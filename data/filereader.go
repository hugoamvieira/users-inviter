package data

import (
	"bufio"
	"os"
)

// FileReader is a structure that wraps over a bufio scanner
// This was created because, as it stands, we cannot read the file in its entirety at once (incorrect JSON),
// rather we have to read it line by line
type FileReader struct {
	FilePath string
	Queue    *Queue
}

// NewFileReader returns a new FileReader reference with the passed file path and a queue,
// where the data will be written to
func NewFileReader(fp string) *FileReader {
	return &FileReader{
		FilePath: fp,
		Queue:    NewQueue(),
	}
}

// ReadLinesToQueue will open the file and read, line by line, and dump the results in a queue
func (fr *FileReader) ReadLinesToQueue() error {
	f, err := os.Open(fr.FilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	s := bufio.NewScanner(f)
	for s.Scan() {
		fr.Queue.Enqueue(s.Bytes())
	}

	return nil
}
