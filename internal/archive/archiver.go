package archive

type Archiver interface {
	Archive(archiveName string, filePaths []string) error
	Unarchive(path string) error
}
