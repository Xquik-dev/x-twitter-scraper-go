// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package paramutil

import "sync"

// SyncMap is a typed map that is safe for concurrent use, with the methods of
// [sync.Map] that the encoder caches use. Its zero value is empty.
type SyncMap[K comparable, V any] struct {
	mu     sync.RWMutex
	values map[K]V
}

// Load returns the value stored under key and whether there is one.
func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.values[key]
	return value, ok
}

// LoadOrStore returns the value stored under key and true when there is one.
// Otherwise it stores value under key and returns value and false.
func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if existing, ok := m.values[key]; ok {
		return existing, true
	}
	m.store(key, value)
	return value, false
}

// Store sets the value under key.
func (m *SyncMap[K, V]) Store(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.store(key, value)
}

func (m *SyncMap[K, V]) store(key K, value V) {
	if m.values == nil {
		m.values = map[K]V{}
	}
	m.values[key] = value
}
