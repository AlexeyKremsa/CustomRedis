package storage

import "errors"

type Str struct {
	Value string
}

type Storage struct {
	keyValues map[string]interface{}
}

func Init() *Storage {
	strg := &Storage{}
	strg.keyValues = make(map[string]interface{})
	return strg
}

func (s *Storage) SetStr(key, value string) {
	item := Str{Value: value}
	s.keyValues[key] = item
}

func (s *Storage) SetStrNX(key, value string) error {
	if _, ok := s.keyValues[key]; ok {
		return newErrCustom("Key already exists")
	}

	item := Str{Value: value}
	s.keyValues[key] = item

	return nil
}

func (s *Storage) GetStr(key string) (interface{}, error) {
	if v, ok := s.keyValues[key]; ok {
		if s, ok := v.(Str); ok {
			return s.Value, nil
		}
		return nil, errors.New(errWrongType)
	}

	return nil, nil
}
