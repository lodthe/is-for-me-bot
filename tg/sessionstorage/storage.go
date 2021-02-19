package sessionstorage

import (
	"sync"
)

type Storage struct {
	sessionLockers map[int]sync.Locker
	lock           sync.Locker
}

func NewStorage() *Storage {
	return &Storage{
		sessionLockers: map[int]sync.Locker{},
		lock:           &sync.Mutex{},
	}
}

func (s *Storage) AcquireLock(telegramID int) sync.Locker {
	s.lock.Lock()
	defer s.lock.Unlock()

	sessionLocker, exists := s.sessionLockers[telegramID]
	if !exists {
		sessionLocker = &sync.Mutex{}
		s.sessionLockers[telegramID] = sessionLocker
	}

	return sessionLocker
}
