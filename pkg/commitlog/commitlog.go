// Package commitlog provide a simple commit log file (and index) implementation.
// A commit log file is an append only file to put arbitrary data and read it sequentially.
package commitlog

import (
	"io"
	"sync"
	"time"
)

// CommitLog struct holds all handlers (files and channels) to work with commit logs.
type CommitLog struct {
	Name string

	writer *commitLog

	io.Writer
	io.ReaderAt
	io.Closer
}

type action int

const (
	write action = iota
	flush        = iota
	stop         = iota
)

type commitLog struct {
	input chan *commitLogEntry
}

type commitLogEntry struct {
	action    action
	data      []byte
	timestamp time.Time
}

var (
	guards = make(map[string]*sync.Once)
	cmlogs = make(map[string]*CommitLog)
)

// GetInstance returns an instance of a CommitLog by name
// A CommitLog instance holds a pointer to the underling
// single writer and a dedicated reader.
func GetInstance(name string) (*CommitLog, error) {
	guards[name].Do(func() {
		cmlogs[name] = &CommitLog{
			Name: name,
		}
	})
	return cmlogs[name], nil
}

// Finish closes permanently (writer and all readers) of
// a given CommitLog
func Finish(name string) error {
	return nil
}

// FinishAll closes permanently all CommitLogs
func FinishAll() error {
	return nil
}

// Close finishes the specific reader of a CommitLog
func (cmlog *CommitLog) Close() error {
	return nil
}

// Write adds data to commit log.
func (cmlog *CommitLog) Write(data []byte) (n int, err error) {
	cmlog.writer.input <- &commitLogEntry{write, data, time.Now()}
	return len(data), nil
}

// ReadAt reads an entry at a given offset
func (cmlog *CommitLog) ReadAt(p []byte, off int64) (n int, err error) {
	return 0, nil
}
