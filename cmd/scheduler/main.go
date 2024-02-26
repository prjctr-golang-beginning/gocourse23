package main

import (
	"fmt"
	"time"
)

type Differ interface {
	GetDiff() []string
}

func main() {
	t := TrackTimeBegin()
	defer func() { fmt.Printf("Time pased: %d", t.Passed()) }()
	defer Recover()
}

type TimeTracker int64

func TrackTimeBegin() TimeTracker {
	return TimeTracker(time.Now().UnixNano())
}

func (t TimeTracker) Passed() int64 {
	return time.Now().UnixNano() - int64(t)
}

func Recover() {
	if r := recover(); r != nil {
		newErr := fmt.Errorf("recovered with message: '%s', stack trace:\r\n", r)
		fmt.Println(newErr)
	} else {
		fmt.Println(`Panics not found`)
	}
}
