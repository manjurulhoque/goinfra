package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Config represents the root configuration structure
type Config struct {
	Version  string            `yaml:"version"`
	AWS      AWSConfig        `yaml:"aws"`
	Resources ResourceConfig  `yaml:"resources"`
}

// AWSConfig represents AWS-specific configuration
type AWSConfig struct {
	Region  string `yaml:"region"`
	Profile string `yaml:"profile"`
}

// ResourceConfig represents all resource definitions
type ResourceConfig struct {
	EC2 []EC2Instance `yaml:"ec2"`
	S3  []S3Bucket   `yaml:"s3"`
	RDS []RDSInstance `yaml:"rds"`
}

// EC2Instance represents an EC2 instance configuration
type EC2Instance struct {
	Name         string            `yaml:"name"`
	InstanceType string            `yaml:"instance_type"`
	AMI          string            `yaml:"ami"`
	Tags         map[string]string `yaml:"tags"`
}

// S3Bucket represents an S3 bucket configuration
type S3Bucket struct {
	Name   string            `yaml:"name"`
	Region string            `yaml:"region"`
	Tags   map[string]string `yaml:"tags"`
}

// RDSInstance represents an RDS instance configuration
type RDSInstance struct {
	Name             string            `yaml:"name"`
	Engine           string            `yaml:"engine"`
	InstanceClass    string            `yaml:"instance_class"`
	AllocatedStorage int               `yaml:"allocated_storage"`
	Tags             map[string]string `yaml:"tags"`
}

// LoadConfig loads the configuration from a YAML file
func LoadConfig(path string) (*Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("failed to parse config file: %w", err)
	}

	return &config, nil
}

// SaveConfig saves the configuration to a YAML file
func SaveConfig(config *Config, path string) error {
	data, err := yaml.Marshal(config)
	if err != nil {
		return fmt.Errorf("failed to marshal config: %w", err)
	}

	if err := os.WriteFile(path, data, 0644); err != nil {
		return fmt.Errorf("failed to write config file: %w", err)
	}

	return nil
} 