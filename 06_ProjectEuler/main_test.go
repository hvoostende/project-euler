package main

import "testing"

func Test_divisableBy(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "Test1",
			args: args{10, 5},
			want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := divisableBy(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("divisableBy(%v, %v) = %v, want %v", tt.args.x, tt.args.y, got, tt.want)
			}
		})
	}
}
