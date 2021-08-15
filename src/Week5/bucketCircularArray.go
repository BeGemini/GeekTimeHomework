package Week5

type bucketCircularArray struct {
	size int
	head int
	tail int
	data []*bucket
}

func (e *bucketCircularArray) getTail() *bucket {
	if e.size == 0 {
		return nil
	}
	return e.data[e.size-1]
}

func (e *bucketCircularArray) convert(index int) int {
	return (index + e.head) % len(e.data)
}
