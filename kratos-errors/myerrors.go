package main

import (
	"fmt"

	"github.com/superproj/superproj-examples/kratos-errors/myerrors"
)

func main() {
	notFoundErr := myerrors.ErrorUserNotFound("user %s not found", "colin")
	fmt.Printf("ErrorUserNotFound: %q\n", notFoundErr)

	if myerrors.IsUserNotFound(notFoundErr) {
		fmt.Printf("IsUserNotFound: %v\n", true)
	}
}
