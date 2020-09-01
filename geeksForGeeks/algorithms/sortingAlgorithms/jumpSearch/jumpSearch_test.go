package jumpSearch

import "testing"

func TestJumpSearch(t *testing.T) {
	type args struct {
		slice        []int
		element      int
		jumpingIndex int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Success",
			args:args{
				slice:   []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 11, 12, 13, 15, 16, 16},
				element: 10,
				jumpingIndex:4,
			},
			want:11,
		},
		{
			name:"Error",
			args:args{
				slice:   []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 11, 12, 13, 15, 16, 16},
				element: 0,
				jumpingIndex:3,
			},
			want:-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JumpSearch(tt.args.slice, tt.args.element, tt.args.jumpingIndex); got != tt.want {
				t.Errorf("JumpSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}