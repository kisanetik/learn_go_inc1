package logger

import (
	"fmt"
	"log"
)

func Fatalf(str string, err error) {
	log.Fatalf("FATAL: "+str, err)
}

func Errorf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("ERROR: %s\n", msg)
}
