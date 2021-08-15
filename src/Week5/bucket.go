package Week5

const (
	success = iota
	failure
	timeout
	rejection
)

type bucket struct {
	success     int64
	failure     int64
	timeout     int64
	rejection   int64
	windowStart int64
}

func newBucket(time int64) *bucket {
	_bucket := &bucket{
		windowStart: time,
	}
	return _bucket
}
