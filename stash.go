package stash

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

const namespace string = "orbital-core-auth-"

// Entity object
type entity struct {
	Message  string   `json:"message"`
	Level    string   `json:"level"`
	Function string   `json:"function"`
	Details  *Details `json:"details"`
}

// Details object
type Details struct {
	TransactionID string `json:"transaction_id,omitempty"`
	RoundID       string `json:"round_id,omitempty"`
	SessionID     string `json:"session_id,omitempty"`
	UserID        string `json:"user_id,omitempty"`
	GameID        int    `json:"game_id,omitempty"`
	CurrencyID    int    `json:"currency_id,omitempty"`
}

// Debug logs on debug level
func Debug(msg string, d *Details) {
	e := entity{
		Message:  msg,
		Level:    "DEBUG",
		Function: "test",
		Details:  d,
	}
	logE(e)
}

// Info logs on info level
func Info(msg string, d *Details) {
	e := entity{
		Message:  msg,
		Level:    "INFO",
		Function: "test",
		Details:  d,
	}
	logE(e)
}

// Warn logs on warn level
func Warn(msg string, d *Details) {
	e := entity{
		Message:  msg,
		Level:    "WARN",
		Function: "test",
		Details:  d,
	}
	logE(e)
}

// Error logs on error level
func Error(msg string, d *Details) {
	e := entity{
		Message:  msg,
		Level:    "ERROR",
		Function: "test",
		Details:  d,
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
	path := os.Getenv("LOG_PATH")
	b = append(b, "\n"...)
	filename := fmt.Sprintf("%s/%s%s", path, namespace, name)
	fmt.Printf("append: %v", filename)
	var f *os.File
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		f, err = os.Create(filename)
	} else {
		f, err = os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			return err
		}
	}
	defer f.Close()
	_, err = f.Write(b)
	return err
}

// File returns file to log in request/response
func File(name string) (io.WriteCloser, error) {
	path := os.Getenv("LOG_PATH")
	filename := fmt.Sprintf("%s/%s%s", path, namespace, name)
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
