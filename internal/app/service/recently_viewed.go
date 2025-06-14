package service

import "sync"

type RecentlyViewedStore interface {
	Add(userID, articleID int64)
	Get(userID int64) []int64
}

type InMemoryRecentlyViewedStore struct {
	mu      sync.Mutex
	data    map[int64][]int64
	maxSize int
}

func NewInMemoryRecentlyViewedStore(maxSize int) *InMemoryRecentlyViewedStore {
	return &InMemoryRecentlyViewedStore{
		data:    make(map[int64][]int64),
		maxSize: maxSize,
	}
}

func (s *InMemoryRecentlyViewedStore) Add(userID, articleID int64) {
	s.mu.Lock()
	defer s.mu.Unlock()

	views := s.data[userID]
	// Remove if already present
	for i, id := range views {
		if id == articleID {
			views = append(views[:i], views[i+1:]...)
			break
		}
	}
	// Prepend
	views = append([]int64{articleID}, views...)
	// Remove extras
	if len(views) > s.maxSize {
		views = views[:s.maxSize]
	}
	s.data[userID] = views
}

func (s *InMemoryRecentlyViewedStore) Get(userID int64) []int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.data[userID]
}
