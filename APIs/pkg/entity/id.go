package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func ParseID(id string) (ID, error) {
	newId, err := uuid.Parse(id)
	return newId, err
}
