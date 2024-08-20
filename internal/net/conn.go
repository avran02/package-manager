package net

import "os"

type RemoteServerConn interface {
	DownloadFile(file *os.File, name string) error
	UploadFile(file *os.File, name string) error
	Close() error
}
