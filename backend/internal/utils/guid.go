package utils

import "github.com/google/uuid"

type GUIDGenerator struct{}

func NewGUIDGenerator() *GUIDGenerator {
	return &GUIDGenerator{}
}

func (g *GUIDGenerator) GenerateGUID() string {
	uuidObj := uuid.New()
	return uuidObj.String()
}
