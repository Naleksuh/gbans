package cmd

import (
	"github.com/leighmacdonald/gbans/internal/config"
	"github.com/leighmacdonald/gbans/internal/store"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var downAll = false

// migrateCmd loads the db schema
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Create or update the database schema",
	Run: func(cmd *cobra.Command, args []string) {
		act := store.MigrateUp
		if downAll {
			act = store.MigrateDn
		}
		db, err := store.New(config.DB.DSN)
		if err != nil {
			log.Fatalf("Failed to initialize db connection: %v", err)
		}
		if err := db.Migrate(store.MigrationAction(act)); err != nil {
			if err.Error() == "no change" {
				log.Infof("Migration at latest version")
			} else {
				log.Fatalf("Could not migrate schema: %v", err)
			}
		} else {
			log.Infof("Migration completed successfully")
		}
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().BoolVarP(&downAll, "down", "d", false, "Fully reverts all migrations")
}
