package section1

import (
	"log"
	"os"
)

type Logger interface {
	Log(message string)
}

type FileLogger struct {
	FileName string
}

func (f *FileLogger) Log(message string) {
	os.WriteFile(f.FileName, []byte(message), os.ModePerm)
}

type ConsoleLogger struct {
}

func (f *ConsoleLogger) Log(message string) {
	log.Println(message)
}
