package Week5

import "time"

type RollingNumber struct {
	startMS      int64
	cap          int
	perBucketCap int
	head         int
	tail         int
	buckets      []*bucket
}

type Event int

const (
	Success Event = iota
	Failure
	Timeout
	Rejection
)

func NewRollingNmber(cap int, perBucketCap int) *RollingNumber {
	start := time.Now().UnixNano() / 1e6
	return &RollingNumber{
		startMS:      start,
		cap:          cap,
		perBucketCap: perBucketCap,
		head:         0,
		tail:         0,
		buckets:      make([]*bucket, cap),
	}
}

func (e *RollingNumber) Increase(event Event) {
	currentTime := time.Now().UnixNano() / 1e6

}

func (e *RollingNumber) GetRollingSum() {

}
