package console

import (
	_ "github.com/lib/pq"

	"github.com/rhtyx/bayarind-service.git/config"
	"github.com/rhtyx/bayarind-service.git/db"
	"github.com/rhtyx/bayarind-service.git/utils"

	"github.com/rubenv/sql-migrate"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use: "migrate",
	Run: doMigrate,
}

func init() {
	migrateCmd.PersistentFlags().Int("step", 0, "maximum migration steps")
	migrateCmd.PersistentFlags().String("direction", "up", "migration direction")
	RootCmd.AddCommand(migrateCmd)
}

func doMigrate(cmd *cobra.Command, _ []string) {
	db.InitPostgresDB()
	direction := cmd.Flag("direction").Value.String()

	migrations := &migrate.FileMigrationSource{
		Dir: "./migrations",
	}

	migrate.SetTable("schema_migrations")
	db, err := db.PostgresDB.DB()
	if err != nil {
		logrus.WithField("DatabaseDSN", config.DatabaseDSN()).Fatal("Failed to connect database: ", err)
	}

	var n int
	if direction == "down" {
		n, err = migrate.Exec(db, "postgres", migrations, migrate.Down)
	} else {
		n, err = migrate.Exec(db, "postgres", migrations, migrate.Up)
	}
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"db":         db,
			"migrations": utils.Dump(migrations),
			"direction":  direction}).
			Fatal("Failed to migrate database: ", err)
	}

	logrus.Infof("Applied %d migrations!\n", n)
}
