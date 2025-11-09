package package_managers

import (
	"fmt"
	"github.com/jmilne22/distrohop/internal/middleware"
)

func GetPackageManager() (PackageManager, error) {
	distro, err := middleware.DetectDistro()
	if err != nil {
		return nil, err
	}

	switch distro {
	case "void":
		return &Xbps{}, nil
	default:
		return nil, fmt.Errorf("Package maanger not detected %s", distro)
	}
}
