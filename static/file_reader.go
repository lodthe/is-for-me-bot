package static

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

type FileReader struct {
	path string
	name string
}

func NewFileReader(fileName string) *FileReader {
	return &FileReader{
		path: "static/" + fileName,
		name: fileName,
	}
}

func (s *FileReader) Name() string {
	return s.name
}

func (s *FileReader) Reader() (io.Reader, error) {
	return os.Open(s.path)
}

func (s *FileReader) Size() int64 {
	file, err := os.Stat(s.path)
	if err != nil {
		log.WithFields(log.Fields{
			"path": s.path,
		}).WithError(err).Error("failed to get the path description")
		return 0
	}

	return file.Size()
}
