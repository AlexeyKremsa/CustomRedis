package storage

func (s *Storage) SetMap(key string, value map[string]string, expirationSec int64) {
	s.set(key, value, expirationSec)
}

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
