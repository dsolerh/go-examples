package searcher

type Line interface {
	Next() bool
	Item() ([]byte, int)
}
