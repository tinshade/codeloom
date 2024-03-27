package helpers

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func RegisterEnvVars(logfile string) (map[string]string, error) {
	f, err := os.OpenFile(logfile, os.O_RDONLY, os.ModePerm)
	dict := map[string]string{}
	if err != nil {
		log.Fatalf("open file error: %v", err)
		return dict, err
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	for sc.Scan() {
		var line string = sc.Text() // GET the line string
		if len(line) > 0 {
			tempEnv := strings.Split(line, "=")
			key := tempEnv[0]
			value := tempEnv[1]
			dict[key] = value
			os.Setenv(key, value)
		}
	}

	if err := sc.Err(); err != nil {
		log.Fatalf("scan file error: %v", err)
		return dict, err
	}
	return dict, nil
}
