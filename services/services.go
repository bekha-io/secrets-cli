package services

import (
	"fmt"
)

func NewError(text string, a ...interface{}) error {
	return fmt.Errorf(text, a...)
}

type serviceManager struct {
	groupsService
	secretsService
}

var Services = &serviceManager{}