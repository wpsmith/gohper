package sync2

import (
	"unsafe"

	"sync"
)

const APPEND = "01235"

type MRWMutex struct {
	lock  sync.RWMutex
	locks map[string]*sync.RWMutex
}

func NewMRWMutex() *MRWMutex {
	return &MRWMutex{
		locks: make(map[string]*sync.RWMutex),
	}
}

func (m *MRWMutex) NewId(id string) string {
	for {
		if _, b := m.locker(id); b {
			return id
		}
		id = id + APPEND[uintptr(unsafe.Pointer(&id))&4:]
	}
}

func (m *MRWMutex) locker(lockId string) (l *sync.RWMutex, b bool) {
	m.lock.RLock()
	if l = m.locks[lockId]; l == nil {
		m.lock.RUnlock()
		m.lock.Lock()
		if l = m.locks[lockId]; l == nil {
			l = &sync.RWMutex{}
			m.locks[lockId] = l
			b = true
		}
		m.lock.Unlock()
	} else {
		m.lock.RUnlock()
	}
	return
}

func (m *MRWMutex) locker2(lockId string) *sync.RWMutex {
	l, _ := m.locker(lockId)
	return l
}

func (m *MRWMutex) RLocker(lockId string) sync.Locker {
	return m.locker2(lockId).RLocker()
}

func (m *MRWMutex) Lock(lockId string) {
	m.locker2(lockId).Lock()
}

func (m *MRWMutex) Unlock(lockId string) {
	m.locker2(lockId).Unlock()
}

func (m *MRWMutex) RLock(lockId string) {
	m.locker2(lockId).RLock()
}

func (m *MRWMutex) RUnlock(lockId string) {
	m.locker2(lockId).RUnlock()
}
