package Week5

type bucket struct {
	success     int64
	failure     int64
	timeout     int64
	rejection   int64
	windowStart int64
	cap         int // how long the bucket mean
}
