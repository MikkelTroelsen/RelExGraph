package main

import (
	"context"
	"fmt"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Neo4jClient struct {
	driver neo4j.DriverWithContext
}

func GetNeo4jClient() (*Neo4jClient, error) {
	ctx := context.Background()
	driver, err := neo4j.NewDriverWithContext(
		"bolt://localhost:7687",
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
