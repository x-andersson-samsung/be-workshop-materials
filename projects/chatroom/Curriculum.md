---
title: Curriculum
level:
tags: []
created_at: 2026-01-26 07-12-12
modified_at: 2026-03-09 09-32-01
---

## Project Overview

ChatRoom is a real-time messaging application that teaches participants backend concepts. Starting with HTTP polling mechanisms, moving into WebSockets for real-time communication, file sharing capabilities with MinIO, and persistent data management with Redis. 
## Learning Objectives

1. Implement real-time communication using both HTTP polling and WebSockets
2. Design service-oriented architecture with clean separation of concerns
3. Integrate multiple storage solutions including in-memory, Redis, and object storage
4. Handle concurrent connections and resource management in Go applications
5. Apply message queuing and pub/sub patterns for scalable systems
6. Practice performance optimization and system refactoring techniques

## Technologies Covered
#go #redis #docker #rest

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
- Redis:
	- https://redis.io/

## Curriculum

### Meeting 1: ChatRoom HTTP Core Implementation

__Topics Covered:__
- HTTP polling for real-time-like behavior
- Service and repository layer architecture
- RESTful API design patterns
- In-memory data structures for storage
- Data modeling for chat applications

__Prerequisites:__
- Go environment set up
- Docker and Docker Compose installed

__In-Class Activities:__
- Set up ChatRoom project structure with service/repository layers
- Implement in-memory storage using maps and slices
- Design data structures for users, rooms, and messages
- Create HTTP endpoints for chat operations
- Implement message polling mechanism

__Homework Assignment:__

__Task__: Extend ChatRoom HTTP functionality
1. Create room creation and management APIs
2. Add message filtering and simple pagination
3. Implement basic error handling

__Deliverables:__
- Working HTTP-based chat application with in-memory storage
- Service and repository layer implementation
- Basic API documentation

__Learning Objectives:__
- Practice service/repository architecture patterns
- Understand HTTP-based polling for real-time behavior
- Learn proper data modeling for chat applications
- Gain experience with in-memory data structures

### Meeting 2: WebSocket Integration

__Topics Covered:__
- WebSocket protocol fundamentals
- Upgrading HTTP connections to WebSockets
- Real-time message broadcasting
- Concurrent connection management
- Handler separation for HTTP and WebSocket

__Prerequisites:__
- Completion of Meeting 1 homework
- Running ChatRoom with HTTP polling
- Basic understanding of HTTP and TCP concepts

__In-Class Activities:__
- Implement WebSocket handlers using existing service layer
- Create connection upgrade mechanism
- Build real-time message broadcasting system
- Handle concurrent WebSocket connections
- Test real-time messaging vs HTTP polling

__Homework Assignment:__

__Task__: Complete WebSocket integration
1. Add user presence detection with WebSockets
2. Implement room joining/leaving via WebSocket
3. Create message acknowledgment system
4. Add graceful connection handling
	1. Connection cleanup
	2. Heartbeat/ping
	3. Error handling

__Deliverables:__
- Fully functional WebSocket chat system
- Integrated service layer with both HTTP and WebSocket
- Connection management implementation
- Integration tests for real-time features

__Learning Objectives:__
- Understand WebSocket protocol implementation
- Learn connection upgrade patterns
- Practice real-time message broadcasting
- Master concurrent connection handling

### Meeting 3: File Sharing with MinIO

__Topics Covered:__
- Object storage concepts
- MinIO integration for file storage
- File metadata management
- Secure file upload/download patterns
- Content-Type handling for different file types

__Prerequisites:__
- Completion of Meeting 2 homework
- Working ChatRoom with WebSocket integration
- Basic understanding of file handling concepts

__In-Class Activities:__
- Set up MinIO container with Docker Compose
- Implement file upload endpoint with MinIO integration
- Create file metadata storage in-memory
- Build file download and sharing mechanisms
- Test file operations with different file types

__Homework Assignment:__

__Task__: Complete file sharing functionality
1. Add file type validation and size limits
2. Create file expiration and cleanup mechanisms
3. Write tests for file sharing operations

__Deliverables:__
- Fully integrated file sharing system with MinIO
- File metadata management

__Learning Objectives:__
- Learn object storage integration patterns
- Understand file metadata management
- Practice secure file handling techniques
- Master content-type handling for diverse files

### Meeting 4: Redis Data Persistence

__Topics Covered:__
- Redis data structures for chat applications
- Data migration from in-memory to Redis
- Pub/Sub messaging patterns
- Persistence strategies in Redis
- Connection management with Redis

__Prerequisites:__
- Completion of Meeting 3 homework
- Working ChatRoom with file sharing

__In-Class Activities:__
- Set up Redis container with Docker Compose
- Migrate in-memory storage to Redis data structures
- Implement Redis pub/sub for message broadcasting
- Create data persistence layers
- Test data consistency between components

__Homework Assignment:__

__Task__: Complete Redis integration
1. Implement user session storage in Redis
2. Add message history persistence
3. Create room state management in Redis

__Deliverables:__
- Fully integrated Redis persistence layer
- Data migration from in-memory to Redis
- Session and state management in Redis
- Tests for all Redis operations

__Learning Objectives:__
- Understand Redis data structures for applications
- Learn data migration patterns
- Practice persistence strategies
- Master Redis connection management

### Meeting 5: Performance Optimization & Refinement

__Topics Covered:__
- Performance profiling in Go
- Memory optimization techniques
- Connection pooling strategies
- Load testing methodologies
- Code refactoring for maintainability

__Prerequisites:__
- Completion of all previous homework assignments
- Fully functional ChatRoom application
- Understanding of performance concepts

__In-Class Activities:__
- Profile ChatRoom application for bottlenecks
- Optimize data structures and algorithms
- Implement connection pooling for Redis
- Conduct load testing with multiple users
- Refactor code for better organization

__Homework Assignment:__

__Task__: Final ChatRoom optimization
1. Complete performance optimizations based on profiling
2. Add comprehensive logging and monitoring hooks
3. Refactor and clean up codebase
4. Create deployment documentation
5. Prepare benchmarking report

__Deliverables:__
- Fully optimized ChatRoom application
- Performance benchmarking results
- Clean, well-documented codebase
- Deployment documentation
- Benchmarking and optimization report

__Learning Objectives:__
- Learn performance profiling techniques
- Master memory and connection optimization
- Understand load testing methodologies
- Practice system refactoring and documentation
