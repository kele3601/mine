package app

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log/slog"
	"mine/internal/db/model"
	"os"
	"path/filepath"
)

func NewDB(conf Config) (*gorm.DB, error) {
	var (
		err error
		db  *gorm.DB
	)
	switch conf.DB.Type {
	case "sqlite3":
		slog.Info("1")
		db, err = loadSqlite(conf.DB.DSN)
	default:
		slog.Info("暂不支持的数据库：" + conf.DB.Type)
		return nil, fmt.Errorf("暂不支持的数据库：" + conf.DB.Type)
	}
	if err = autoMigrate(db); nil != err {
		return nil, err
	}
	return db, err
}

func loadSqlite(dsn string) (*gorm.DB, error) {
	// 检查文件是否存在，若不存在直接创建
	if _, err := os.Stat(dsn); os.IsNotExist(err) {
		dir := filepath.Dir(dsn)
		// 递归创建目录
		if err := os.MkdirAll(dir, 0755); nil != err {
			return nil, err
		}
		// 创建文件
		if file, err := os.Create(dsn); nil != err {
			defer file.Close()
			return nil, err
		}
	}
	// 加载数据库
	return gorm.Open(sqlite.Open(dsn), &gorm.Config{})
}

func autoMigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&model.User{}); nil != err {
		return err
	}
	return nil
}
