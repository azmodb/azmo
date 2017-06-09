// Copyright 2016 Markus Sonderegger. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate protoc --go_out=. clock.proto

package pb

func (c Clock) Swap(i, j int) {
	c.Items[i], c.Items[j] = c.Items[j], c.Items[i]
}

func (c Clock) Len() int { return len(c.Items) }

func (c Clock) Less(i, j int) bool {
	return c.Items[i].ID < c.Items[j].ID
}
