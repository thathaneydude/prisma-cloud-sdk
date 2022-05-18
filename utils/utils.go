package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/thathaneydude/prisma-cloud-sdk/internal"
	"io"
	"net/http"
)

func ToBytes[T any](obj T) []byte {
	ret, err := json.Marshal(obj)
	if err != nil {
		logrus.Errorf("Failed to convert object to bytes: %v", err)
	}
	return ret
}

func UnmarshalResponse(httpResponse *http.Response, response interface{}) error {
	if httpResponse == nil {
		return &internal.GenericError{Msg: fmt.Sprintf("Error while reading response: No data found")}
	}

	defer httpResponse.Body.Close()
	tmp, readErr := io.ReadAll(httpResponse.Body)
	if readErr != nil {
		return &internal.GenericError{Msg: fmt.Sprintf("Error while reading response body: %v", readErr)}
	}
	unmarshalErr := json.Unmarshal(tmp, response)
	if unmarshalErr != nil {
		return &internal.GenericError{Msg: fmt.Sprintf("Error while unmarshaling response: %v", unmarshalErr)}
	}
	return nil
}
