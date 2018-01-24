package storage

func (s *Storage) SetList(key string, value []string, expirationSec int64) {
	s.set(key, value, expirationSec)
}

func (s *Storage) GetList(key string) ([]string, error) {
	val, err := s.get(key)
	if err != nil {
		return nil, err
	}

	if list, ok := val.([]string); ok {
		return list, nil
	}

	return nil, newErrCustom(errWrongType)
}

func (s *Storage) ListPush(key string, items []string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if item, ok := s.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return newErrCustom(errNotExist)
		}

		if list, ok := item.Value.([]string); ok {
			list = append(list, items...)
			item.Value = list
			s.keyValues[key] = item
			return nil
		}
		return newErrCustom(errWrongType)
	}
	return newErrCustom(errNotExist)
}

func (s *Storage) ListPop(key string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if item, ok := s.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return "", newErrCustom(errNotExist)
		}

		if list, ok := item.Value.([]string); ok {
			lastElem := list[len(list)-1]
			item.Value = list[:len(list)-1]
			s.keyValues[key] = item
			return lastElem, nil
		}
		return "", newErrCustom(errWrongType)
	}
	return "", newErrCustom(errNotExist)
}

func (s *Storage) ListIndex(key string, index int) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if item, ok := s.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return "", newErrCustom(errNotExist)
		}

		if list, ok := item.Value.([]string); ok {
			if index >= len(list) || index < 0 {
				return "", newErrCustom(errIndexOutOfRange)
			}
			return list[index], nil
		}
		return "", newErrCustom(errWrongType)
	}
	return "", newErrCustom(errNotExist)
}
