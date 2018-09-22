package strings

import (
	"testing"
)

func TestIsOneAway(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s1: "abcde",
				s2: "abfde",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s1: "abcde",
				s2: "abfdegk",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				s1: "abcde",
				s2: "abde",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				s1: "xyz",
				s2: "xyaz",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOneAway(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsOneAway() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsOneAway2ndApproach(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				s1: "abcde",
				s2: "abfde",
			},
			want: true,
		},
		{
			name: "test2",
			args: args{
				s1: "abcde",
				s2: "abfdegk",
			},
			want: false,
		},
		{
			name: "test3",
			args: args{
				s1: "abcde",
				s2: "abde",
			},
			want: true,
		},
		{
			name: "test4",
			args: args{
				s1: "xyz",
				s2: "xyaz",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsOneAway2ndApproach(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("IsOneAway2ndApproach() = %v, want %v", got, tt.want)
			}
		})
	}
}
