package advicer

import (
	"fmt"
	"os"
)

func Error(line string) {
	fmt.Println(line)
	os.Exit(1)
}
