package service

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"github.com/rhtyx/bayarind-service.git/controller"
)

func parseError(err error, data string) error {
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		return errors.Join(controller.ErrNotFound, fmt.Errorf(": %s", data))
	case errors.Is(err, gorm.ErrDuplicatedKey):
		return errors.Join(controller.ErrDuplicate, fmt.Errorf(": %s", data))
	default:
		return controller.ErrInternalServer
	}
}
