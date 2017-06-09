//go:generate protoc --proto_path=. --go_out=plugins=grpc:. azmo.proto

package azmopb

const (
	Decrement = Event_DECREMENT
	Increment = Event_INCREMENT
	Put       = Event_PUT
	Delete    = Event_DELETE
	Get       = Event_GET
	Range     = Event_RANGE
	Watch     = Event_WATCH
)

func NewEvent(t Event_Type, key []byte, data interface{}, created, current int64) *Event {
	m := &Event{}
	m.Init(t, key, data, created, current)
	return m
}

func (m *Event) Init(t Event_Type, key []byte, data interface{}, created, current int64) {
	m.Type = t
	m.Key = key
	m.Current = current
	m.Created = created
	switch t := data.(type) {
	case []byte:
		m.Content = t
	case int64:
		m.Numeric = t
	}
}

func (m *BatchRequest) Decrement(key []byte, value int64, tombstone bool) {
	m.Arguments = append(m.Arguments, &Argument{
		&Argument_Decrement{
			&DecrementRequest{
				Key:       key,
				Value:     value,
				Tombstone: tombstone,
			},
		},
	})
}

func (m *BatchRequest) Increment(key []byte, value int64, tombstone bool) {
	m.Arguments = append(m.Arguments, &Argument{
		&Argument_Increment{
			&IncrementRequest{
				Key:       key,
				Value:     value,
				Tombstone: tombstone,
			},
		},
	})
}

func (m *BatchRequest) Put(key []byte, value []byte, tombstone bool) {
	m.Arguments = append(m.Arguments, &Argument{
		&Argument_Put{
			&PutRequest{
				Key:       key,
				Value:     value,
				Tombstone: tombstone,
			},
		},
	})
}

func (m *BatchRequest) Delete(key []byte) {
	m.Arguments = append(m.Arguments, &Argument{
		&Argument_Delete{
			&DeleteRequest{
				Key: key,
			},
		},
	})
}
