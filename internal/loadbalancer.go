package main

import "sync/atomic"

type LoadBalancer struct {
	Backends []*Backend
	current  uint64
}

func (lb *LoadBalancer) NextBackend() *Backend {
	next := atomic.AddUint64(&lb.current, uint64(1)) % uint64(len(lb.Backends))

	for i := 0; i < len(lb.Backends); i++ {
		idx := (int(next) + i) % len(lb.Backends)
		if lb.Backends[idx].IsAlive() {
			return lb.Backends[idx]
		}
	}
	return nil
}
