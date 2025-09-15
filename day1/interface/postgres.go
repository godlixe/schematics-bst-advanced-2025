package main

import (
	"errors"
	"fmt"
)

type PostgresAdapter struct {
	store map[int]string
}

func NewPostgresAdapter() *PostgresAdapter {
	return &PostgresAdapter{store: make(map[int]string)}
}

func (p *PostgresAdapter) GetData(id int) (string, error) {
	val, ok := p.store[id]
	if !ok {
		return "", errors.New("record not found")
	}
	return val, nil
}

func (p *PostgresAdapter) CreateData(data string) (string, error) {
	id := len(p.store) + 1
	p.store[id] = data
	return fmt.Sprintf("created with id %d", id), nil
}

func (p *PostgresAdapter) UpdateData(id int, data string) (string, error) {
	if _, ok := p.store[id]; !ok {
		return "", errors.New("record not found")
	}
	p.store[id] = data
	return fmt.Sprintf("updated id %d", id), nil
}

func (p *PostgresAdapter) DeleteData(id int) error {
	if _, ok := p.store[id]; !ok {
		return errors.New("record not found")
	}
	delete(p.store, id)
	return nil
}
