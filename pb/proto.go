//go:generate protoc --proto_path=$GOPATH/src:. --go_out=plugins=grpc:. azmo.proto

package pb
