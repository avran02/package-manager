package app

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/avran02/package-manager/config"
	"github.com/avran02/package-manager/internal/archive"
	"github.com/avran02/package-manager/internal/models"
	"github.com/avran02/package-manager/internal/net"
	"github.com/avran02/package-manager/internal/parser"
)

type App struct {
	config   config.Config
	conn     net.RemoteServerConn
	archiver archive.Archiver
}

func (a *App) CreatePackage(path string) {
	pkgFile, err := parser.ParcePackageFile(path)
	if err != nil {
		log.Fatalf("can't create package on %s: %s", path, err.Error())
	}

	files := make([]string, 0)
	for _, target := range pkgFile.Targets {
		matchedFiles, err := filepath.Glob(target.Path)
		if err != nil {
			log.Fatalf("can't create package on %s: %s", path, err.Error())
		}
		for _, filePath := range matchedFiles {
			skip := false
			if target.Exclude != "" {
				matchedExcludeFiles, err := filepath.Glob(filepath.Join(target.Exclude))
				if err != nil {
					slog.Error("Error matching exclude files:", "error", err.Error())
					continue
				}
				for _, excludeFile := range matchedExcludeFiles {
					if excludeFile == filePath {
						slog.Info("Skipping file:", "file", filePath)
						skip = true
					}
				}
			}
			if !skip {
				files = append(files, filePath)
			}
		}
	}

	archiveName := pkgFile.FullName() + ".zip"
	if err := a.archiver.Archive(archiveName, files); err != nil {
		log.Fatalf("can't create package on %s: %s", path, err.Error())
	}

	f, err := os.Open(archiveName)
	if err != nil {
		log.Fatalf("can't open archive on %s: %s", path, err.Error())
	}

	if err := a.conn.UploadFile(f, pkgFile.RemotePath()); err != nil {
		log.Fatalf("can't upload archive %s: %s", path, err.Error())
	}
}

func (a *App) UpdatePackage(path string) {
	pkgFile, err := parser.ParseDownloadPkgsFile(path)
	if err != nil {
		log.Fatalf("can't update package: %s", err.Error())
	}

	archivesPath := config.ArchivesFolderName
	slog.Info("Downloading", "path", archivesPath)
	if err := os.MkdirAll(archivesPath, 0777); err != nil {
		log.Fatalf("can't create dependencies directory: %s", err.Error())
	}

	for _, pkg := range pkgFile.Packages {
		pkgPath := filepath.Join(archivesPath, pkg.FullName())
		f, err := os.Create(pkgPath)
		if err != nil {
			log.Fatalf("can't create archive file: %s", err.Error())
		}
		defer f.Close()

		if err := a.conn.DownloadFile(f, pkg.RemotePath()); err != nil {
			slog.Error("Can't download", "pkg name", pkg.Name, "version", pkg.Version, "error", err.Error())
			continue
		}
		a.InstallPackage(pkgPath)
	}

	for _, pkg := range pkgFile.Packages {
		pkgPath := filepath.Join(archivesPath, pkg.FullName())
		if err := os.Rename(pkgPath+".package", pkg.FullName()+".package"); err != nil {
			slog.Error("Can't rename", "pkg name", pkg.Name, "version", pkg.Version, "error", err.Error())
			continue
		}
	}
	if err := os.RemoveAll(archivesPath); err != nil {
		log.Fatalf("can't remove dependencies directory: %s", err.Error())
	}
}

func (a *App) DownloadPackages(packages ...models.PackageInfo) {
	for _, pkg := range packages {
		pkgPath := filepath.Join("", pkg.FullName())
		f, err := os.Create(pkgPath)
		if err != nil {
			log.Fatalf("can't create archive file: %s", err.Error())
		}
		defer f.Close()

		if err := a.conn.DownloadFile(f, pkg.RemotePath()); err != nil {
			log.Fatalf("can't download archive file %s: %s", pkg.FullName(), err.Error())
		}
	}
}

func (a *App) InstallPackage(archivePath string) {
	if err := a.archiver.Unarchive(archivePath); err != nil {
		slog.Error("can't extruct file from archive", "file", archivePath, "error", err)
	}

	pkgFilePath := filepath.Join(archivePath+".package", config.PackageFileName)

	pkgFile, err := parser.ParcePackageFile(pkgFilePath)
	if err != nil {
		slog.Error("Can't read package metadata", "file", archivePath, "error", err)
		return
	}

	for _, pkg := range pkgFile.Dependencies {
		a.UpdatePackage(pkg.RemotePath())
	}
}

func New() *App {
	conf := config.New()
	return &App{
		config:   conf,
		conn:     net.NewRemoteServerConnection(conf.SSHConfig),
		archiver: archive.NewZipArchiver(),
	}
}
