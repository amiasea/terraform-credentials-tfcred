package commands

import (
	"fmt"

	"github.com/amiasea/terraform-credentials-tfcred/internal/store"
	"github.com/spf13/cobra"
)

// NewInitCmd creates the init command.
func NewInitCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize tfcred storage",
		Long: `Initializes tfcred local storage.

The installer is responsible for installing the Terraform credentials helper
and configuring Terraform integration.`,
		Run: func(cmd *cobra.Command, _ []string) {
			domain, _ := cmd.Flags().GetString("domain")

			if domain == "" {
				domain = promptDefaultDomain()
			}

			runInit(domain)

			fmt.Println("[tfcred] ✅ Initialization completed successfully.")
		},
	}

	cmd.Flags().String("domain", "", "default Terraform domain")

	return cmd
}

func runInit(defaultDomain string) {
	// Initialize tfcred internal storage only.
	//
	// Installation concerns are intentionally handled by the installer:
	// - Terraform helper deployment
	// - terraform.tfrc configuration
	// - TF_CLI_CONFIG_FILE environment variable
	// - command alias registration

	store.Init(defaultDomain)

	fmt.Println("[tfcred] ✅ tfcred storage initialized.")
}
