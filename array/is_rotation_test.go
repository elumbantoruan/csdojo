package array

import "testing"

func TestIsRotation(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	listA := []int{1, 2, 3, 4, 5, 6, 7}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test1",
			args: args{
				a: listA,
				b: []int{4, 5, 6, 7, 8, 1, 2, 3},
			},
			want: false,
		},
		{
			name: "test2",
			args: args{
				a: listA,
				b: []int{4, 5, 6, 7, 1, 2, 3},
			},
			want: true,
		},
		{
			name: "test3",
			args: args{
				a: []int{4, 5, 6, 7, 1, 2, 3},
				b: []int{4, 5, 6, 9, 1, 2, 3},
			},
			want: false,
		},
		{
			name: "test4",
			args: args{
				a: listA,
				b: []int{4, 6, 5, 7, 1, 2, 3},
			},
			want: false,
		},
		{
			name: "test5",
			args: args{
				a: listA,
				b: []int{4, 5, 6, 7, 0, 2, 3},
			},
			want: false,
		},
		{
			name: "test6",
			args: args{
				a: listA,
				b: []int{1, 2, 3, 4, 5, 6, 7},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsRotation(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("IsRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
