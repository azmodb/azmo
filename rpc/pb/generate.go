//go:generate brick-protoc -o codec.pb.go  codec.proto

package pb

func (m *Error) Error() string { return m.Description }
