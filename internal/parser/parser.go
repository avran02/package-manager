package parser

import (
	"path/filepath"

	"github.com/avran02/package-manager/internal/models"
)

func ParcePackageFile(path string) (models.PackageFile, error) {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parceYamlPackageFile(path)
	case ".json":
		return parceJsonPackageFile(path)
	default:
		return models.PackageFile{}, ErrUnknownAFileExt
	}
}

func ParseDownloadPkgsFile(path string) (models.DownloadPackagesFile, error) {
	switch filepath.Ext(path) {
	case ".yaml", ".yml":
		return parceYamlDownloadPkgsFile(path)
	case ".json":
		return parceJsonDownloadPkgsFile(path)
	default:
		return models.DownloadPackagesFile{}, ErrUnknownAFileExt
	}
}
