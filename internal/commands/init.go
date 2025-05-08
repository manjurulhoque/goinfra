package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a new infrastructure project",
	Long: `Initialize a new infrastructure project by creating the necessary directory structure
and configuration files. This will create:
- goinfra.yaml: Main configuration file
- .goinfra/: Directory for state and temporary files`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Create .goinfra directory
		if err := os.MkdirAll(".goinfra", 0755); err != nil {
			return fmt.Errorf("failed to create .goinfra directory: %w", err)
		}

		// Create initial goinfra.yaml if it doesn't exist
		if _, err := os.Stat("goinfra.yaml"); os.IsNotExist(err) {
			config := `# GoInfra Configuration
version: "1.0"

# AWS Configuration
aws:
  region: us-east-1
  profile: default

# Resource Definitions
resources:
  # EC2 Instances
  ec2:
    - name: example-instance
      instance_type: t2.micro
      ami: ami-0c55b159cbfafe1f0
      tags:
        Environment: development

  # S3 Buckets
  s3:
    - name: example-bucket
      region: us-east-1
      tags:
        Environment: development

  # RDS Instances
  rds:
    - name: example-db
      engine: mysql
      instance_class: db.t3.micro
      allocated_storage: 20
      tags:
        Environment: development
`
			if err := os.WriteFile("goinfra.yaml", []byte(config), 0644); err != nil {
				return fmt.Errorf("failed to create goinfra.yaml: %w", err)
			}
		}

		fmt.Println("✅ Project initialized successfully!")
		fmt.Println("📝 Edit goinfra.yaml to define your infrastructure")
		return nil
	},
}

func GetInitCommand() *cobra.Command {
	return initCmd
}
