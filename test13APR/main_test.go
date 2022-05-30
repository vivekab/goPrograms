package main

import (
	"reflect"
	"testing"
)

func TestEvaluateSlice(t *testing.T) {
	tests := []struct {
		name   string
		input  []string
		output map[string]bool
	}{
		{
			name:   "EmptySlice",
			input:  nil,
			output: map[string]bool{},
		},
	}

	for i := range tests {
		t.Run(tests[i].name,func(t *testing.T) {
			got := evaluateSlice(tests[i].input)
			if !reflect.DeepEqual(tests[i].output, got){
				t.Errorf("Expected: %v , Got: %v",tests[i].output,got)
			}
		})
	}
}
