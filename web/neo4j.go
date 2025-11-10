package main

import (
	"context"
	"fmt"
	"os"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jClient struct {
	driver neo4j.DriverWithContext
}

func GetNeo4jClient() (*Neo4jClient, error) {
	ctx := context.Background()

	// Get Neo4j URI from environment variable, default to localhost
	neo4jURI := os.Getenv("NEO4J_URI")
	if neo4jURI == "" {
		neo4jURI = "bolt://localhost:7687"
	}

	driver, err := neo4j.NewDriverWithContext(
		neo4jURI,
		neo4j.NoAuth())

	if err != nil {
		return nil, err
	}

	err = driver.VerifyConnectivity(ctx)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connection established.")
	return &Neo4jClient{driver: driver}, nil
}

func (c *Neo4jClient) Close(ctx context.Context) {
	c.driver.Close(ctx)
}

func (c *Neo4jClient) ExecuteQuery(ctx context.Context, query string) {
	_, err := neo4j.ExecuteQuery(ctx, c.driver, query, nil, neo4j.EagerResultTransformer)
	if err != nil {
		panic(err)
	}
}
