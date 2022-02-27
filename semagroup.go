package semagroup

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type Group struct {
	n    int64
	sema *semaphore.Weighted
}

func New(n int64) *Group {
	group := &Group{
		n:    n,
		sema: semaphore.NewWeighted(n),
	}
	return group
}

func (g *Group) AcquireAndGo(ctx context.Context, fn func()) (err error) {
	err = g.sema.Acquire(ctx, 1)
	if err == nil {
		go func() {
			defer g.sema.Release(1)
			fn()
		}()
	}
	return
}
func (g *Group) Wait(ctx context.Context) (err error) {
	return g.sema.Acquire(ctx, g.n)
}
