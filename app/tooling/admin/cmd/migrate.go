package cmd

import (
	"context"
	"fmt"
	"time"

	"github.com/ardanlabs/liarsdice/business/data/migrate"
	"github.com/ardanlabs/liarsdice/business/data/sqldb"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate the database",
	Long:  `Migrates the database to its most current schema.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		dbConfig := sqldb.Config{
			User:         "postgres",
			Password:     "postgres",
			HostPort:     "database-service.liars-system.svc.cluster.local",
			Name:         "postgres",
			MaxIdleConns: 2,
			MaxOpenConns: 0,
			DisableTLS:   true,
		}

		return performMigrate(dbConfig)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}

func performMigrate(cfg sqldb.Config) error {
	db, err := sqldb.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrate.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")
	return nil
}
