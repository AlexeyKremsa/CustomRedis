package storage

func (s *Storage) SetMap(key string, value map[string]string, expirationSec int64) {
	s.set(key, value, expirationSec)
}

func (s *Storage) GetMap(key string) (map[string]string, error) {
	val, err := s.get(key)
	if err != nil {
		return nil, err
	}

	if item, ok := val.(map[string]string); ok {
		return item, nil
	}

	return nil, newErrCustom(errWrongType)
}

func (s *Storage) GetMapItem(key, itemKey string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if item, ok := s.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return "", newErrCustom(errNotExist)
		}

		if m, ok := item.Value.(map[string]string); ok {
			return m[itemKey], nil
		}
		return "", newErrCustom(errWrongType)
	}
	return "", newErrCustom(errNotExist)
}
