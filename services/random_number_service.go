package services

import "math/rand"

// A service providing a random number
type RandomNumberService interface {
	GenerateRandomNumber() int
}

// Default implementation of RandomNumberService
type DefaultRandomNumberService struct {
	Seed int64
}

// Generates a random number
func (rns *DefaultRandomNumberService) GenerateRandomNumber() int {
	rand.Seed(rns.Seed)
	return rand.Int()
}
