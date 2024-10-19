package fileutils

import (
	"log"
	"os"
)

func ReadTextFile(filename string) (string) {
    content, err := os.ReadFile(filename)

    if err != nil {
        log.Fatal(err)
    }

    return string(content)
}