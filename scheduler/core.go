package scheduler

import (
	"context"
)

type Core interface {
	Listen()
	Shut()
	CacheSize() int
	Push(interface{})
	Error() <-chan error
}

type BaseCore struct {
	Ctx        context.Context
	Cancel     context.CancelFunc
	HandleFunc func(interface{}) error
	Cache      chan interface{}
	ErrCh      chan error
}

func NewBaseCore(f func(interface{}) error, cacheSize int, errChSize int) *BaseCore {
	ctx, cancel := context.WithCancel(context.Background())
	return &BaseCore{
		Ctx:        ctx,
		Cancel:     cancel,
		HandleFunc: f,
		Cache:      make(chan interface{}, cacheSize),
		ErrCh:      make(chan error, errChSize),
	}
}

func (c *BaseCore) Listen() {
	go func() {
		for {
			select {
			case <-c.Ctx.Done():
				return
			case event := <-c.Cache:
				c.HandleFunc(event)
			}
		}
	}()
}

func (c *BaseCore) Shut() {
	c.Cancel()
	close(c.Cache)
	close(c.ErrCh)
}

func (c *BaseCore) CacheSize() int {
	return cap(c.Cache)
}

func (c *BaseCore) Push(event interface{}) {
	c.Cache <- event
}

func (c *BaseCore) Error() <-chan error {
	return c.ErrCh
}
