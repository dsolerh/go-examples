package searcher

type SearchFn = func(data []byte) []byte

type lineSearcher struct {
	line         []byte
	searchers    []SearchFn
	found        []byte
	currentIndex int
	at           int
}

func (s *lineSearcher) Next() bool {
	for ; s.currentIndex < len(s.line); s.currentIndex++ {
		for _, searchFn := range s.searchers {
			found := searchFn(s.line[s.currentIndex:])
			// if found something
			if len(found) > 0 {
				s.found = found
				s.at = s.currentIndex
				s.currentIndex += len(found)
				return true
			}
		}
	}

	return false
}

func (s *lineSearcher) Item() ([]byte, int) {
	return s.found, s.at
}

func NewLineSearcher(line []byte, searchers ...SearchFn) Line {
	return &lineSearcher{
		line:         line,
		searchers:    searchers,
		currentIndex: 0,
	}
}
