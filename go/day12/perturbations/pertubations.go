package perturbations

import (
	"bytes"
	"encoding/gob"
)

type Params struct {
	Record []rune
	Counts []int
	Size   int
}

func (p Params) Hash() string {
	var b bytes.Buffer
	gob.NewEncoder(&b).Encode(p)
	return b.String()
}

type Solver struct {
	Cache map[string]int
}

func (s *Solver) ValidArrangements(record []rune, counts []int, size int) int {
	if len(record) == 0 {
		if (len(counts) == 1 && counts[0] == size) || (len(counts) == 0 && size == 0) {
			return 1
		}
		return 0
	}
	params := Params{record, counts, size}.Hash()
	if v, ok := s.Cache[params]; ok {
		return v
	}
	v := s.validArrangements(record, counts, size)
	s.Cache[params] = v
	return v
}

func (s *Solver) validArrangements(record []rune, counts []int, size int) int {
	if len(record) == 0 {
		if (len(counts) == 1 && counts[0] == size) || (len(counts) == 0 && size == 0) {
			return 1
		}
		return 0
	}

	spring := record[0]
	springs := record[1:]
	count := 0
	if len(counts) > 0 {
		count = counts[0]
	}
	if spring == '?' {
		return s.ValidArrangements(append([]rune{'#'}, springs...), counts, size) + s.ValidArrangements(append([]rune{'.'}, springs...), counts, size)
	}
	if spring == '#' {
		if size > count {
			return 0
		}
		return s.ValidArrangements(springs, counts, size+1)
	}
	if spring == '.' {
		if size == 0 {
			return s.ValidArrangements(springs, counts, 0)
		}
		if size == count {
			return s.ValidArrangements(springs, counts[1:], 0)
		}
		return 0
	}
	panic("err")
}
