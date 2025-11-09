package middleware

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Store the detected distro
var cachedDistro string
var distroErr error

func InitDistr() error {
	cachedDistro, distroErr = DetectDistro()
	if distroErr != nil {
		return distroErr
	}
	fmt.Println("Detected Distro:", cachedDistro)
	return nil
}

// use this everywhere to get the distro
func GetDistr() (string, error) {
	return cachedDistro, distroErr
}

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
