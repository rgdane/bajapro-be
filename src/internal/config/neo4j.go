package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j/config"
)

var Neo4jDriver neo4j.DriverWithContext

func InitNeo4j() {
	uri := os.Getenv("NEO4J_URI")       // example: neo4j://localhost:7687
	user := os.Getenv("NEO4J_USER")     // neo4j
	pass := os.Getenv("NEO4J_PASSWORD") // password
	timeout := 5 * time.Second

	driver, err := neo4j.NewDriverWithContext(
		uri,
		neo4j.BasicAuth(user, pass, ""),
		func(c *config.Config) {
			c.MaxConnectionPoolSize = 10
			c.SocketConnectTimeout = timeout
		},
	)

	if err != nil {
		log.Fatalf("❌ Failed to create Neo4j driver: %v", err)
	}

	if err := driver.VerifyConnectivity(context.Background()); err != nil {
		log.Fatalf("❌ Failed to connect to Neo4j: %v", err)
	}

	log.Println("✅ Connected to Neo4j:", uri)
	Neo4jDriver = driver
}

func GetNeo4j() neo4j.DriverWithContext {
	return Neo4jDriver
}

func CloseNeo4j() {
	if Neo4jDriver != nil {
		Neo4jDriver.Close(context.Background())
		log.Println("🔌 Neo4j connection closed")
	}
}
