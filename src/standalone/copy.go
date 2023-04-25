package standalone

import (
	"fmt"
	"io"
	"os"

	"path/filepath"

	"github.com/pkg/errors"
	"github.com/spf13/afero"
)

func copyFolder(fs afero.Fs, source string, destination string) error {
	sourceInfo, err := fs.Stat(source)
	if err != nil {
		return err
	}
	if !sourceInfo.IsDir() {
		return errors.Errorf("%s is not a directory", source)
	}

	err = fs.MkdirAll(destination, sourceInfo.Mode())
	if err != nil {
		return err
	}

	entries, err := afero.ReadDir(fs, source)
	if err != nil {
		return err
	}
	for _, entry := range entries {
		sourcePath := filepath.Join(source, entry.Name())
		destinationPath := filepath.Join(destination, entry.Name())
		if entry.IsDir() {
			err = copyFolder(fs, sourcePath, destinationPath)
			if err != nil {
				return err
			}
		} else {
			log.Info(fmt.Sprintf("copying from %s to %s ", sourcePath, destinationPath))
			err = copyFile(fs, sourcePath, destinationPath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func copyFile(fs afero.Fs, sourcePath string, destinationPath string) error {
	sourceFile, err := fs.Open(sourcePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	sourceInfo, err := sourceFile.Stat()
	if err != nil {
		return err
	}

	destinationFile, err := fs.OpenFile(destinationPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, sourceInfo.Mode())
	if err != nil {
		return err
	}
	defer destinationFile.Close()
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	err = destinationFile.Sync()
	if err != nil {
		return err
	}
	return nil
}
