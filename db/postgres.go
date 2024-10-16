package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/rhtyx/bayarind-service.git/config"

	"github.com/sirupsen/logrus"
)

var (
	PostgresDB *gorm.DB
)

func InitPostgresDB() {
	conn, err := OpenPostgresDB(config.DatabaseDSN())
	if err != nil {
		logrus.WithField("dsn", config.DatabaseDSN()).Fatal("Failed to connect to database ", err)
	}

	PostgresDB = conn
	logrus.Info("Successfully connected to database")
}

func OpenPostgresDB(dsn string) (*gorm.DB, error) {
	dialector := postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	conn, err := db.DB()
	if err != nil {
		logrus.Fatal(err)
	}
	conn.SetMaxIdleConns(config.PostgresMaxIdleConnections())
	conn.SetMaxOpenConns(config.PostgresMaxOpenConnections())
	conn.SetConnMaxLifetime(config.PostgresMaxConnectionLifetime())

	return db, nil
}
