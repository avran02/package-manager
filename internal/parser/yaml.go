package parser

import (
	"fmt"
	"os"

	"github.com/avran02/package-manager/internal/models"
	"gopkg.in/yaml.v3"
)

func parceYamlPackageFile(path string) (models.PackageFile, error) {
	yamlFile, err := os.Open(path)
	if err != nil {
		return models.PackageFile{}, fmt.Errorf("Error opening file %s: %w", path, err)
	}
	defer yamlFile.Close()

	var pkgFile models.PackageFile
	if err := yaml.NewEncoder(yamlFile).Encode(pkgFile); err != nil {
		return models.PackageFile{}, fmt.Errorf("Error unmarshal file %s: %w", path, err)
	}

	return pkgFile, nil
}

func parceYamlDownloadPkgsFile(path string) (models.DownloadPackagesFile, error) {
	yamlFile, err := os.Open(path)
	if err != nil {
		return models.DownloadPackagesFile{}, fmt.Errorf("Error opening file %s: %d", path, err)
	}
	defer yamlFile.Close()

	var dpFile models.DownloadPackagesFile

	if err := yaml.NewEncoder(yamlFile).Encode(dpFile); err != nil {
		return models.DownloadPackagesFile{}, fmt.Errorf("Error unmarshal file %s: %w", path, err)
	}

	return dpFile, nil
}
