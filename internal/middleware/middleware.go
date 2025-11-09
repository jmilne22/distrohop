package middleware

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func DetectDistro() (string, error) {
	path := "/etc/os-release"
	fileHandle, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer fileHandle.Close()

	scanner := bufio.NewScanner(fileHandle)

	for scanner.Scan() {
		textLine := scanner.Text()

		// skip empty lines
		if strings.TrimSpace(textLine) == "" || strings.HasPrefix(textLine, "#") {
			continue
		}

		// extract the distro name
		if distroName, found := strings.CutPrefix(textLine, "ID="); found {
			distroName = strings.Trim(distroName, `"`)
			return distroName, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	return "", fmt.Errorf("ID field not found in %s", path)
}
