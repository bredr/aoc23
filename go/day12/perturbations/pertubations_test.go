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
		want []string
	}{
		{
			args{"#.#.###", []int{1, 1, 3}},
			[]string{"#.#.###"},
		},
		{
			args{"??", []int{1}},
			[]string{".#", "#."},
		},
		{
			args{"?###????????", []int{3, 2, 1}},
			[]string{".###....##.#", ".###...##..#", ".###..##...#", ".###.##....#", ".###...##.#.", ".###..##..#.", ".###.##...#.", ".###..##.#..", ".###.##..#..", ".###.##.#..."},
		},
	}
	for i, tt := range tests {
		t.Run(fmt.Sprintf("test_%d", i), func(t *testing.T) {
			if got := Perturbations(tt.args.record, tt.args.counts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Perturbations() = %v, want %v", got, tt.want)
			}
		})
	}
}
