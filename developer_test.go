//"https://itunes.apple.com/lookup?id=284882218&country=us&entity=software"
package main

import (
	"testing"
)

func TestDeveloperLookup(t *testing.T) {
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
				id:      368677371,
				country: "us",
			},
			want: "https://apps.apple.com/us/developer/uber-technologies-inc/id368677371?uo=4",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {
			if got := DeveloperLookup(tt.args.id, tt.args.country); got.DeveloperURL != tt.want || got.CountApps != 8 {
				t.Errorf("DeveloperLookup() = %v, want %v", got, tt.want)
			}
		})
	}
}
