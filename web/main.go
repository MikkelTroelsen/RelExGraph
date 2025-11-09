package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router with default middleware (logger and recovery)
	r := gin.Default()

	// Serve static files (CSS, JS, images, etc.)
	r.Static("/static", "./static")

	// Serve HTML Templates
	r.LoadHTMLGlob("templates/*")

	// Connecto to neo4j
	_, err := GetNeo4jClient()
	if err != nil {
		log.Fatal("Neo4j vonnection failed:", err)
	}

	// Input page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "input_page.html", nil)
	})

	// Handle form submission
	r.POST("/generate-graph", func(c *gin.Context) {
		context := c.PostForm("context")
		result := callRelik(context)
		client, err := GetNeo4jClient()
		if err != nil {
			panic(err)
		}
		client.ExecuteQuery(c, CreateQuery(result))
		c.HTML(http.StatusOK, "results.html", gin.H{
			"context":   context,
			"graphData": CreateQuery(result), // the response from your API
		})
	})
	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run()
}

func CreateQuery(relations []Relation) string {
	uniqueNodes := make(map[string]string)

	var firstQueryPart string
	var secoundQueryPart string
	for _, relation := range relations {
		subject := FormatStrings(relation.subject)
		object := FormatStrings(relation.object)
		uniqueNodes[subject] = subject
		uniqueNodes[object] = object

		label := strings.ToUpper(FormatStrings(relation.label))

		secoundQueryPart += fmt.Sprintf("MERGE(%s)-[:%s]->(%s)\n",
			subject,
			label,
			object,
		)
	}
	for key, value := range uniqueNodes {
		firstQueryPart += fmt.Sprintf("Merge(%s:%s)\n",
			key,
			value,
		)
	}
	return firstQueryPart + secoundQueryPart
}

func FormatStrings(str string) string {
	str = strings.ReplaceAll(str, " ", "_")
	str = strings.ReplaceAll(str, "-", "_")
	str = strings.ReplaceAll(str, ".", "")
	str = strings.ReplaceAll(str, ",", "")

	first := str[0]
	if first >= '0' && first <= '9' {
		return "_" + str
	}

	return str
}
