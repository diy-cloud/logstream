package badger

import (
	"errors"
	"fmt"
	"path/filepath"

	"github.com/dgraph-io/badger/v3"
	"github.com/snowmerak/msgbuf/store"
)

type Badger struct {
	db *badger.DB
}

func New(path ...string) (store.Store, error) {
	b := &Badger{}
	db, err := badger.Open(badger.DefaultOptions(filepath.Join(path...)))
	if err != nil {
		return nil, fmt.Errorf("Badger.New: %w", err)
	}
	b.db = db
	return b, nil
}

func (b *Badger) Get(key string) (string, error) {
	var value string
	if err := b.db.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		if err := item.Value(func(val []byte) error {
			if val == nil {
				return errors.New("key not found")
			}
			value = string(val)
			return nil
		}); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return "", fmt.Errorf("Badger.Get: %w", err)
	}
	return value, nil
}

func (b *Badger) Set(key, value string) error {
	if err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), []byte(value))
	}); err != nil {
		return fmt.Errorf("Badger.Set: %w", err)
	}
	return nil
}

func (b *Badger) SetBytes(key []byte, value []byte) error {
	if err := b.db.Update(func(txn *badger.Txn) error {
		return txn.Set([]byte(key), value)
	}); err != nil {
		return fmt.Errorf("Badger.Set: %w", err)
	}
	return nil
}

func (b *Badger) Close() error {
	if err := b.db.Close(); err != nil {
		return fmt.Errorf("Badger.Close: %w", err)
	}
	return nil
}
