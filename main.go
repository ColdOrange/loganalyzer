package main

import (
	"flag"

	"loganalyzer/loganalyzer"
	log "loganalyzer/loganalyzer/logging"
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
	loganalyzer.Analyze()
	server := loganalyzer.NewServer("127.0.0.1:8080")
	server.ListenAndServe()
}
