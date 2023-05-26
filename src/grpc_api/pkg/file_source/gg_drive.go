package file_source

type GoogleDriveFile struct {
	File
}

func NewGoogleDriveFile(file File) IFile {
	return &GoogleDriveFile{File: file}
}
