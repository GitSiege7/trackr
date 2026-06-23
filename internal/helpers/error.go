package helpers

import "fmt"

// Tests err, if not nil prints error output
func HandlerError(err error, msg string) {
	if err != nil {
		fmt.Printf("ERROR: %v: %v\n", msg, err)
	}
}
