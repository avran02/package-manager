package models

import (
	"path/filepath"
)

type PackageFile struct {
	PackageInfo
	Targets      []Target      `json:"targets,omitempty" yaml:"targets,omitempty"`
	Dependencies []PackageInfo `json:"deps,omitempty" yaml:"deps,omitempty"`
}

type DownloadPackagesFile struct {
	Packages []PackageInfo `json:"packages,omitempty" yaml:"packages,omitempty"`
}

type PackageInfo struct {
	Name    string `json:"name,omitempty" yaml:"name,omitempty"`
	Version string `json:"ver,omitempty" yaml:"ver,omitempty"`
}

func (p PackageInfo) RemotePath() string {
	return filepath.Join(p.Name, p.Version, p.FullName()+".zip")
}

func (p PackageInfo) FullName() string {
	return p.Name + "@" + p.Version
}

type Target struct {
	Path    string `json:"path,omitempty" yaml:"path,omitempty"`
	Exclude string `json:"exclude,omitempty" yaml:"exclude,omitempty"`
}
