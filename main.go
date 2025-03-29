package main

import (
	"fmt"
	"os"

	alternatelog "github.com/sirupsen/logrus"
)

/*
Log Levels

Trace
Debug
Info
Warn
Error
Fatal
Panic
*/

func main() {
	file, err := os.OpenFile("E:\\Kode Kloud\\Golang\\advanced-go-code\\go-logging\\app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		fmt.Println(err)
		return
	}
	alternatelog.SetOutput(file) // file is of type *os.File. It has a method func (f *File) WriteTo(w io.Writer) (n int64, err error).
	alternatelog.Println("Hello world")
	alternatelog.Warn("Printing just hello world is waste of time")
	alternatelog.Error("Error occured during execution")
	alternatelog.SetLevel(alternatelog.DebugLevel)
	alternatelog.Debug("System non-responsive")
	alternatelog.SetLevel(alternatelog.TraceLevel)
	alternatelog.Trace("System trace")
	// alternatelog.Fatal("Server process not responding")

}
