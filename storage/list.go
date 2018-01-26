package storage

// SetList sets key to hold the string array value
func (s *Storage) SetList(key string, value []string, expirationSec int64) {
	s.set(key, value, expirationSec)
}

// GetList returns string array accroding to the key
func (s *Storage) GetList(key string) ([]string, error) {
	val := s.get(key)

	if list, ok := val.([]string); ok {
		return list, nil
	}

	return nil, newErrCustom(errWrongType)
}

// ListPush adds string elements to the end of the array stored by the specified key
func (s *Storage) ListPush(key string, items []string) error {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	if item, ok := shard.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return newErrCustom(errNotExist)
		}

		if list, ok := item.Value.([]string); ok {
			list = append(list, items...)
			item.Value = list
			shard.keyValues[key] = item
			return nil
		}
		return newErrCustom(errWrongType)
	}
	return newErrCustom(errNotExist)
}

// ListPop returns the last string element of the array stored by the specified key
func (s *Storage) ListPop(key string) (string, error) {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	if item, ok := shard.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return "", newErrCustom(errNotExist)
		}

		if list, ok := item.Value.([]string); ok {
			lastElem := list[len(list)-1]
			item.Value = list[:len(list)-1]
			shard.keyValues[key] = item
			return lastElem, nil
		}
		return "", newErrCustom(errWrongType)
	}
	return "", newErrCustom(errNotExist)
}

// ListIndex returns array element at specified index
func (s *Storage) ListIndex(key string, index int) (string, error) {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	if item, ok := shard.keyValues[key]; ok {
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
