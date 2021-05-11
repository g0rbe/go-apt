package apt

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func init() {
	// Set DEBIAN_FRONTEND to "noninteractive" to silent questions
	if os.Getenv("DEBIAN_FRONTEND") != "noninteractive" {
		if err := os.Setenv("DEBIAN_FRONTEND", "noninteractive"); err != nil {
			panic(fmt.Sprintf("failed to set DEBIAN_FRONTEND: %s", err))
		}
	}
}

// Update updates the package cache
func Update() error {

	output, err := exec.Command("/usr/bin/apt", "update", "-y").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run update: %s\noutput: %s", err, output)
	}

	return nil
}

// Upgrade upgrades the packages
func Upgrade() error {

	// Run the external command
	output, err := exec.Command("/usr/bin/apt", "upgrade", "-y").CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run upgrade: %s\noutput: %s", err, output)
	}

	return nil
}

// ListUpgradable lists the available upgradable packages
func ListUpgradable() ([]string, error) {

	output, err := exec.Command("/usr/bin/apt", "list", "--upgradable").CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("failed to list upgradable: %s\noutput: %s", err, output)
	}

	packages := make([]string, 0)

	for _, line := range strings.Split(string(output), "\n") {

		switch true {
		case line == "":
			continue
		case strings.Contains(line, "apt does not have a stable CLI interface"):
			continue
		case strings.Contains(line, "Listing..."):
			continue
		}

		packages = append(packages, strings.Split(line, "/")[0])
	}

	return packages, nil
}

// Install install the given package
func Install(pkg string) error {

	output, err := exec.Command("/usr/bin/apt", "-y", "install", pkg).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run update: %s\noutput: %s", err, output)
	}

	return nil
}

// Remove remove teh given package
func Remove(pkg string) error {

	output, err := exec.Command("/usr/bin/apt", "-y", "remove", pkg).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run update: %s\noutput: %s", err, output)
	}

	return nil
}

// Purge purge (remove the config files too) the package
func Purge(pkg string) error {

	output, err := exec.Command("/usr/bin/apt", "-y", "purge", pkg).CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to run update: %s\noutput: %s", err, output)
	}

	return nil
}
