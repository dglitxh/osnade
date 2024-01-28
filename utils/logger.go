package utils

import (
	"log"
	"os"

	"log/slog"
)

func CreateLogFile(name string) *os.File {
	f, err := os.OpenFile(name+".txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

var fn *os.File = CreateLogFile("log")

func Logger(txt string) {

	killog := log.New(fn, "", 1)
	killog.SetOutput(fn)
	killog.SetFlags(log.LstdFlags | log.Lshortfile)
	killog.Println(txt)
	log.Println(txt)

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil)
	myslog := slog.New(jsonHandler)
	myslog.Info("hi there")

}
