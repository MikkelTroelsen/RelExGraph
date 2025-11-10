# RelExGraph Docker Setup

## Quick Start

### Prerequisites
- Docker Desktop installed and running
- Docker Compose installed (comes with Docker Desktop on Windows)

### Starting the Application

1. **Build and start all services:**
   ```powershell
   docker-compose up --build
   ```

2. **Start in detached mode (background):**
   ```powershell
   docker-compose up -d
   ```

3. **View logs:**
   ```powershell
   # All services
   docker-compose logs -f

   # Specific service
   docker-compose logs -f web
   docker-compose logs -f relik
   docker-compose logs -f neo4j
   ```

4. **Stop all services:**
   ```powershell
   docker-compose down
   ```

5. **Stop and remove volumes (clean slate):**
   ```powershell
   docker-compose down -v
   ```

### Accessing the Services

- **Web Application**: http://localhost:8080
- **Neo4j Browser**: http://localhost:7474
- **Relik API**: http://localhost:5000
  - Health check: http://localhost:5000/health
  - Relations endpoint: POST to http://localhost:5000/get-relations

### Service Architecture

```
┌─────────────┐
│   Browser   │
└──────┬──────┘
       │ :8080
       ▼
┌─────────────┐
│  Web (Go)   │────────┐
└──────┬──────┘        │
       │               │
       │ :5000         │ :7687
       ▼               ▼
┌─────────────┐  ┌─────────────┐
│   Relik     │  │   Neo4j     │
│  (Python)   │  │  Database   │
└─────────────┘  └─────────────┘
```

### Development Tips

#### Rebuild a specific service
```powershell
docker-compose build web
docker-compose up -d web
```

#### Execute commands inside containers
```powershell
# Access web container shell
docker-compose exec web sh

# Access relik container shell
docker-compose exec relik bash

# Check Neo4j status
docker-compose exec neo4j cypher-shell
```

#### View container status
```powershell
docker-compose ps
```

#### Restart a service
```powershell
docker-compose restart web
```

### Troubleshooting

#### Service won't start
1. Check logs: `docker-compose logs <service-name>`
2. Verify ports aren't already in use
3. Ensure Docker Desktop is running

#### Neo4j connection issues
- Wait 10-20 seconds after starting for Neo4j to initialize
- Check health: http://localhost:7474
- View logs: `docker-compose logs neo4j`

#### Relik model download
- First startup takes longer (downloads ML model)
- Monitor: `docker-compose logs -f relik`
- The model is ~100MB and cached in the container

#### Port conflicts
If ports are in use, edit `docker-compose.yml`:
```yaml
ports:
  - "9080:8080"  # Change host port (left side)
```

### Production Considerations

1. **Enable Neo4j authentication:**
   Edit `docker-compose.yml`:
   ```yaml
   environment:
     - NEO4J_AUTH=neo4j/your-secure-password
   ```

2. **Use environment files:**
   Create `.env` file:
   ```env
   NEO4J_PASSWORD=your-password
   GIN_MODE=release
   ```

3. **Persist volumes:**
   Neo4j data is automatically persisted in Docker volumes

4. **Resource limits:**
   Add to services in `docker-compose.yml`:
   ```yaml
   deploy:
     resources:
       limits:
         cpus: '2'
         memory: 4G
   ```

### File Structure

```
RelExGraph/
├── docker-compose.yml          # Orchestration
├── .dockerignore               # Global ignore
├── RelikService/
│   ├── Dockerfile              # Python service build
│   ├── .dockerignore
│   ├── extraction_service.py
│   └── requirements.txt
└── web/
    ├── Dockerfile              # Go service build
    ├── .dockerignore
    ├── main.go
    ├── neo4j.go
    ├── relik.go
    ├── go.mod
    ├── static/
    └── templates/
```

### Next Steps

1. Start the application: `docker-compose up --build`
2. Wait for all services to be healthy (~30-60 seconds first time)
3. Open http://localhost:8080 in your browser
4. Test the relation extraction functionality

### Cleaning Up

```powershell
# Stop and remove containers
docker-compose down

# Remove volumes (database data)
docker-compose down -v

# Remove images
docker-compose down --rmi all
```
