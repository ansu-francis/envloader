package envloader

import (
	"bufio"
	"io"
	"os"
	"strings"
)

func LoadEnv(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	envMap, err := parseFile(file)
	if err != nil {
		return err
	}

	if err = load(envMap); err != nil {
		return err
	}
	return nil
}

func load(envMap map[string]string) error {
	for key, value := range envMap {
		os.Setenv(key, value)
	}
	return nil
}

func parseFile(r io.Reader) (map[string]string, error) {
	envMap := make(map[string]string)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "=")
		if len(parts) != 2 {
			continue
		}
		envMap[strings.TrimSpace(parts[0])] = strings.TrimSpace(strings.Trim(parts[1], "\""))
	}
	if err := scanner.Err(); err != nil {
		return envMap, err
	}
	return envMap, nil
}
