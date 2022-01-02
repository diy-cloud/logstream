package textfile

import (
	"errors"
	"fmt"
	"os"

	"github.com/snowmerak/msgbuf/store"
	"github.com/snowmerak/msgbuf/unlock"
)

type TextFile struct {
	unlock.TLock

	path string
	file *os.File
}

func New(path string) (store.Store, error) {
	tf := &TextFile{
		path: path,
	}
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return nil, fmt.Errorf("TextFile.New: %w", err)
	}
	tf.file = f
	return tf, nil
}

func (tf *TextFile) Get(key string) (string, error) {
	return "", errors.New("TF.Get: not implemented")
}

func (tf *TextFile) Set(key, value string) error {
	tf.Lock()
	defer tf.Unlock()
	if _, err := tf.file.WriteString(fmt.Sprintf("%s:%s\n", key, value)); err != nil {
		return fmt.Errorf("TextFile.Set: %w", err)
	}
	return nil
}

func (tf *TextFile) SetBytes(key []byte, value []byte) error {
	tf.Lock()
	defer tf.Unlock()
	if _, err := tf.file.Write(value); err != nil {
		return fmt.Errorf("TextFile.SetBytes: %w", err)
	}
	return nil
}

func (tf *TextFile) Close() error {
	if err := tf.file.Close(); err != nil {
		return fmt.Errorf("TextFile.Close: %w", err)
	}
	return nil
}
