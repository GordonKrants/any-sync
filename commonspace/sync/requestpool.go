package sync

import (
	"context"
	"sync"
)

type RequestPool interface {
	TryTake(peerId, objectId string) bool
	Release(peerId, objectId string)
	QueueRequestAction(peerId, objectId string, action func(ctx context.Context)) (err error)
}

type requestPool struct {
	mu       sync.Mutex
	taken    map[string]struct{}
	pools    map[string]*tryAddQueue
	ctx      context.Context
	cancel   context.CancelFunc
	isClosed bool
}

func NewRequestPool() RequestPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &requestPool{
		ctx:    ctx,
		cancel: cancel,
		taken:  make(map[string]struct{}),
		pools:  make(map[string]*tryAddQueue),
	}
}

func (rp *requestPool) TryTake(peerId, objectId string) bool {
	rp.mu.Lock()
	defer rp.mu.Unlock()
	if rp.isClosed {
		return false
	}

	id := fullId(peerId, objectId)
	if _, exists := rp.taken[id]; exists {
		return false
	}
	rp.taken[id] = struct{}{}
	return true
}

func (rp *requestPool) Release(peerId, objectId string) {
	rp.mu.Lock()
	defer rp.mu.Unlock()

	id := fullId(peerId, objectId)
	delete(rp.taken, id)
}

func (rp *requestPool) QueueRequestAction(peerId, objectId string, action func(ctx context.Context)) (err error) {
	rp.mu.Lock()
	if rp.isClosed {
		rp.mu.Unlock()
		return nil
	}
	var (
		pool   *tryAddQueue
		exists bool
	)
	pool, exists = rp.pools[peerId]
	if !exists {
		pool = newTryAddQueue(100, 100)
		rp.pools[peerId] = pool
		pool.Run()
	}
	rp.mu.Unlock()
	var wrappedAction func()
	wrappedAction = func() {
		if !rp.TryTake(peerId, objectId) {
			pool.TryAdd(objectId, wrappedAction, func() {})
			return
		}
		action(rp.ctx)
		rp.Release(peerId, objectId)
	}
	pool.Replace(objectId, wrappedAction, func() {})
	return nil
}

func (rp *requestPool) Close() {
	rp.mu.Lock()
	defer rp.mu.Unlock()
	rp.isClosed = true
	for _, pool := range rp.pools {
		_ = pool.Close()
	}
}
