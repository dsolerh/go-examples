package main

type Stack struct {
	itemsCounter int
	values       map[int][]int
}

func NewStack() *Stack {
	return &Stack{
		itemsCounter: 0,
		values:       map[int][]int{},
	}
}

func (s *Stack) push(v int) {
	if _, exist := s.values[v]; exist {
		s.values[v] = append(s.values[v], s.itemsCounter)
	} else {
		s.values[v] = []int{s.itemsCounter}
	}
	s.itemsCounter++
}

func (s *Stack) pop() *int {
	var moreFrequent *int = nil
	if len(s.values) == 0 {
		return moreFrequent
	}

	maxCount := 0
	lastTimestamp := 0
	for key, value := range s.values {
		if len(value) > maxCount {
			moreFrequent = &key
			maxCount = len(value)
			lastTimestamp = value[len(value)-1]
		} else if len(value) == maxCount {
			if value[len(value)-1] > lastTimestamp {
				moreFrequent = &key
				lastTimestamp = value[len(value)-1]
			}
		}
	}
	if moreFrequent != nil {
		val := *moreFrequent
		// remove the last element of the timestamps
		s.values[val] = s.values[val][:len(s.values[val])-1]

		// remove the key from the map if there're no more elements inside
		if len(s.values[val]) == 0 {
			delete(s.values, val)
		}
	}

	return moreFrequent
}

func main() {}
