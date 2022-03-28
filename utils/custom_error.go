package utils

import (
	"errors"
	"fmt"
)

type AppError struct {
	ErrorCode string
	Err       error
}

func (e *AppError) Error() string {
	return fmt.Sprintf("code : %s, err : %s", e.ErrorCode, e.Err)
}

func RequiredError() error {
	return &AppError{
		ErrorCode: "X01",
		Err:       errors.New("Input cant be empty"),
	}
}
func UnauthorizedError() error {
	return &AppError{
		ErrorCode: "X04",
		Err:       errors.New("Unauthorized user"),
	}
}
func DataNotFoundError() error {
	return &AppError{
		ErrorCode: "X02",
		Err:       errors.New("No Data Found"),
	}
}
