package config

type PackageConfig map[string]PackageConfigItem

type PackageConfigItem struct {
	// Hidden indicates whether the package is hidden from the pkg.go.dev index.
	// If true, the package will not be listed on pkg.go.dev, and its documentation will not be generated.
	Hidden bool

	// Target specifies the target platform or environment for the package.
	Target string
	// Branch specifies the branch of the repository to use for the package.
	Branch string
}
