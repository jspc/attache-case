package attachecase

import ()

const (
	fsFile = iota + 1000
	fsDir
)

// FSData is an interface which any type of
// filesystem node must implement.
type FSData interface {
	String() string
}

// Filesystem represents a filesystem
type Filesystem struct {
	ID   string
	Type int
	Root Directory
}

// File represents a file, including its contents and so forth
type File struct {
	data string
}

// String returns the base64 encoded contents of a file
func (f File) String() string {
	return f.data
}

// Directory embeds anything which implements FSData- so other directories
// or files
type Directory struct {
	contents []FSData
}

// String returns an empty string; it exists to ensure a directory can be
// treated as implementing FSData
func (d Directory) String() string {
	return ""
}
