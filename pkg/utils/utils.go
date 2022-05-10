package utils

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"prisma-cloud-sdk/pkg"
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
		return &pkg.GenericError{Msg: fmt.Sprintf("Error while reading response: No data found")}
	}

	defer httpResponse.Body.Close()
	tmp, readErr := io.ReadAll(httpResponse.Body)
	logrus.Debugf("Response: %v", string(tmp))
	if readErr != nil {
		return &pkg.GenericError{Msg: fmt.Sprintf("Error while reading response body: %v", readErr)}
	}
	unmarshalErr := json.Unmarshal(tmp, response)
	if unmarshalErr != nil {
		return &pkg.GenericError{Msg: fmt.Sprintf("Error while unmarshaling response: %v", unmarshalErr)}
	}
	return nil
}
