package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/swlee3306/osbeat/beater"
)

func main() {
	err := beat.Run("osbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
