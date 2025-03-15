package protocol

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	s := id.String()
	s = strings.Replace(s, "-", "", -1)
	return s, nil
}
