package filesystem

import (
	"log"
	"os"

	"github.com/hashicorp/terraform-ls/internal/source"
)

type Document interface {
	DocumentHandler
	Text() ([]byte, error)
	Lines() source.Lines
	LanguageID() string
	Version() int
}

type DocumentHandler interface {
	URI() string
	FullPath() string
	Dir() string
	Filename() string
}

type VersionedDocumentHandler interface {
	DocumentHandler
	Version() int
}

type DocumentChange interface {
	Text() string
	Range() *Range
}

type DocumentChanges []DocumentChange

type DocumentStorage interface {
	// LS-specific methods
	CreateDocument(DocumentHandler, string, []byte) error
	CreateAndOpenDocument(DocumentHandler, string, []byte) error
	GetDocument(DocumentHandler) (Document, error)
	CloseAndRemoveDocument(DocumentHandler) error
	ChangeDocument(VersionedDocumentHandler, DocumentChanges) error
	HasOpenFiles(path string) (bool, error)
}

type Filesystem interface {
	DocumentStorage

	SetLogger(*log.Logger)

	// direct FS methods
	ReadFile(name string) ([]byte, error)
	ReadDir(name string) ([]os.FileInfo, error)
	Open(name string) (File, error)
	Stat(name string) (os.FileInfo, error)
}

// File represents an open file in FS
// See io/fs.File in http://golang.org/s/draft-iofs-design
type File interface {
	Stat() (os.FileInfo, error)
	Read([]byte) (int, error)
	Close() error
}
