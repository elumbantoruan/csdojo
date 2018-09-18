package strings

import "testing"

func TestNonRepeatingChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "Test1",
			args: args{s: "aabcb"},
			want: 'c',
		},
		{
			name: "Test2",
			args: args{s: "aaacdddbg"},
			want: 'c',
		},
		{
			name: "Test3",
			args: args{s: "aabb"},
			want: ' ',
		},
		{
			name: "Test4",
			args: args{s: "aabbd"},
			want: 'd',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NonRepeatingChar(tt.args.s); got != tt.want {
				t.Errorf("NonRepeatingChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
