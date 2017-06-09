// Copyright (c) 2010-2014 - Gustavo Niemeyer <gustavo@niemeyer.net>
// Copyright (c) 2016 - Markus Sonderegger <marraison@gmail.com>
//
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice, this
//    list of conditions and the following disclaimer.
// 2. Redistributions in binary form must reproduce the above copyright notice,
//    this list of conditions and the following disclaimer in the documentation
//    and/or other materials provided with the distribution.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
// DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT OWNER OR CONTRIBUTORS BE LIABLE FOR
// ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
// (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
// LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
// ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
// SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

// Package clock offers a vector clock implementation.
package clock

import (
	"sort"
	"time"

	"github.com/azmodb/clock/pb"
)

// Condition constants define how to compare a vector clock against
// another.
type Condition int

const (
	Equals Condition = 1 << iota
	Ancestor
	Descendant
	Concurrent
)

// Clock represents a vector clock.
type Clock struct {
	*pb.Clock
}

// New creates and initializes a new Clock using data as its initial
// contents, which can be nil.
func New(data *pb.Clock) *Clock {
	if data != nil {
		return &Clock{Clock: data}
	}

	now := time.Now().UTC().UnixNano()
	return &Clock{
		Clock: &pb.Clock{
			Items:    []*pb.Item{},
			Created:  now,
			Modified: now,
		},
	}
}

// Reset resets the clock state.
func (c *Clock) Reset() { c.Clock.Reset() }

// Copy returns a copy of c.
func (c *Clock) Copy() *Clock {
	nc := &Clock{
		Clock: &pb.Clock{
			Items:    make([]*pb.Item, 0, len(c.Items)),
			Created:  c.Created,
			Modified: c.Modified,
		},
	}
	for _, item := range c.Items {
		nc.Items = append(nc.Items, &pb.Item{
			Tick: item.Tick,
			ID:   item.ID,
			When: item.When,
		})
	}
	return nc
}

// Merge merges clock into c, so that c becomes a descendant of clock.
// This means that every clock tick in clock which does not exist in c
// or which is smaller in c will be copied from clock to c.
func (c *Clock) Merge(clock *Clock) {
	n := 0
	for _, item := range clock.Items {
		if index, found := c.find(item.ID); found {
			v := c.Items[index]
			if v.Tick < item.Tick {
				v.Tick = item.Tick
			}
		} else {
			n++
		}
	}

	if n > 0 {
		items := make([]*pb.Item, len(c.Items), len(c.Items)+n)
		copy(items, c.Items)
		c.Items = items[:len(items):cap(items)]

		for _, item := range clock.Items { // TODO: optimize
			if _, found := c.find(item.ID); !found {
				c.Items = append(c.Items, &pb.Item{
					Tick: item.Tick,
					ID:   item.ID,
					When: item.When,
				})
			}
		}
		sort.Sort(c.Clock)
	}
}

func (c *Clock) find(id string) (index int, found bool) {
	index = sort.Search(len(c.Items), func(i int) bool {
		return id < c.Items[i].ID
	})
	if index > 0 && c.Items[index-1].ID >= id {
		return index - 1, true
	}
	return index, false
}

/*
func (c *Clock) Find(id string) (*pb.Item, bool) {
	index, found := c.find(id)
	if !found {
		return nil, false
	}
	item := c.Items[index]
	return &pb.Item{
		Tick: item.Tick,
		ID:   item.ID,
		When: item.When,
	}, true
}
*/

// Update increments id's vector clock ticks in c. The when update time
// is associated with id and may be used for pruning the vector clock.
// When may have any unit.
func (c *Clock) Update(id string, when int64) {
	index, found := c.find(id)
	if found {
		c.Items[index].Tick++
		if when > c.Items[index].When {
			c.Items[index].When = when
		}
	} else {
		if cap(c.Items) >= len(c.Items)+1 {
			c.Items = c.Items[:len(c.Items)+1]
		} else {
			item := make([]*pb.Item, len(c.Items)+1)
			copy(item, c.Items)
			c.Items = item
		}
		copy(c.Items[index+1:], c.Items[index:])
		c.Items[index] = &pb.Item{Tick: 1, ID: id}
	}
	if c.Items[index].Tick > c.Last {
		c.Last = c.Items[index].Tick
	}
	c.Modified = time.Now().UTC().UnixNano()
}

// Compare returns whether b matches any one of the conditions within
// condition (Equal, Ancestor, Descendant, or Concurrent).
func Compare(a, b *Clock, condition Condition) bool {
	na, nb := len(a.Items), len(b.Items)
	var state Condition

	switch {
	case na < nb:
		if condition&(Descendant|Concurrent) == 0 {
			return false
		}
		state = Descendant
	case na > nb:
		if condition&(Ancestor|Concurrent) == 0 {
			return false
		}
		state = Ancestor
	default:
		state = Equals
	}

	diff := nb - na
	for i := 0; i < len(b.Items); i++ {
		if index, found := a.find(b.Items[i].ID); found {
			itema, itemb := a.Items[index], b.Items[i]
			switch {
			case itemb.Tick > itema.Tick:
				if state == Equals {
					if condition&Descendant == 0 {
						return false
					}
					state = Descendant
				} else if state == Ancestor {
					return condition&Concurrent != 0
				}
			case itemb.Tick < itema.Tick:
				if state == Equals {
					if condition&Ancestor == 0 {
						return false
					}
					state = Ancestor
				} else if state == Descendant {
					return condition&Concurrent != 0
				}
			}
		} else {
			if state == Equals {
				return condition&Concurrent == 0
			} else if diff--; diff < 0 {
				return condition&Concurrent != 0
			}
		}
	}
	return condition&state != 0
}

/*
func (c *Clock) Size() int {
	if c == nil {
		return 0
	}
	return proto.Size(c.Clock)
}

func (c *Clock) MustMarshal() []byte {
	b, err := c.Marshal()
	if err != nil {
		panic("marshal clock: " + err.Error())
	}
	return b
}

func (c *Clock) MustUnmarshal(b []byte) {
	if err := c.Unmarshal(b); err != nil {
		panic("unmarshal clock: " + err.Error())
	}
}

func (c *Clock) Marshal() ([]byte, error) {
	buf := proto.NewBuffer(nil)
	if err := buf.Marshal(c.Clock); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (c *Clock) Unmarshal(b []byte) error {
	buf := proto.NewBuffer(b)
	if c.Clock == nil {
		c.Clock = &pb.Clock{}
	}
	if err := buf.Unmarshal(c.Clock); err != nil {
		return err
	}
	sort.Sort(c.Clock)
	return nil
}
*/
