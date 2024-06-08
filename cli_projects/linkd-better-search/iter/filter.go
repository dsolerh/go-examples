package iter

//
// type filterIter[T any] struct {
// 	inner Iterator[T]
// 	pred  func(T) bool
// }
//
// // Filter returns an Iterator adapter that yields elements from the underlying
// // Iterator for which pred returns true.
// func Filter[T any](it Iterator[T], pred func(T) bool) Iterator[T] {
// 	return &filterIter[T]{
// 		inner: it,
// 		pred:  pred,
// 	}
// }
//
// func (it *filterIter[T]) Next() Option[T] {
// 	var v Option[T]
// 	for v = it.inner.Next(); v.IsSome(); v = it.inner.Next() {
// 		if it.pred(v.Unwrap()) {
// 			break
// 		}
// 	}
// 	return v
// }
