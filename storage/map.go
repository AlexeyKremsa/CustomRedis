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
