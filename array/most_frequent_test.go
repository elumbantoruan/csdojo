package array

import "testing"

func TestMostFrequent(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{nums: []int{1, 3, 1, 3, 2, 1}},
			want: 1,
		},
		{
			name: "test 2",
			args: args{nums: []int{3, 3, 1, 3, 2, 1}},
			want: 3,
		},
		{
			name: "test 3",
			args: args{nums: []int{}},
			want: -1,
		},
		{
			name: "test 4",
			args: args{nums: []int{0}},
			want: 0,
		},
		{
			name: "test 5",
			args: args{nums: []int{0, -1, 10, 10, -1, 10, -1, -1, -1, 1}},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MostFrequent(tt.args.nums); got != tt.want {
				t.Errorf("MostFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
