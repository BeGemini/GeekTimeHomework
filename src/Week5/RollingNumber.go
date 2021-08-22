package Week5

import (
	"time"
)

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
	currentTime := getCurrentTime()
	timeDis := currentTime - e.startMS
	index := (timeDis / int64(e.perBucketCap)) % int64(e.cap)
	e.head = int(timeDis / int64(e.perBucketCap))
	if int(index) > e.tail {
		e.tail += 1
	}
	switch event {
	case Success:
		e.buckets[index].success += 1
	case Timeout:
		e.buckets[index].timeout += 1
	case Rejection:
		e.buckets[index].rejection += 1
	case Failure:
		e.buckets[index].failure += 1
	}
}

func (e *RollingNumber) GetRollingSum(event Event) int64 {
	var sum int64
	for i := 0; i < e.cap; i++ {
		switch event {
		case Success:
			sum += e.buckets[i].success
		case Failure:
			sum += e.buckets[i].failure
		case Timeout:
			sum += e.buckets[i].timeout
		case Rejection:
			sum += e.buckets[i].rejection
		}
	}
	return sum
}

func getCurrentTime() int64 {
	return time.Now().UnixNano() / 1e6
}
