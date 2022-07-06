package internal

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// ToBytes is a generic utility which converts a struct to json bytes.
//func ToBytes[T any](obj T) []byte {
//	ret, err := json.Marshal(obj)
//	if err != nil {
//		logrus.Errorf("Failed to convert object to bytes: %v", err)
//	}
//	return ret
//}

// UnmarshalResponse is a generic utility that reads a http response and attempts to unmarshal it to whatever
// struct (interface) is provided
func UnmarshalResponse(httpResponse *http.Response, response interface{}) error {
	if httpResponse == nil {
		return &GenericError{Msg: fmt.Sprintf("Error while reading response: No data found")}
	}

	defer httpResponse.Body.Close()
	tmp, readErr := io.ReadAll(httpResponse.Body)
	if readErr != nil {
		return &GenericError{Msg: fmt.Sprintf("Error while reading response body: %v", readErr)}
	}
	unmarshalErr := json.Unmarshal(tmp, response)
	if unmarshalErr != nil {
		return &GenericError{Msg: fmt.Sprintf("Error while unmarshaling response: %v", unmarshalErr)}
	}
	return nil
}
