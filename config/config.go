package config

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	DefaultApplicationRefreshTokenDuration = 24 * time.Hour
	DefaultApplicationAccessTokenDuration  = 5 * time.Minute
	DefaultPostgresMaxIdleConns            = 3
	DefaultPostgresMaxOpenConns            = 5
	DefaultPostgresMaxConnLifetime         = 1 * time.Hour
)

func GetConfig() {
	viper.SetEnvPrefix("svc")

	viper.AddConfigPath(".")
	viper.AddConfigPath("./..")
	viper.SetConfigName("config")

	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		logrus.WithField("%v", err)
	}
}

func ApplicationName() string {
	return viper.GetString("application.name")
}

func Port() string {
	return viper.GetString("application.port")
}

func LogLevel() string {
	return viper.GetString("log-level")
}

func RefreshTokenDuration() time.Duration {
	cfg := viper.GetString("application.refresh-token-duration")
	res, err := time.ParseDuration(cfg)
	if err != nil {
		return DefaultApplicationRefreshTokenDuration
	}

	return res
}

func AccessTokenDuration() time.Duration {
	cfg := viper.GetString("application.access-token-duration")
	res, err := time.ParseDuration(cfg)
	if err != nil {
		return DefaultApplicationAccessTokenDuration
	}

	return res
}

func PostgresHost() string {
	return viper.GetString("postgres.host")
}

func PostgresPort() string {
	return viper.GetString("postgres.port")
}

func PostgresDatabase() string {
	return viper.GetString("postgres.database")
}

func PostgresUsername() string {
	return viper.GetString("postgres.username")
}

func PostgresPassword() string {
	return viper.GetString("postgres.password")
}

func PostgresSSLMode() string {
	return viper.GetString("postgres.sslmode")
}

func PostgresMaxIdleConnections() int {
	if viper.GetInt("postgres.max-idle-connections") <= 0 {
		return DefaultPostgresMaxIdleConns
	}
	return viper.GetInt("postgres.max-idle-connections")
}

func PostgresMaxOpenConnections() int {
	if viper.GetInt("postgres.max-open-connections") <= 0 {
		return DefaultPostgresMaxOpenConns
	}
	return viper.GetInt("postgres.max-open-connections")
}

func PostgresMaxConnectionLifetime() time.Duration {
	cfg := viper.GetString("postgres.max-connection-lifetime")
	time, err := time.ParseDuration(cfg)
	if err != nil {
		return DefaultPostgresMaxConnLifetime
	}

	return time
}

func DatabaseDSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		PostgresUsername(),
		PostgresPassword(),
		PostgresHost(),
		PostgresPort(),
		PostgresDatabase(),
		PostgresSSLMode(),
	)
}
