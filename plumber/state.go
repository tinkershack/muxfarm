package plumber

import (
	tspb "google.golang.org/protobuf/types/known/timestamppb"
)

func NewStatus(code State) *Status {
	status := &Status{
		State:           *code.Enum(),
		UpdateTimestamp: tspb.Now(),
	}
	return status
}
