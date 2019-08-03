package desync

// UntarFilesystem is a filesystem implementation that supports untar'ing
// a catar archive to.
type UntarFilesystem interface {
	CreateDir(n NodeDirectory, opts UntarOptions) error
	CreateFile(n NodeFile, opts UntarOptions) error
	CreateSymlink(n NodeSymlink, opts UntarOptions) error
	CreateDevice(n NodeDevice, opts UntarOptions) error
}
