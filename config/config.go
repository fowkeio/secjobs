package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	// Initialize SQLite
	db, err = InitializeSQLite()
	if err != nil {
		return fmt.Errorf("Error initializing sqlite: %v", err)
	}
	return nil
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
