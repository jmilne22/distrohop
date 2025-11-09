package package_managers

import (
	"fmt"
	"os/exec"
	"strings"
)

type Xbps struct{}

func (x *Xbps) List() ([]string, error) {
	cmd := exec.Command("xpkg", "-m")

	// capture the output
	output, err := cmd.Output()

	if err != nil {
		return nil, fmt.Errorf("failed to run xpkg -m %w", err)
	}

	// convert output to string and split by newlines
	outputStr := string(output)
	lines := strings.Split(strings.TrimSpace(outputStr), "\n")

	// filter out empty lines if there are any
	var packages []string
	for _, line := range lines {
		if line != "" {
			packages = append(packages, line)
		}
	}

	return packages, nil
}

func (x *Xbps) Install(packageName ...string) error {
	fmt.Println("placeholder")
	return nil
}

func (x *Xbps) Remove(packageName ...string) error {
	fmt.Println("placeholder")
	return nil
}
