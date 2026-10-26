// SPDX-FileCopyrightText: 2026 Xquik contributors
//
// SPDX-License-Identifier: Apache-2.0

package paramutil

import "sync"

type SyncMap[K comparable, V any] struct {
	mu     sync.RWMutex
	values map[K]V
}

func (m *SyncMap[K, V]) Load(key K) (V, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	value, ok := m.values[key]
	return value, ok
}

func (m *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if existing, ok := m.values[key]; ok {
		return existing, true
	}
	m.store(key, value)
	return value, false
}

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
