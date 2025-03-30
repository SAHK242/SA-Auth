package cache

import (
	"sync"

	"github.com/eko/gocache/lib/v4/codec"
)

const (
	OpGet Op = "get"
	OpSet Op = "set"
)

type (
	Stats struct {
		stats *codec.Stats
		mu    sync.Mutex
	}

	Op string

	cacheStatsItem struct {
		op      Op
		success bool
	}
)

func newStats() *Stats {
	return &Stats{
		stats: &codec.Stats{},
	}
}

func (s *Stats) Hits() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stats.Hits
}

func (s *Stats) Miss() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stats.Miss
}

func (s *Stats) SetSuccess() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stats.SetSuccess
}

func (s *Stats) SetError() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.stats.SetError
}

func (s *Stats) get() *codec.Stats {
	s.mu.Lock()
	defer s.mu.Unlock()
	stats := *s.stats
	return &stats
}

func (s *Stats) increaseHits() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stats.Hits++
}

func (s *Stats) increaseMiss() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.stats.Miss++
}

func (s *Stats) increaseSet(success bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if success {
		s.stats.SetSuccess++
	} else {
		s.stats.SetError++
	}
}

func (s *Stats) updateStats(op Op, success bool) {
	switch op {
	case OpGet:
		if success {
			s.increaseHits()
		} else {
			s.increaseMiss()
		}
	case OpSet:
		s.increaseSet(success)
	}
}
