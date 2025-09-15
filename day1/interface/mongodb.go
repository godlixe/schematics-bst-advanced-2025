package main

import (
	"errors"
	"fmt"
)

type MongoAdapter struct {
	store map[int]string
}

func NewMongoAdapter() *MongoAdapter {
	return &MongoAdapter{store: make(map[int]string)}
}

func (m *MongoAdapter) GetData(id int) (string, error) {
	val, ok := m.store[id]
	if !ok {
		return "", errors.New("record not found")
	}
	return val, nil
}

func (m *MongoAdapter) CreateData(data string) (string, error) {
	id := len(m.store) + 1
	m.store[id] = data
	return fmt.Sprintf("created with id %d", id), nil
}

func (m *MongoAdapter) UpdateData(id int, data string) (string, error) {
	if _, ok := m.store[id]; !ok {
		return "", errors.New("record not found")
	}
	m.store[id] = data
	return fmt.Sprintf("updated id %d", id), nil
}

func (m *MongoAdapter) DeleteData(id int) error {
	if _, ok := m.store[id]; !ok {
		return errors.New("record not found")
	}
	delete(m.store, id)
	return nil
}
