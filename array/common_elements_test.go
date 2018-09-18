package array

import (
	"reflect"
	"testing"
)

func TestCommonElements(t *testing.T) {
	var empty []int
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Test1",
			args: args{
				nums1: []int{1, 3, 4, 6, 7, 9},
				nums2: []int{1, 2, 4, 5, 9, 10},
			},
			want: []int{1, 4, 9},
		},
		{
			name: "Test2",
			args: args{
				nums1: []int{1, 2, 9, 10, 11, 12},
				nums2: []int{0, 1, 2, 3, 4, 5, 8, 9, 10, 12, 14, 15},
			},
			want: []int{1, 2, 9, 10, 12},
		},
		{
			name: "Test3",
			args: args{
				nums1: []int{1, 2, 3, 4, 5},
				nums2: []int{6, 7, 8, 9, 10, 11},
			},
			want: empty,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CommonElements(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CommonElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
