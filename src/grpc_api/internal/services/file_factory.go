package services

import (
	"errors"
	"powergate.com/hr_timesheet/pkg/file_source"
)

func getFile(fileType string) (file_source.IFile, error) {
	if fileType == "internal" {
		return file_source.NewInternalFile(file_source.File{}), nil
	}
	return nil, errors.New("file type is not support")
}
