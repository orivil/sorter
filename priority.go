package sorter
import (
	"sort"
	"errors"
)

var ErrPriorityExist = errors.New("sorter.PrioritySorter: priority exist")

// PrioritySorter could cache the sorted data, also could add a priority, or delete a priority
type PrioritySorter struct {
	priorities map[string]int
	cache []string
	reverse []string
	keys []string
}

func NewPrioritySorter(data map[string]int) *PrioritySorter {
	return &PrioritySorter {
		priorities: data,
	}
}

func (s *PrioritySorter) Del(name string) {
	delete(s.priorities, name)
	s.cache = nil
	s.reverse = nil
}

func (s *PrioritySorter) Add(name string, priority int) error {
	if _, ok := s.priorities[name]; ok {
		return ErrPriorityExist
	} else {
		s.priorities[name] = priority
		s.cache = nil
		s.reverse = nil
		return nil
	}
}

func (s *PrioritySorter) Sort() []string {
	if s.cache == nil {
		s.sort(false)
	}
	return s.cache
}

func (s *PrioritySorter) SortReverse() []string {
	if s.reverse == nil {
		s.sort(true)
	}
	return s.reverse
}

func (s *PrioritySorter) sort(reverse bool) {
	s.keys = make([]string, len(s.priorities))
	index := 0
	for name, _ := range s.priorities {
		s.keys[index] = name
		index++
	}
	if reverse {
		sort.Sort(sort.Reverse(s))
		s.reverse = s.keys
	} else {
		sort.Sort(s)
		s.cache = s.keys
	}
}

func (s *PrioritySorter) Len() int {
	return len(s.keys)
}

func (s *PrioritySorter) Swap(i, j int) {
	s.keys[i], s.keys[j] = s.keys[j], s.keys[i]
}

func (s *PrioritySorter) Less(i, j int) bool {
	return s.priorities[s.keys[i]] < s.priorities[s.keys[j]]
}
