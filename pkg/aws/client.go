package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/rds"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// Client represents an AWS client with all necessary service clients
type Client struct {
	EC2 *ec2.Client
	S3  *s3.Client
	RDS *rds.Client
}

// NewClient creates a new AWS client with the specified region and profile
func NewClient(ctx context.Context, region, profile string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithSharedConfigProfile(profile),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	return &Client{
		EC2: ec2.NewFromConfig(cfg),
		S3:  s3.NewFromConfig(cfg),
		RDS: rds.NewFromConfig(cfg),
	}, nil
}

// ValidateCredentials checks if the AWS credentials are valid
func (c *Client) ValidateCredentials(ctx context.Context) error {
	// Try to make a simple API call to validate credentials
	_, err := c.EC2.DescribeRegions(ctx, &ec2.DescribeRegionsInput{})
	if err != nil {
		return fmt.Errorf("failed to validate AWS credentials: %w", err)
	}
	return nil
}
