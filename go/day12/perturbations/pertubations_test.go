package perturbations

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPerturbations(t *testing.T) {
	type args struct {
		record string
		counts []int
	}
	tests := []struct {
		args args
		want int
	}{
		{
			args{"#.#.###", []int{1, 1, 3}},
			1,
		},
		{
			args{"??", []int{1}},
			2,
		},
		{
			args{"?###????????", []int{3, 2, 1}},
			10,
		},
	}
	for i, tt := range tests {
		s := Solver{Cache: make(map[string]int)}
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			if got := s.ValidArrangements([]rune(tt.args.record), tt.args.counts, 0); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Perturbations() = %v, want %v", got, tt.want)
			}
		})
	}
}
