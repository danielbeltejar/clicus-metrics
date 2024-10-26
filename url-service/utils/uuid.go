package utils

import "github.com/google/uuid"

func GenerateShortID() string {
	return uuid.New().String()[:8]
}
