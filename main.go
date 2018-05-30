package main

import (
	"flag"

	"github.com/ColdOrange/loganalyzer/loganalyzer"
	log "github.com/ColdOrange/loganalyzer/loganalyzer/logging"
)

func init() {
	debug := flag.Bool("debug", false, "Print debug level logs")
	dev := flag.Bool("dev", false, "Run in development mode (with webpack watcher)")
	flag.Parse()
	if *debug {
		log.SetLevel("Debug")
	}
	if *dev {
		go loganalyzer.StartWebpackWatcher()
	}
}

func main() {
	server := loganalyzer.NewServer("127.0.0.1:8080")
	server.ListenAndServe()
}
