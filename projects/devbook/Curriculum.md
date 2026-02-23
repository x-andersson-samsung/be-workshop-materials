## Project Overview

DevBook is a simple application where developers can save, categorize, and manage useful resources like articles, tutorials, tools, and libraries. Users can create resources with titles, URLs, descriptions and tags.
#level_basic 

## Learning Objectives
- Introduce listeners to basic back end development flow.
- Provide fundamentals for future projects

## Technologies Covered
#go #postgres #sql #docker #rest

## Prerequisites
- Basic programming knowledge (any language)
- Basic git usage

## Estimated Time

- 5 meetings
- 3 hours each
## Resources

- Go:
	- https://go.dev/learn/
	- https://go.dev/ref/spec
- SQL:
	- https://www.postgresql.org/
	- https://www.w3schools.com/sql/
- Docker:
	- https://www.docker.com/

## Curriculum

### Meeting 1: Go Basics

**Topics Covered:**
- Go syntax
- Go project structure and modules
- Go tooling
- Structs and data modeling
- In-memory storage with maps/slices
- Running Go applications

__Prerequisites__:
- (for Windows) setup WSL
- Set up Go environment

**In-Class Activities:**
- Create basic project structure
- Build simple app with simple terminal UI
- Implement in-memory storage for resources

__Homework Assignment:__

**Task**: Simple go exercises
1. Fix non compiling programs
2. Write tests
3. Implementing basic algorithms

**Deliverables**: 
- Fixed files

**Learning Objectives**:
- Practice Go syntax and structuring
- Get comfortable with basic data manipulation


### Meeting 2: REST API

__Topics Covered:__
- RESTful API principles
- HTTP methods (GET, POST, PUT, DELETE)
- JSON request/response handling
- Error handling and HTTP status codes
- Query parameters and path variables
- Input validation and sanitization

__In-Class Activities:__
- Add http server to application
- Add support for creating and retrieving single entries
- Implement a search functionality that finds resources by title (partial match)
- Implement proper HTTP status codes
- Create a simple HTML page (static file) that lists all resources
- Add query parameter support (filtering, pagination)
- Implement input validation
- Create consistent error response format

__Homework Assignment:__

**Task**: Build a complete REST API
1. Implement missing functions (update, delete)
2. Implement pagination for resource listing
3. Add filtering by tags
4. Create search endpoint with query parameters
5. Add sorting options (by date, title, etc.)
6. Implement proper validation for all endpoints
7. **Bonus**: Add bulk operations (create multiple resources)

**Deliverables**:
- Complete REST API with all endpoints
- API documentation (simple markdown file)
- (optional) Test scripts showing all functionality

**Learning Objectives**:
- Learn REST API design principles
- Practice basic HTTP handling
- Learn input validation techniques
- Understand API documentation basics

### Meeting 3: SQL Database Integration

__Topics Covered:__
- PostgreSQL basics and setup
- Database design for resources and tags
- Go database/sql package
- pgx driver for PostgreSQL
- Environment variables and configuration
- Basic CRUD operations with SQL

__Prerequisites__:
- Install PostgreSQL

__In-Class Activities:__
- Design database schema for resources
- Set up database connection in Go
- Implement CREATE and READ operations
- Environment configuration with .env files

__Homework Assignment:__

**Task**: Complete the database integration
1. Implement UPDATE and DELETE operations for resources
2. Create a function to list resources by tag
3. Add proper error handling for database operations

**Deliverables**:
- Database schema diagram
- SQL migration scripts
- Updated Go application with full CRUD
- Test data insertion script

**Learning Objectives**:
- Understand database relationships
- Practice SQL operations
- Learn proper database connection management
- Handle errors in database operations

### Meeting 4: Docker & Containerization

__Topics Covered:__
- Docker basics and concepts
- Dockerfile creation
- Docker Compose for multi-container setup
- Environment configuration
- Data persistence with volumes
- Networking between containers

__Prerequisites__:
- Install Docker

__In-Class Activities:__
- Create Dockerfile for Go application
- Set up Docker Compose with Go app and PostgreSQL
- Configure environment variables
- Implement data persistence
- Test containerized application

__Homework Assignment:__

**Task**: Production-ready containerization
1. Add FE container
2. Add health check endpoints and Docker health checks
3. Add environment-specific configurations (dev/staging/prod)

**Deliverables**:
- Complete docker-compose.yml with all services
- Environment configuration files
- Documentation for running the application

**Learning Objectives**:
- Docker containerization
- Understand multi-container applications
- Learn environment configuration management
- Practice deployment and operations concepts

### Meeting 5: Testing & Best Practices (Optional)

__Topics Covered:__
- Unit testing in Go
- Integration testing
- API testing with tools
- Documentation generation
- Security considerations
- Performance optimization
- Code quality tools (golintci, go vet, go sec, go fmt)

__In-Class Activities:__
- Write unit tests for business logic
- Implement integration tests for API endpoints
- Set up automated testing
- Generate API documentation

__Homework Assignment:__

**Task**: Quality assurance and documentation
1. Write comprehensive unit tests (80%+ coverage)
2. Generate OpenAPI/Swagger documentation
3. Add logging and basic monitoring
4. **Bonus**: Set up a simple CI pipeline (GitHub Actions/GitLab CI)

**Deliverables**:
- Test suite with coverage report
- API documentation
- CI configuration file

**Learning Objectives**:
- Testing strategies in Go
- Learn documentation best practices
- Understand security fundamentals
- Practice quality assurance techniques

## Project Features Summary:

### Core Functionality:
- Create, read, update, delete resources
- Categorize resources
- Tag resources for better organization
- Search and filter capabilities
- RESTful API design
- Containerized deployment

### Data Model:
```
Resources:
- ID (int)
- Title (string)
- URL (string)
- Description (text)
- CreatedAt (timestamp)
- UpdatedAt (timestamp)

Tags:
- ID (int)
- Name (string)

ResourceTags: (junction table)
- ResourceID (foreign key)
- TagID (foreign key)
```

### API Endpoints:
- `GET /api/resources` - List all resources (with pagination/filtering)
- `GET /api/resources/{id}` - Get specific resource
- `POST /api/resources` - Create new resource
- `PUT /api/resources/{id}` - Update resource
- `DELETE /api/resources/{id}` - Delete resource
- `GET /api/resources?tag={tag}` - Filter by tag
- `GET /api/resources?search={term}` - Search resources
