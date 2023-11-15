package file_core

import (
	"errors"
	"fmt"
)

func Panic(r interface{}) error {
	return errors.New(fmt.Sprintf("Recoverd from panic: %v", r))
}

func Wrapf(err error, format string, args ...interface{}) error {
	return fmt.Errorf(format+": %w", append(args, err)...)
}

func New(error_info string) error {
	return errors.New(error_info)
}

func Newf(error_info string) error {
	return errors.New(error_info)
}

func PanicPrint(r interface{}, error_info string) int {
	fmt.Printf("%s\n", error_info)

	if intValue, ok := r.(int); ok {
		fmt.Printf("Successfully converted to int: %d\n", intValue)
		return intValue
	}
	return 0
}
