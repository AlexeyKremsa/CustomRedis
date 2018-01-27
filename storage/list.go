package storage

// SetList sets key to hold the string array value
func (s *Storage) SetList(key string, value []string, expirationSec uint64) {
	s.set(key, value, expirationSec)
}

// GetList returns string array accroding to the key
func (s *Storage) GetList(key string) ([]string, error) {
	item := s.get(key)
	if item == nil {
		return nil, nil
	}

	if list, ok := item.Value.([]string); ok {
		return list, nil
	}

	return nil, newErrCustom(errWrongType)
}

// ListInsert adds string elements to the end of the array stored by the specified key and returns array`s length
func (s *Storage) ListInsert(key string, items []string) (int, error) {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	if item, ok := shard.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return 0, nil
		}

		if list, ok := item.Value.([]string); ok {
			list = append(list, items...)
			item.Value = list
			shard.keyValues[key] = item
			return len(list), nil
		}
		return 0, newErrCustom(errWrongType)
	}
	return 0, nil
}

// ListPop removes and returns the last string element of the array stored by the specified key
func (s *Storage) ListPop(key string) (string, error) {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	if item, ok := shard.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return "", nil
		}

		if list, ok := item.Value.([]string); ok {
			lastElem := list[len(list)-1]
			item.Value = list[:len(list)-1]
			shard.keyValues[key] = item
			return lastElem, nil
		}
		return "", newErrCustom(errWrongType)
	}
	return "", nil
}

// ListIndex returns array element at specified index
func (s *Storage) ListIndex(key string, index int) (string, error) {
	shard := s.getShard(key)

	shard.mutex.Lock()
	defer shard.mutex.Unlock()

	if item, ok := shard.keyValues[key]; ok {
		if isExpired(item.Expiration) {
			return "", nil
		}

		if list, ok := item.Value.([]string); ok {
			if index >= len(list) || index < 0 {
				return "", newErrCustom(errIndexOutOfRange)
			}
			return list[index], nil
		}
		return "", newErrCustom(errWrongType)
	}
	return "", nil
}
