package net

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/avran02/package-manager/config"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type sshRemoteServerConn struct {
	sshClient  *ssh.Client
	sftpCilent *sftp.Client
}

func (s sshRemoteServerConn) Close() error {
	s.sftpCilent.Close()
	return s.sshClient.Close()
}

func (s sshRemoteServerConn) DownloadFile(dst *os.File, name string) error {
	remoteFile, err := s.sftpCilent.Open(name)
	if err != nil {
		return fmt.Errorf("can't open file %s: %w", name, err)
	}
	defer remoteFile.Close()

	if _, err = io.Copy(dst, remoteFile); err != nil {
		return err
	}
	return nil
}

func (s sshRemoteServerConn) UploadFile(src *os.File, name string) error {
	if err := s.sftpCilent.MkdirAll(filepath.Dir(name)); err != nil {
		return fmt.Errorf("can't create directory %s: %w", filepath.Dir(name), err)
	}

	remoteFile, err := s.sftpCilent.Create(name)
	if err != nil {
		return fmt.Errorf("can't create file %s: %w", name, err)
	}
	defer remoteFile.Close()

	if _, err = io.Copy(remoteFile, src); err != nil {
		return err
	}
	return nil
}

func NewRemoteServerConnection(config config.SSHConfig) RemoteServerConn {
	clientConfig := ssh.ClientConfig{
		User: config.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(config.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	sshClient, err := ssh.Dial("tcp", config.Addr(), &clientConfig)
	if err != nil {
		log.Fatalf("can't connect to remote server: %s", err.Error())
	}

	sftpCilent, err := sftp.NewClient(sshClient)
	if err != nil {
		log.Fatalf("can't create scp client: %s", err.Error())
	}

	return sshRemoteServerConn{
		sshClient:  sshClient,
		sftpCilent: sftpCilent,
	}
}
