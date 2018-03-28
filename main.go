package main

import (
	"loganalyzer/loganalyzer"
)

func main() {
	const logFile = "/Users/Orange/Desktop/毕设/log/ssg-surfer/apache-access.log"
	loganalyzer.Analyze(logFile)
	server := loganalyzer.NewServer("127.0.0.1:8080")
	server.ListenAndServe()
}
