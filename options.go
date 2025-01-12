package glumberjack

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/ayinke-llc/hermes"
)

// ENUM(gzip)
type CompressionType uint8

// Option is a custom type to configure your logger
type Option func(*Logger)

// WithFileName allows you configure the name of the file to write the
// logs to. If not provided , it would defer to <processname>-glumberjack.log
func WithFileName(name string) Option {
	return func(l *Logger) {
		l.Filename = name
	}
}

// WithMaxSize allows you configure the max size of the log file
func WithMaxSize(maxSize int) Option {
	return func(l *Logger) {
		if maxSize <= 0 {
			return
		}

		l.MaxSize = maxSize
	}
}

// WithMaxAge allows you configure the max days of the log file
func WithMaxAge(maxAge int) Option {
	return func(l *Logger) {
		l.MaxAge = maxAge
	}
}

// WithMaxBackups allows you configure the max number of files to retain.
func WithMaxBackups(maxBackups int) Option {
	return func(l *Logger) {
		l.MaxBackups = maxBackups
	}
}

// WithCompression allows you configure how you want to compress
// the log files when backing them up
func WithCompression(c CompressionType) Option {
	return func(l *Logger) {
		l.compressionType = c
	}
}

// BackupNameGeneratorFunc allows you configure the format for your backups.
// it takes in just one argument which is the name of the current log file being written to at the moment
type BackupNameGeneratorFunc func(string) string

func defaultBackupFunc(useLocalTime bool) func(string) string {
	return func(name string) string {

		dir := filepath.Dir(name)
		filename := filepath.Base(name)
		ext := filepath.Ext(filename)
		prefix := filename[:len(filename)-len(ext)]
		t := currentTime()
		if !useLocalTime {
			t = t.UTC()
		}

		timestamp := t.Format(backupTimeFormat)
		return filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, timestamp, ext))
	}
}

// DefaultBackupName uses a custom layout with the date and time component to
// store the file
func DefaultBackupName(useLocalTime bool) Option {
	return func(l *Logger) {
		l.backupNameFunc = defaultBackupFunc(useLocalTime)
	}
}

// DateBackupName allows you configure your date layout and how you want
// to control how the backup file is named.
func DateBackupName(layout string) Option {
	if hermes.IsStringEmpty(layout) {
		layout = time.DateOnly
	}

	return func(l *Logger) {
		l.backupNameFunc = func(name string) string {

			dir := filepath.Dir(name)
			filename := filepath.Base(name)
			ext := filepath.Ext(filename)
			prefix := filename[:len(filename)-len(ext)]

			return filepath.Join(dir, fmt.Sprintf("%s-%s%s", prefix, currentTime().Format(layout), ext))
		}
	}
}
