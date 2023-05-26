package file_source

type InternalFile struct {
	File
}

func NewInternalFile(file File) IFile {
	return &InternalFile{File: file}
}
