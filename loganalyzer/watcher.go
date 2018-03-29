package loganalyzer

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"path"
	"path/filepath"

	log "loganalyzer/loganalyzer/logging"
)

func StartNodeWatcher() {
	webpackPath, _ := filepath.Abs(path.Join(ProjectPath, "asserts/node_modules/webpack/bin/webpack.js"))
	log.Debugln("Webpack path:", webpackPath)
	cmd := exec.Command("node", webpackPath, "--mode", "development", "--colors")
	cmd.Dir, _ = filepath.Abs(path.Join(ProjectPath, "asserts/")) // Webpack should exec under `asserts` directory

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("Start node watcher error:", err)
	}
	err = cmd.Start()
	if err != nil {
		log.Fatalln("Start node watcher error:", err)
	}
	log.Infoln("Webpack start watching files")

	reader := bufio.NewReader(stdout)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Fatalln("Start node watcher error:", err)
			}
			break
		}
		fmt.Print(line)
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatalln("Start node watcher error:", err)
	}
}
