

## Core Backend Development

- __Language__: Go (Golang)
- __Web Framework__: Standard library HTTP
- __Templating__: Go Templ
- __API Design__: RESTful APIs with JSON, potentially GraphQL
- Makefiles: make or just
## Database Technologies

### __Primary Database__:

- PostgreSQL (relational, SQL)
- Redis (in-memory key-value store)

### __Alternative Databases__:

  - MongoDB (document database)
  - TimescaleDB (time-series data)
  - Neo4j (graph database)
  - CockroachDB (distributed SQL)

## Infrastructure & Containerization

- __Containerization__: Docker
- __Orchestration__: Docker Compose (development), potentially Kubernetes (advanced)
- __Reverse Proxy__: Nginx
- __Load Balancing__: Traefik or Nginx

## Messaging & Communication

- __Message Queues__: RabbitMQ, Apache Kafka, or NATS
- __Real-time Communication__: WebSockets
- __Service Communication__: gRPC, HTTP REST
## Authentication & Authorization

- __Authentication__: OAuth2, JWT (JSON Web Tokens)
- __Authorization__: Custom ACL (Access Control List) system
- __Identity Management__: Self-built OAuth2 server or Keycloak
- __Secrets Management__: HashiCorp Vault or Consul

## Storage Solutions

- __Object Storage__: MinIO (S3-compatible local storage)
- __File Storage__: Local filesystem with Docker volumes

## CI/CD & Development Tools

- __Version Control__: Git

- __CI/CD Platform__: TBD
	  - Woodpecker CI (lightweight)
	  - Jenkins (comprehensive)
	  - GitLab CI (integrated with GitLab)

- __Infrastructure as Code__: Docker Compose, Terraform

## Monitoring & Observability

- __Metrics__: Prometheus
- __Visualization__: Grafana
- __Logging__: ELK Stack (Elasticsearch, Logstash, Kibana) or Loki
- __Distributed Tracing__: Jaeger, OpenTelemetry, or Tempo
- __Alerting__: Alertmanager
- __Health Checks__: Custom endpoints with status reporting

## Testing & Quality Assurance

- __API Testing__: Postman, Hurl, Newman
- __Load Testing__: k6, Vegeta
- __Integration Testing__: Testcontainers
- __Contract Testing__: Pact
- __Code Quality__: SonarQube Community Edition
- __Security Scanning__: Snyk, Dependabot, ClamAV

## Advanced Concepts & Patterns

- __Microservices Patterns__: CQRS, Event Sourcing, Saga pattern
- __Data Processing__: Apache Spark, Flink (advanced)
- __Workflow Orchestration__: Apache Airflow
- __Service Discovery__: Consul, etcd
- __Configuration Management__: Consul, etcd, Viper
- __CLI Tools__: Cobra for building command-line applications
- __ORM__: GORM, Ent
- __Data Visualization__: Metabase, Superset, Redash

## Self-Hosted Solutions

- __Git Platform__: GitLab Community Edition, Gitea, Gogs
- __CI/CD__: Woodpecker CI, Jenkins, GitLab CI
- __Monitoring__: All Prometheus ecosystem tools are open source
- __Logging__: ELK Stack, Loki/Promtail/Grafana stack
- __Identity__: Keycloak (if not building custom OAuth2 server)
