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

func StartWebpackWatcher() {
	webpackPath, _ := filepath.Abs(path.Join(ProjectPath, "assets/node_modules/webpack/bin/webpack.js"))
	log.Debugln("Resolve webpack path:", webpackPath)
	cmd := exec.Command("node", webpackPath, "--mode", "development", "--watch", "--colors")
	cmd.Dir, _ = filepath.Abs(path.Join(ProjectPath, "assets")) // Webpack should exec under `assets` directory

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatalln("Start webpack watcher error:", err)
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatalln("Start webpack watcher error:", err)
	}

	err = cmd.Start()
	if err != nil {
		log.Fatalln("Start webpack watcher error:", err)
	}
	log.Infoln("Webpack start watching files")

	go copyOutput(stdout)
	go copyOutput(stderr)

	err = cmd.Wait()
	if err != nil {
		log.Fatalln("Start webpack watcher error:", err)
	}
}

func copyOutput(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
