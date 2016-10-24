package main

import (
	"fmt"

	pb "github.com/azmodb/azmo/azmopb"
)

type fmtEncoder struct{}

func (e fmtEncoder) Encode(ev *pb.Event) error {
	_, err := fmt.Printf("%s\n", ev)
	return err
}
