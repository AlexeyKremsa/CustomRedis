package storage

// SetStr sets key to hold the string value
func (s *Storage) SetStr(key, value string, expirationSec uint64) {
	s.set(key, value, expirationSec)
}

// SetStrNX sets key to hold string value if key does not exist
func (s *Storage) SetStrNX(key, value string, expirationSec uint64) error {
	return s.setNX(key, value, expirationSec)
}

// GetStr gets string value of the key
func (s *Storage) GetStr(key string) (string, error) {
	item := s.get(key)
	if item == nil {
		return "", nil
	}

	if str, ok := item.Value.(string); ok {
		return str, nil
	}
	return "", newErrCustom(errWrongType)
}
