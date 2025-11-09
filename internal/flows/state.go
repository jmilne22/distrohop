package flows

import (
	"fmt"
	"github.com/jmilne22/distrohop/internal/config"
	"github.com/jmilne22/distrohop/internal/package_managers"
)

func State(packageNames ...string) error {
	// Step 1: Load config
	if err := config.LoadConfig(); err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	// Step 2: Get package manager
	pm, err := package_managers.GetPackageManager()
	if err != nil {
		return fmt.Errorf("failed to get package manager: %w", err)
	}

	// Step 3: Get installed packages
	installed, err := pm.List()
	if err != nil {
		return fmt.Errorf("failed to list installed packages: %w", err)
	}

	// Step 4: Convert installed to map for fast lookup
	installedMap := make(map[string]bool)
	for _, pkg := range installed {
		installedMap[pkg] = true
	}

	// Step 5: Find missing packages (in config but not installed)
	var missing []string
	for _, pkg := range config.Cfg.Distro.Packages {
		if !installedMap[pkg] {
			missing = append(missing, pkg)
		}
	}
	// Step 6: Print results (ADD THIS SECTION)
	fmt.Printf("Total installed packages: %d\n", len(installed))
	fmt.Printf("Packages in config: %d\n", len(config.Cfg.Distro.Packages))

	if len(missing) > 0 {
		fmt.Printf("\nMissing packages (%d):\n", len(missing))
		for _, pkg := range missing {
			fmt.Printf("  - %s\n", pkg)
		}
	} else {
		fmt.Println("\n✓ All config packages are installed!")
	}

	return nil // Don't forget this!

}
