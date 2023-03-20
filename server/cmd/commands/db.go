package commands

import (
	"github.com/oseducation/knowledge-graph/model"
	"github.com/oseducation/knowledge-graph/store"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var dbCmd = &cobra.Command{
	Use:   "db",
	Short: "Management of Database",
}

var dbCreateAdmin = &cobra.Command{
	Use:     "create",
	Short:   "Create an admin",
	Long:    `Create an admin.`,
	Example: `  db create --email adminEmail --password adminPassword`,
	RunE:    createAdminCmdF,
}

func init() {
	dbCreateAdmin.Flags().String("email", "", "Admin email")
	dbCreateAdmin.Flags().String("password", "", "Admin password")
	dbCmd.AddCommand(dbCreateAdmin)
	rootCmd.AddCommand(dbCmd)
}

func createAdminCmdF(command *cobra.Command, args []string) error {
	email, err := command.Flags().GetString("email")
	if err != nil || email == "" {
		return errors.New("email is required")
	}
	password, err := command.Flags().GetString("password")
	if err != nil || password == "" {
		return errors.New("password is required")
	}

	db := store.CreateStore()

	user := &model.User{
		Email:         email,
		EmailVerified: true,
		Password:      password,
		Role:          model.AdminRole,
	}
	_, err = db.User().Save(user)
	if err != nil {
		return errors.Wrap(err, "can't save admin in DB")
	}
	return nil
}