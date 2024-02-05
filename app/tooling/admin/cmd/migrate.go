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
		dbUser, err := cmd.Flags().GetString(dbUser)
		if err != nil {
			return fmt.Errorf("retrieve dbUser: %w", err)
		}

		dbPass, err := cmd.Flags().GetString(dbPass)
		if err != nil {
			return fmt.Errorf("retrieve dbPass: %w", err)
		}

		dbHost, err := cmd.Flags().GetString(dbHost)
		if err != nil {
			return fmt.Errorf("retrieve dbHost: %w", err)
		}

		dbName, err := cmd.Flags().GetString(dbName)
		if err != nil {
			return fmt.Errorf("retrieve dbName: %w", err)
		}

		dbConfig := sqldb.Config{
			User:         dbUser,
			Password:     dbPass,
			HostPort:     dbHost,
			Name:         dbName,
			MaxIdleConns: 2,
			MaxOpenConns: 0,
			DisableTLS:   true,
		}

		return performMigrate(dbConfig)
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().String(dbUser, "postgres", "user for access to the db")
	migrateCmd.Flags().String(dbPass, "postgres", "password for access to the db")
	migrateCmd.Flags().String(dbHost, "database-service.liars-system.svc.cluster.local", "host and port to db")
	migrateCmd.Flags().String(dbName, "postgres", "name of the db to access")
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
