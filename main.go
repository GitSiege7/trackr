package main

import (
	"fmt"
	"os"
)

// Prints error output and exits
func handlerError(err error) {
	if err != nil {
		fmt.Println("ERROR: ", err)
	}
	os.Exit(1)
}

func main() {
	const DB_PATH = "~/.config/trackr/times.db"

}
