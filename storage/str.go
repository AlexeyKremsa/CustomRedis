package storage

func (s *Storage) SetStr(key, value string, expirationSec int64) {
	s.set(key, value, expirationSec)
}

func (s *Storage) SetStrNX(key, value string, expirationSec int64) error {
	return s.setNX(key, value, expirationSec)
}

func (s *Storage) GetStr(key string) (string, error) {
	val, err := s.get(key)
	if err != nil {
		return "", err
	}

	if str, ok := val.(string); ok {
		return str, nil
	}
	return "", newErrCustom(errWrongType)
}
