package stash

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

var namespace, path string

// New sets namespace for the file
func New(p, ns string) {
	namespace = ns
	path = p
}

// Status type for status
type Status int

// Status possible values
const (
	StatusSuccess Status = iota + 1
	StatusGameInternal
	StatusCasinoHTTPError
	StatusGameBadRequest
	StatusCoreBadRequest
	StatusCasinoBadResponse
)

// ObjectLog type for storing object
type ObjectLog struct {
	Object      interface{} `json:"object"`
	Status      Status      `json:"status"`
	Timestamp   time.Time   `json:"timestamp"`
	IsProcessed bool        `json:"is_processed"`
}

// Object log into file
func Object(name string, t interface{}, status Status) {
	tl := ObjectLog{Object: t, Status: status, Timestamp: time.Now()}
	b, _ := json.Marshal(tl) // Imedia arasdros moxdeba aq error
	err := myappend(name, b)
	if err != nil {
		log.Fatalf("could not append object: %v", err)
	}
}

// Entity object
type entity struct {
	Message  string `json:"message"`
	Level    string `json:"level"`
	Function string `json:"function"`
}

// Debug logs on debug level
func Debug(msg string) {
	e := entity{
		Message:  msg,
		Level:    "DEBUG",
		Function: "test",
	}
	logE(e)
}

// Info logs on info level
func Info(msg string) {
	e := entity{
		Message:  msg,
		Level:    "INFO",
		Function: "test",
	}
	logE(e)
}

// Warn logs on warn level
func Warn(msg string) {
	e := entity{
		Message:  msg,
		Level:    "WARN",
		Function: "test",
	}
	logE(e)
}

// Error logs on error level
func Error(msg string) {
	e := entity{
		Message:  msg,
		Level:    "ERROR",
		Function: "test",
	}
	logE(e)
}

func logE(e entity) {
	b, _ := json.Marshal(e)
	err := myappend("log", b)
	if err != nil {
		log.Printf("could not append log: %v", err)
	}
}

func myappend(name string, b []byte) (err error) {
	b = append(b, "\n"...)
	f, err := File(name)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(b)
	return err
}

// File returns file to log in request/response
func File(name string) (io.WriteCloser, error) {
	filename := fmt.Sprintf("%s/%s%s.log", path, namespace, name)
	var f *os.File
	var err error
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, err = os.Create(filename)
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return nil, err
		}
	}
	return f, err
}
