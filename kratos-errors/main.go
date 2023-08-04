package main

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
)

var (
	// Token invalidated error.
	tokenErr = errors.New(400, "InvalidParameter", "Token has been invalidated")
)

func main() {
	// new error examples
	if err := errors.New(400, "InvalidParameter", "Token has been invalidated"); err != nil {
		fmt.Printf("%-28sError is: %q\n", "[new error]", err)
	}

	if err := errors.New(400, "InvalidParameter", "Token has been invalidated").WithMetadata(map[string]string{"token": "xxx"}); err != nil {
		fmt.Printf("%-28sError is: %q\n", "[new error with metadata]", err)
	}

	if err := errors.Newf(400, "InvalidParameter", "Token %s has been invalidated", "xxx"); err != nil {
		fmt.Printf("%-28sError is: %q\n", "[formatting creation error]", err)
	}

	if err := errors.FromError(fmt.Errorf("new error with fmt.Errorf")); err != nil {
		fmt.Printf("%-28sError is: %+v\n", "[FromError example]", err)
	}

	// output error examples
	fmt.Printf("%-28sCode: %d | Reason: %s | Message: %s | Metadata: %v\n",
		"[output error]",
		errors.Code(tokenErr),
		errors.Reason(tokenErr),
		tokenErr.GetMessage(),
		tokenErr.GetMetadata(),
	)

	// predefined error examples
	unauthorizedErr := errors.Unauthorized("Unauthorized", "No permission to delete clusters")
	if unauthorizedErr != nil {
		fmt.Printf("%-28sError is: %q\n", "[predefined error]", unauthorizedErr)
	}
	if errors.IsUnauthorized(unauthorizedErr) {
		fmt.Printf("%-28sIs unauthorized: %v\n", "[predefined error]", true)
	}

	// error suger examples
	if err := withCause(); err != nil {
		fmt.Printf("%-28s`err` is `tokenErr`: %v\n", "[error suger]", errors.Is(err, tokenErr))
	}
	var asErr *errors.Error
	if errors.As(tokenErr, &asErr) {
		// Hint: return the first error of type errors.Error in asErr.
		fmt.Printf("%-28sFirst error is: %v\n", "[error suger]", asErr)
	}
}

// withCause with the underlying cause of the error.
func withCause() error {
	return errors.InternalServer("InternalServer", "Internal call exception").WithCause(tokenErr)
}
