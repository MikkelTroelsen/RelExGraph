# RelExGraph

RelExGraph is a web application that automatically extracts relationships from text using the **ReLiK model** and visualizes them as knowledge graphs in **Neo4j**. Simply input your text, and the system will identify subjects, objects, and the relationships between them, creating an interactive graph database representation.

## About This Project

This project was created as an experiment to explore applications with neural language models, especially smaller specialized models that can run locally. It also served as my first Go application to learn and play around with how Go works.

**Key Motivations:**
- Explore the capabilities of smaller, locally-runnable neural models for NLP tasks
- Test relation extraction quality with resource-constrained, task-specific models
- Learn Go web development and its ecosystem
- Explore the basics of GraphRAG (Graph Retrieval-Augmented Generation) by building the foundation: relation extraction and graph storage, without the LLM querying component

**Findings:**
While it's definitely possible to extract relations with a small, specialized model like ReLiK, the results include some noise that makes it challenging for production use. The model can identify many relationships correctly, but also produces some incorrect or irrelevant extractions. Most of these issues could likely be solved with larger models but that was outside the scope of this experiment which focused on local, lightweight solutions.

This project demonstrates the foundational components of a GraphRAG system - the extraction and storage phases - showing how unstructured text can be transformed into a queryable knowledge graph.

> **Note:** The HTML/CSS for the web interface was generated with assistance from LLMs.

## Features

- **Automatic Relation Extraction**: Uses the ReLiK (Retriever-Reader) neural model to identify entities and their relationships
- **Knowledge Graph Visualization**: Stores extracted relations in Neo4j for easy exploration and visualization
- **Simple Web Interface**: Clean, user-friendly interface built with Go and Gin framework
- **Docker Support**: Fully containerized setup for easy deployment
- **Real-time Processing**: Submit text and get immediate graph generation

## Architecture

The project consists of three main components:

1. **Web Frontend** (Go + Gin)
   - Serves the user interface
   - Handles form submissions
   - Coordinates between ReLiK and Neo4j services

2. **ReLiK Service** (Python + Flask)
   - Runs the ReLiK neural model for relation extraction
   - Exposes REST API for text analysis
   - Returns structured relation data (subject, label, object)

3. **Neo4j Database**
   - Stores and visualizes knowledge graphs
   - Provides graph query capabilities
   - Accessible via browser interface

## Prerequisites

- Docker and Docker Compose
- At least 4GB RAM (the neural model requires memory)
- Ports 5000, 7474, 7687, and 8080 available

## Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/MikkelTroelsen/RelExGraph.git
   cd RelExGraph
   ```

2. **Start all services with Docker Compose**
   ```bash
   docker-compose up --build
   ```

   This will start:
   - ReLiK service on `http://localhost:5000`
   - Neo4j database on `http://localhost:7474` (browser UI) and `bolt://localhost:7687` (database)
   - Web application on `http://localhost:8080`

3. **Access the application**
   - Open your browser and go to `http://localhost:8080`
   - Enter text in the input field
   - Click "Analyze" and wait for processing
   - View results and the generated Cypher query

4. **View the graph in Neo4j**
   - Navigate to `http://localhost:7474`
   - Run queries to explore your knowledge graph

## Usage

### Web Interface

1. Enter or paste your text into the input field
2. Click the "Analyze" button
3. Wait for the ReLiK model to process (this may take a moment)
4. View the extracted relations and generated Cypher query
5. Check Neo4j browser to visualize the graph

### Example Input

```
Apple was founded by Steve Jobs in California.
Microsoft was created by Bill Gates.
Both companies are based in the United States.
```

This will extract relations like:
- `Apple` -[FOUNDED_BY]-> `Steve_Jobs`
- `Apple` -[LOCATED_IN]-> `California`
- `Microsoft` -[CREATED_BY]-> `Bill_Gates`

## Development

### Running Without Docker

**ReLiK Service:**
```bash
cd RelikService
pip install -r requirements.txt
python extraction_service.py
```

**Web Application:**
```bash
cd web
go mod download
go run .
```

**Neo4j:**
Install Neo4j locally or use Docker:
```bash
docker run -p 7474:7474 -p 7687:7687 -e NEO4J_AUTH=none neo4j:latest
```

### Environment Variables

- `RELIK_URL`: URL of the ReLiK service (default: `http://127.0.0.1:5000`)
- `NEO4J_URI`: Neo4j connection URI (default: `bolt://localhost:7687`)

## Project Structure

```
RelExGraph/
├── RelikService/          # Python Flask service for relation extraction
│   ├── Dockerfile
│   ├── extraction_service.py
│   └── requirements.txt
├── web/                   # Go web application
│   ├── Dockerfile
│   ├── main.go           # Main application logic
│   ├── neo4j.go          # Neo4j client
│   ├── relik.go          # ReLiK API client
│   ├── static/           # CSS and static assets
│   └── templates/        # HTML templates
├── docker-compose.yml     # Docker Compose configuration
└── README.md
```

## Notes

- **Processing Time**: The ReLiK neural model may take several seconds to process text, especially on first run
- **Memory Usage**: The ReLiK service requires significant memory for the neural model
- **Production Use**: Change Neo4j authentication settings for production deployments
- **Model Performance**: Relation extraction quality depends on text clarity and structure. The small model produces some noise alongside correct extractions
- **Model Type**: ReLiK is a specialized neural information extraction model, not a general-purpose LLM
- **GraphRAG Foundation**: This project builds the extraction and storage components of a GraphRAG pipeline, without the final LLM querying step

## Links

- [ReLiK Model](https://github.com/SapienzaNLP/relik)
- [Neo4j Documentation](https://neo4j.com/docs/)
- [Gin Web Framework](https://gin-gonic.com/)
