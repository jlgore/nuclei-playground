package main

import (
	"github.com/pulumi/pulumi-aws/sdk/v6/go/aws/rds"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// Create a public RDS instance
		publicRdsInstance, err := rds.NewInstance(ctx, "public-rds-instance", &rds.InstanceArgs{
			InstanceClass:      pulumi.String("db.t3.micro"),
			AllocatedStorage:   pulumi.Int(20),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("8.0"),
			Username:           pulumi.String("admin"),
			Password:           pulumi.String("password123"),
			DbName:             pulumi.String("exampledb"),
			SkipFinalSnapshot:  pulumi.Bool(true),
			PubliclyAccessible: pulumi.Bool(true), // Set to true for public instance
		})
		if err != nil {
			return err
		}

		// Create a private RDS instance
		privateRdsInstance, err := rds.NewInstance(ctx, "private-rds-instance", &rds.InstanceArgs{
			InstanceClass:      pulumi.String("db.t3.micro"),
			AllocatedStorage:   pulumi.Int(20),
			Engine:             pulumi.String("mysql"),
			EngineVersion:      pulumi.String("8.0"),
			Username:           pulumi.String("admin"),
			Password:           pulumi.String("password123"),
			DbName:             pulumi.String("exampledb"),
			SkipFinalSnapshot:  pulumi.Bool(true),
			PubliclyAccessible: pulumi.Bool(false), // Set to false for private instance
		})
		if err != nil {
			return err
		}

		// Export the endpoint of the public RDS instance
		ctx.Export("publicRdsEndpoint", publicRdsInstance.Endpoint)

		// Export the endpoint of the private RDS instance
		ctx.Export("privateRdsEndpoint", privateRdsInstance.Endpoint)

		return nil
	})
}
