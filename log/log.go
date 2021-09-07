package log

import (
	"log"
	"os"
)

func New() *log.Logger {
	return log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
