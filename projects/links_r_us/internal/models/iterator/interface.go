package iterator

type Iterator[T any] interface {
	// gets the item of the iterator
	Item() T
	// Next advances the iterator. If no more items are available or an
	// error occurs, calls to Next() return false.
	Next() bool
	// Error returns the last error encountered by the iterator.
	Error() error
	// Close releases any resources associated with an iterator.
	Close() error
}
