package main

import (
	"loganalyzer/loganalyzer"
)

func main() {
	loganalyzer.Analyze()
	//go loganalyzer.StartNodeWatcher() // TODO: fix webpack watcher
	server := loganalyzer.NewServer("127.0.0.1:8080")
	server.ListenAndServe()
}
