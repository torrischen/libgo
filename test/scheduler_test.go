package test

import (
	"testing"
	"time"

	"github.com/torrischen/libgo/scheduler"
)

type core struct {
	scheduler.BaseCore
}

func TestScheduler(t *testing.T) {
	s := scheduler.NewScheduler()
	defer s.Stop()

	bc := scheduler.NewBaseCore(func(i interface{}) { t.Log(i) }, 100, 1)

	c := &core{
		*bc,
	}

	s.AppendCore("test", c)

	s.Push("test", "hello world")

	time.Sleep(time.Second * 1)
}
