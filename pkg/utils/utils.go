package utils

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
)

func ToBytes[T any](obj T) []byte {
	ret, err := json.Marshal(obj)
	if err != nil {
		logrus.Errorf("Failed to convert object to bytes: %v", err)
	}
	return ret
}
