package storage

// SetMap sets map to specified key
func (s *Storage) SetMap(key string, value map[string]string, expirationSec uint64) {
	s.set(key, value, expirationSec)
}

// GetMap returns map by specified key
func (s *Storage) GetMap(key string) (map[string]string, error) {
	item := s.get(key)
	if item == nil {
		return nil, nil
	}

	if m, ok := item.Value.(map[string]string); ok {
		return m, nil
	}

	return nil, newErrCustom(errWrongType)
}

// GetMapItem returns an element stored in a map which is stored by specified key
func (s *Storage) GetMapItem(key, itemKey string) (string, error) {
	item := s.get(key)
	if item == nil {
		return "", nil
	}

	if m, ok := item.Value.(map[string]string); ok {
		return m[itemKey], nil
	}
	return "", newErrCustom(errWrongType)
}
