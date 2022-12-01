package check

import (
	"fmt"
	"log"
)

// IfError prints error, if it is not nil
func IfError(err error) {
	if err == nil {
		return
	}

	log.Printf("\x1b[31;1m%s\x1b[0m\n", fmt.Sprintf("ERROR: %s", err))
}
