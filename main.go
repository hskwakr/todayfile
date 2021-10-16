package main

import (
	"log"
	"os"
	"time"

	"github.com/hskwakr/todayfile/todayfile"
)

func main() {
	if err := todayfile.Create(todayfile.Date(time.Now()) + ".txt"); err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
