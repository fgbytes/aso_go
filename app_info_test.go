package main

import (
	"testing"
)

func TestLookUp(t *testing.T) {
	type args struct {
		id      int
		country string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test1",
			args: args{
				id:      368677368,
				country: "us",
			},
			want: "https://www.uber.com",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LookUp(tt.args.id, tt.args.country); got.SellerURL != tt.want {
				t.Errorf("LookUp() = %v, want %v", got, tt.want)
			}

		})
	}
}
