package parser

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/avran02/package-manager/internal/models"
)

func parceJsonPackageFile(path string) (models.PackageFile, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return models.PackageFile{}, fmt.Errorf("Error opening file %s: %w", path, err)
	}
	defer jsonFile.Close()

	var pkgFile models.PackageFile
	if err := json.NewDecoder(jsonFile).Decode(&pkgFile); err != nil {
		return models.PackageFile{}, fmt.Errorf("Error unmarshal file %s: %w", path, err)
	}
	return pkgFile, nil
}

func parceJsonDownloadPkgsFile(path string) (models.DownloadPackagesFile, error) {
	jsonFile, err := os.Open(path)
	if err != nil {
		return models.DownloadPackagesFile{}, fmt.Errorf("Error opening file %s: %d", path, err)
	}
	defer jsonFile.Close()

	var dpFile models.DownloadPackagesFile

	if err := json.NewDecoder(jsonFile).Decode(&dpFile); err != nil {
		return models.DownloadPackagesFile{}, fmt.Errorf("Error unmarshal file %s: %w", path, err)
	}

	return dpFile, nil
}
