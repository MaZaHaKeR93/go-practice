package main

import (
	"fmt"
	"log"
	"os"
)

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println(message)
}

type FileLogger struct {
	file *os.File
}

func NewFileLogger(filename string) (*FileLogger, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &FileLogger{file: f}, nil
}

func (f FileLogger) Log(message string) {
	_, err := f.file.WriteString(message + "\n")
	if err != nil {
		log.Println("Error writing to file", err)
	}
}

func (f *FileLogger) Close() {
	f.file.Close()
}

func DoLogging(l Logger, message string) {
	l.Log(message)
}

func main() {
	cl := ConsoleLogger{}
	DoLogging(cl, "Message in console")

	fl, err := NewFileLogger("log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer fl.Close()

	DoLogging(fl, "Message in the file")
	DoLogging(fl, "One more message in the file")
}
