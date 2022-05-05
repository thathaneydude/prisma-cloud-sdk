package pkg

import (
	"github.com/sirupsen/logrus"
)

type GenericError struct {
	Msg string
}

func (m *GenericError) Error() string {
	logrus.Errorf(m.Msg)
	return m.Msg
}
