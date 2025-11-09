package package_managers

type PackageManager interface {
	Install(packageName ...string) error
	Remove(packageName ...string) error
	List() ([]string, error)
}
