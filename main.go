package main

import (
	"flag"

	"loganalyzer/loganalyzer"
	log "loganalyzer/loganalyzer/logging"
)

func init() {
	debug := flag.Bool("d", false, "Print debug level logs")
	flag.Parse()
	if *debug {
		log.SetLevel("Debug")
	}
}

func main() {
	loganalyzer.Analyze()
	//go loganalyzer.StartNodeWatcher() // TODO: fix webpack watcher
	server := loganalyzer.NewServer("127.0.0.1:8080")
	server.ListenAndServe()
}
