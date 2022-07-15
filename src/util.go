package src

import (
	"fmt"
	"os"
)

func Fatal(buf string, exitCode int) {
	fmt.Printf("%s\n", buf)
	os.Exit(exitCode)
}
