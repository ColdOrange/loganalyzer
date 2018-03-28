package loganalyzer

import (
	"fmt"
	"testing"
	"time"
)

func TestAnalyze(t *testing.T) {
	const logFile = "/Users/Orange/Desktop/毕设/log/ssg-surfer/apache-access.log"
	Analyze(logFile)
}
