---
title: Curriculum
level:
tags: []
created_at: 2026-01-26 07-12-12
modified_at: 2026-03-09 08-02-39
---

## Project Overview

TaskFlow is a collaborative task management application that introduces participants to advanced backend concepts through a phased approach, separating core functionality from authentication systems.
#level_basic 

## Learning Objectives
- Introduce authentication and authorization methods
	- user + password
	- oauth2
	- bearer tokens
- Managing and connecting multiple systems
	- Main application
	- Oauth2 server
	- RBAC server

## Technologies Covered
#go #postgres #sql #docker #rest #auth

## Prerequisites
- Basic Go
- Basic git usage

## Estimated Time

- 5 meetings
- 6 hours each
## Resources

- Go:
	- https://go.dev/learn/
	- https://go.dev/ref/spec
- SQL:
	- https://www.postgresql.org/
	- https://www.w3schools.com/sql/
- Docker:
	- https://www.docker.com/
- Oauth2:
	- https://auth0.com/intro-to-iam/what-is-oauth-2
	- https://datatracker.ietf.org/doc/html/rfc6749
- Authentication general:
	- https://zuplo.com/learning-center/top-7-api-authentication-methods-compared
	- https://www.authgear.com/post/what-is-user-authentication-guide-2026
	- https://www.logicmonitor.com/blog/what-are-the-different-types-of-authentication
	- https://testdriven.io/blog/web-authentication-methods








## Curriculum

### Meeting 1: TaskFlow Core Implementation

__Topics Covered:__

- Advanced Go project structure
- Database relationship modeling
- RESTful API design patterns
- Data persistence with PostgreSQL
- Input validation techniques

__Prerequisites:__

- Go environment set up
- Docker and Docker Compose installed
- PostgreSQL basics understanding

__In-Class Activities:__

- Set up TaskFlow project structure
- Design PostgreSQL schema for projects, tasks, and subtasks
- Implement basic CRUD operations for projects
- Create RESTful endpoints for project management
- Set up Docker Compose with PostgreSQL container

__Homework Assignment:__

__Task__: Extend TaskFlow core functionality

1. Implement task creation, retrieval, update, and deletion
2. Create search and filtering capabilities for tasks
3. Write unit tests for all new functionality

__Deliverables:__

- Completed project and task management endpoints
- Docker Compose configuration
- Unit tests with >80% coverage

__Learning Objectives:__

- Practice advanced database relationship modeling
- Understand RESTful API design principles
- Gain experience with Docker for development environments
- Learn proper data validation techniques

### Meeting 2: TaskFlow Error handling & testing

__Topics Covered:__

- Advanced API design patterns
  - Middlewares for logging and errors
  - Wrapping standard handlers
- Error handling in Go services
- Data aggregation and reporting
- Performance considerations with database queries
- API documentation best practices

__Prerequisites:__

- Completion of Meeting 1 homework
- Running TaskFlow project with project/task endpoints
- Basic understanding of HTTP status codes and REST principles

__In-Class Activities:__

- Implement advanced task filtering and search
- Add data aggregation for task statistics
- Create comprehensive error handling middleware

__Homework Assignment:__

__Task__: Complete core TaskFlow functionality

1. Implement task status tracking and progress reporting
2. Create data validation for all input fields

__Deliverables:__

- Fully functional core TaskFlow application

__Learning Objectives:__

- Learn proper error handling patterns
- Practice writing comprehensive tests

### Meeting 3: Authentication Systems Introduction

__Topics Covered:__

- Authentication vs Authorization concepts
- OAuth2 protocol fundamentals
- JWT token structure and usage
- Token-based authentication patterns
- Security considerations for authentication

__Prerequisites:__

- Basic understanding of HTTP headers and cookies
- Familiarity with cryptographic concepts (hashing, signing)
- Understanding of REST API concepts

__In-Class Activities:__

- Create simplified OAuth2 server implementation
- Implement JWT token generation and validation
- Build token endpoints for authentication
- Design client registration system
- Test authentication flows with curl/Postman

__Homework Assignment:__

__Task__: Implement OAuth2 flows

1. Complete Authorization Code flow implementation
2. Implement Client Credentials flow
3. Add token refresh functionality
4. Create token validation middleware
5. Write tests for all authentication endpoints

__Deliverables:__

- Working OAuth2 server implementation
- JWT token generation and validation system
- Complete authentication API
- Tests for all authentication flows

__Learning Objectives:__

- Understand OAuth2 protocol flows
- Learn JWT implementation best practices
- Gain experience with token-based authentication
- Practice secure coding principles

### Meeting 4: Authorization and RBAC Implementation

__Topics Covered:__

- Role-Based Access Control (RBAC) concepts
- Permission modeling and management
- Policy enforcement techniques
- User-role assignment patterns
- Session management basics

__Prerequisites:__

- Completion of Meeting 3 homework
- Understanding of authentication concepts
- Basic knowledge of database relationships

__In-Class Activities:__

- Design RBAC data models for users, roles, and permissions
- Implement role and permission management APIs
- Create policy enforcement points
- Build user-role assignment system
- Test authorization checks with different user roles

__Homework Assignment:__

__Task__: Complete RBAC system

1. Implement permission inheritance and role hierarchies
2. Add fine-grained access control for resources
3. Create session management for users
4. Write comprehensive tests for authorization logic
5. Document RBAC API endpoints

__Deliverables:__

- Fully functional RBAC system
- Role and permission management APIs
- Policy enforcement implementation
- Tests for all authorization scenarios

__Learning Objectives:__

- Master RBAC implementation patterns
- Learn permission modeling techniques
- Understand policy enforcement concepts
- Practice authorization system design

### Meeting 5: Integration and Refinement

__Topics Covered:__

- System integration techniques
- Code refactoring for maintainability
- Performance optimization
- Code organization principles
- Documentation and cleanup

__Prerequisites:__

- Completion of all previous homework assignments
- Working TaskFlow core functionality
- Implemented authentication and authorization systems

__In-Class Activities:__

- Integrate authentication with TaskFlow core
- Implement user assignment to projects and tasks
- Add permission checks to existing endpoints
- Refactor code for better separation of concerns
- Optimize performance and fix any issues

__Homework Assignment:__

__Task__: Final project refinement

1. Complete integration of all system components
2. Refactor and clean up codebase

__Deliverables:__

- Fully integrated TaskFlow application with authentication
- Clean, well-documented codebase

__Learning Objectives:__

- Learn system integration techniques
- Practice code refactoring practices
