package main

import (
	"context"
	"testing"
	addservice "mqx/proto"
)

func PubTest(t *testing.T) {
	s := AddServiceServer{}

	// set up test cases
	tests := []struct{
		name string
		want bool
	} {
		{
			name: "Serie A",
			want: true,
		},
		{
			name: "123",
			want: false,
		},
		{
			name: "GolRoma",
			want: true,
		},
	}

	for _, tt := range tests {
		req := &addservice.TopicMessage{}
		resp, err := s.Pub(context.Background(), req)
		if err != nil {
			t.Errorf("HelloTest(%v) got unexpected error")
		}
		if resp.Ack != tt.want {
			t.Errorf("HelloText(%v)=%v, wanted %v", tt.name, resp.Ack, tt.want)
		}
	}
}
