package encryption

import (
	"fmt"
	"sync"
)

type InMemoryKeyProvider struct {
	mu      sync.RWMutex
	keys    map[string][]byte
	current string
}

func NewInMemoryKeyProvider() *InMemoryKeyProvider {
	return &InMemoryKeyProvider{
		keys: make(map[string][]byte),
	}
}

func (k *InMemoryKeyProvider) GetCurrentKeyID() string {
	k.mu.RLock()
	defer k.mu.RUnlock()
	return k.current
}

func (k *InMemoryKeyProvider) GetKey(keyID string) ([]byte, error) {
	k.mu.RLock()
	defer k.mu.RUnlock()

	key, exists := k.keys[keyID]
	if !exists {
		return nil, fmt.Errorf("key not found: %s", keyID)
	}
	return key, nil
}

func (k *InMemoryKeyProvider) AddKey(keyID string, key []byte) {
	k.mu.Lock()
	defer k.mu.Unlock()

	k.keys[keyID] = key
	if k.current == "" {
		k.current = keyID
	}
}

func (k *InMemoryKeyProvider) SetCurrentKey(keyID string) error {
	k.mu.Lock()
	defer k.mu.Unlock()

	if _, exists := k.keys[keyID]; !exists {
		return fmt.Errorf("key not found: %s", keyID)
	}
	k.current = keyID
	return nil
}
