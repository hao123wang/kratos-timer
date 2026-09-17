package data

import (
	"fmt"

	"kratos-timer/internal/conf"

	"github.com/go-kratos/kratos/v3/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTimerRepo)

// Data holds the long-lived storage clients shared by repos.
type Data struct {
	db *gorm.DB
}

// NewData opens the MySQL client via GORM and returns it with a cleanup function.
func NewData(c *conf.Data) (*Data, func(), error) {
	dc := c.GetDatabase()
	if dc == nil || dc.GetSource() == "" {
		return nil, nil, fmt.Errorf("data.database.source is required")
	}
	if driver := dc.GetDriver(); driver != "" && driver != "mysql" {
		return nil, nil, fmt.Errorf("unsupported database driver %q, expected mysql", driver)
	}

	logLevel := gormlogger.Warn
	if dc.GetDebug() {
		logLevel = gormlogger.Info
	}

	db, err := gorm.Open(mysql.Open(dc.GetSource()), &gorm.Config{
		Logger: gormlogger.Default.LogMode(logLevel),
	})
	if err != nil {
		return nil, nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, nil, err
	}

	if dc.GetAutoMigrate() {
		if err := db.AutoMigrate(); err != nil {
			_ = sqlDB.Close()
			return nil, nil, err
		}
	}

	cleanup := func() {
		log.Info("closing the data resources")
		if err := sqlDB.Close(); err != nil {
			log.Error("failed closing the database", "err", err)
		}
	}
	return &Data{db: db}, cleanup, nil
}
