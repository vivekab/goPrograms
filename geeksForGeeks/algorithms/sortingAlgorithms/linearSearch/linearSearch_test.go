package main

import "testing"

func Test_linearSearch(t *testing.T) {
	type args struct {
		slice   []int
		element int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name:"Success",
			args:args{
				slice:   []int{1, 2, 3, 4, 5, 6, 7, 22, 8, 9, 10, 11, 12, 13, 15, 16, 16},
				element: 10,
			},
		},
		{
			name:"Error",
			args:args{
				slice:   []int{1, 2, 3, 4, 5, 6, 7, 22, 8, 9, 10, 11, 12, 13, 15, 16, 16},
				element: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linearSearch(tt.args.slice, tt.args.element); got != tt.want {
				t.Errorf("linearSearch() = %v, want %v", got, tt.want)
			}
		})
	}
}