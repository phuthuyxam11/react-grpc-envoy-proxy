package file_source

type IFile interface {
	getName() string
	getExtension() string
}

type File struct {
	name      string
	extension string
}

func (f *File) getName() string {
	return f.name
}

//func (f *File) SetName(name string) {
//	f.name = name
//}

func (f *File) getExtension() string {
	return f.extension
}

//func (f *File) SetExtension(extension string) {
//	f.extension = extension
//}
