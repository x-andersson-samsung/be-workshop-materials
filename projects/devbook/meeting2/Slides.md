---
title: "Devbook "
level: basic
tags: []
created_at: 2026-01-26 07-12-12
modified_at: 2026-02-27 07-13-05
slideNumber: "true"
---

 %% Required for proper codeblock width %%
<style>

li,p {
	font-size: 32px;
}

code {
    font-size: 16px;
    line-height: normal;
}

/* left-align all content in Slides */
.reveal .slides {
    text-align: left;
}

</style>

%% Start of slides %%

# Devbook 
## Meeting 2

---

<grid drag="100 10" drop="0 0" align="left">
### Plan
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Introduce REST API concepts
- Add http server to our app
- Basic resource operations (CRUD)
- Go templates
</grid>

---

# REST

--

<grid drag="100 10" drop="0 0" align="left" >
### What is REST?
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
REST (Representational State Transfer) is a set of rules for building web services that use HTTP methods (GET, POST, PUT, DELETE) to work with data. It's stateless, meaning each request contains all needed information, and uses standard web protocols making it simple and widely adopted for APIs.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Rules of REST
**_1. Uniform interface_**
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
All requests made via API must respect that API’s formatting rules. 

No matter what client is making the request, it must put each piece of information where every other client would put it.

One example is the URL used for identifying resources via HTTP.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Rules of REST
**_2. Client-server separation_**
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
REST APIs require that client and server applications are totally independent of each other.

The client need only know the full name of the resource it wants.

The only knowledge client and server have of each other is that exchanged by API transactions.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Rules of REST
**_3. Statelessness_**
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">


Each client request must contain all the information needed to process it, and the server need not hold any information about that request once it has been received.

There is no concept of a session in REST API design and the server is stateless with respect to any particular client.
</grid>

--

<grid drag="100 5" drop="0 0" align="left" >
### Rules of REST
**_4. Cacheability_**
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Resources should be cacheable at one or more points within or between client and server.

In the case of the server, if a particular resource has been served, and it is likely to be requested again within a certain time, it should be cached for a more rapid subsequent response. 

The server should indicate via the API whether a resource can be safely cached locally at the client, including the lifetime of the data, where appropriate.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Rules of REST
**_5. Layered system architecture_**
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
A client can make no assumptions about whether it is communicating directly with a server holding a particular resource, or whether it is being served by an intermediate such as a service broker, load balancer, content delivery system or other subsystem closer to the client than the server.

This provides system and infrastructure designers considerable flexibility to maximise the efficiency and reliability of request satisfaction across the global wired and wireless infrastructure.
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Rules of REST
**_6. Code on demand_**
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
While REST APIs can and frequently do serve only data for consumption by the client, it is increasingly common for code to be delivered to run on the client, such as Java objects or Javascript web apps. If this is implemented, then such code can only be run on demand by the client.
</grid>

---

# Best practices

--

<grid drag="100 10" drop="0 0" align="left" >
### Singleton vs Collection
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
A resource can be a singleton or a collection

For example **_customers_** is a collection while **_customer_** is a singleton

```
/customers      // a collection resource
/customers/{id} // singleton resource
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Sub collections
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Subcollections should extend the location.

For example if **_customer_** has **_account_** resources.

```
/customers/{id}/accounts          // all accounts belonging to customer {id}
/customers/{id}/accounts/{acc_id} // a single account of customer {id}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Best practices
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Use nouns to represent resources
- Use singular for singletons and plural for collections
- Use forward slash (/) to indicate hierarchical relationships
- Do not use file extensions
- Do not use CRUD function names in URIs
- Use query component to filter URI collection
- [And many more](https://restfulapi.net/rest-api-best-practices/)

</grid>

note: 
Do not use file extensions
For file extensions you should rely on `Content-Type` headers

Do not use CRUD function names in URIs
You should use HTTP request methods instead

---

# Go

---

# Server

--

<grid drag="100 10" drop="0 0" align="left" >
### Http servers - handlers
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Handlers process incoming requests.

```go
// from http package

// Function
type HandlerFunc func(ResponseWriter, *Request)

// Interface
type Handler interface {  
  ServeHTTP(ResponseWriter, *Request)  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Http servers - setup
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Go standard library provides a package for setting up both HTTP clients and servers.

###### Simplest
```go [|5-7|11-14]
import "net/http"

// Handling function for std must have the signature:  
// func(http.ResponseWriter, *http.Request)  
func indexHandler(w http.ResponseWriter, r *http.Request) {  
    w.Write([]byte("Hello World!"))  
}  
  
func main() {  
    // Start server  
    if err := http.ListenAndServe(
	    ":8080",
	     http.HandlerFunc(indexHandler)
	 ); err != nil {  
       panic(err)  
    }  
}
```
</grid>

--


<grid drag="100 10" drop="0 0" align="left" >
### Http servers - setup
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Go standard library provides a package for setting up both HTTP clients and servers.

###### Simple
```go [|3-7|9-10|12-15|17-20]
import "net/http"

// Handling function for std must have the signature:  
// func(http.ResponseWriter, *http.Request)  
func indexHandler(w http.ResponseWriter, r *http.Request) {  
    w.Write([]byte("Hello World!"))  
}  
func main() {  
    // Setup HTTP multiplexer  
    mux := http.NewServeMux()  
  
    // Add routes  
    mux.HandleFunc("/", indexHandler)  
    mux.HandleFunc("GET /users", userListHandler)  
    mux.HandleFunc("GET /users/{id}", userGetHandler)  
  
    // Start server  
    if err := http.ListenAndServe(":8080", mux); err != nil {  
       panic(err)  
    }  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Http servers - setup
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Go standard library provides a package for setting up both HTTP clients and servers.

###### Advanced 1 - Handler
```go [|7-11|13-20]
import "net/http"

type  UserStore interface {
	Get(id string) (User, err)
}

// Instead of passing functions we can use our own handler as long as
// it implements `http.Handler` interface
type UserHandler struct {
	user  UserStore
}

func (usrH UserHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// Get id from path
	id := r.PathValue("id")

	 user, err := usrH.user.Get(id)
	 // handle error or return user
	 // ...
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Http servers - setup
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Go standard library provides a package for setting up both HTTP clients and servers.

###### Advanced 2 - Server
```go [|7-8|9|10|11-14|15|16-17]
import "net/http"

func main() {  
	// http.Server gives us a lot of fields to modify server behaviour 
	// instead of using the default one
    srv := &http.Server{  
       Addr:                         "localhost:8080",  
       Handler:                      &UserHandler{},  
       DisableGeneralOptionsHandler: false,  
       TLSConfig:                    nil,  
       ReadTimeout:                  0,  
       ReadHeaderTimeout:            0,  
       WriteTimeout:                 0,  
       IdleTimeout:                  0,  
       MaxHeaderBytes:               0,  
       BaseContext:                  nil,  
       ConnContext:                  nil,  
    }  
    srv.ListenAndServe()                          // !!In code handle error
    srv.ListenAndServeTLS("cert.pem", "key.pem")  // !!In code handle error
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Requests
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Object: **_http.Request_**

```go
type Request struct {
	Method string
	URL *url.URL
	Header Header

	Body io.ReadCloser
	Form url.Values
	PostForm url.Values
	MultipartForm *multipart.Form

	ctx context.Context
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Requests - form data
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go
func handler(w http.ResponseWriter, r *http.Request) {
	// /users/{id}  
	r.PathValue("id")  
  
	// Parse form  
	err := r.ParseForm() // remember to handle error  

	// Url query + post form
	r.Form.Get("key")  
	r.Form["key"]  

	// Only post form
	r.PostForm.Get("key")  
	r.PostForm["key"]
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Requests - body
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
**_io.ReadAll_** is a helper function for getting all data from an **_io.Reader_**
```go
func handler(w http.ResponseWriter, r *http.Request) {

	data, err := io.ReadAll(r.Body)
	
	if err != nil { // handle error }
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Response
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Interface: **_http.ResponseWriter_**

```go
type ResponseWriter interface {
	// Header returns the header map that will be sent by
	// [ResponseWriter.WriteHeader]. 
	Header() Header

	// Write writes the data to the connection as part of an HTTP reply.
	//
	// If [ResponseWriter.WriteHeader] has not yet been called, Write calls
	// WriteHeader(http.StatusOK) before writing the data.
	Write([]byte) (int, error)

	// WriteHeader sends an HTTP response header with the provided
	// status code.
	WriteHeader(statusCode int)
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Response - example
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
**_http_** package contains contants for HTTP methods and status codes.

```go
func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	w.WriteHeader(http.StatusOK)
	
	w.Write(`{"status": "ok"}`)
}
```
</grid>

---

# Client side

--


<grid drag="100 10" drop="0 0" align="left" >
### Basic requests
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
**_http_** package contains helper funcions for most common requests.

```go
Head(url string) (resp *Response, err error)

Get(url string) (resp *Response, err error)

Post(url, contentType string, body io.Reader) (resp *Response, err error)
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Basic requests - example GET
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go[]
package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	req, err := http.Get("https://echo.free.beeceptor.com/sample-request")  
	if err != nil {  
	    panic(err)  
	}  
	  
	body, err := io.ReadAll(req.Body)  
	if err != nil {  
	    panic(err)  
	}  
	  
	fmt.Println(string(body))
}
```
</grid>


--

<grid drag="100 10" drop="0 0" align="left" >
### Basic requests - example POST
</grid>

<grid drag="100 90" drop="0 10" align="left" justify-content="center">
```go[]
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)
func main() {
	req, err := http.Post(
		"https://echo.free.beeceptor.com/sample-request",
		"application/json",
		strings.NewReader(`{"key": "value"}`),
	)
	if err != nil {panic(err)}

	body, err := io.ReadAll(req.Body)
	if err != nil {panic(err)}

	fmt.Println(string(body))
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### More customization
</grid>

<grid drag="100 90" drop="0 10" align="left" justify-content="center">
```go[]
package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)
func main() {
	req, err := http.NewRequest(http.MethodGet, "someURL", nil)

	// We can modify headers
	req.Header.Set("Content-Type", "application/json")

	// We can set auth
	req.SetBasicAuth("username", "password")
	

	resp, err := http.DefaultClient.Do(req)
}
```
</grid>

---

# JSON

--

<grid drag="100 10" drop="0 0" align="left" >
### JSON handling
</grid>

<grid drag="100 90" drop="0 10" align="left" justify-content="center">
Go provides builtin handling with **_encoding/json_** package and tags.
```go[]
type User struct {
	FirstName string `json:"firstname"`
	LastName string `json:"lastname"`
	
	SkippedField string `json:"-"`	
}

const data = `{"firstname": "John", "lastname":"Doe", "address":"123"}`

// Unmarshalling (Decoding)
var u User
err = json.Unmarshall(data, &u)

// Marshalling (Encoding)
data, err = json.Marshall(u)
```
</grid>




---
# Exercises

--

<grid drag="100 10" drop="0 0" align="left" >
### Exercise 1
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Write an "echo" server with its client.

- **_GET /echo_** return "Hello World!".
- **_POST /echo_** return provided body.

</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Exercise 2
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Write 2 handlers. Each will receive 2 query parameters **_a_** and **_b_**.

- AddHandler - returning sum of a & b
- SubHandler - returning difference of a & b

You can assume input will be correct and skip error checking.

Hint: 

- **_strconv.Atoi_** & **_strconv.ItoA_** will help convert types
- You can cast **_string_** to **_[]byte_** to use in **_w.Write_**

</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Exercise 3
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Write a handler processing body parameters.

- AddHandler - returning sum of all values.

Parameters will be provided in JSON body:
```json
{
	"values": [1,2,3,4]
}
```

Expected response format:
```json
{
	"result": 10
}
``` 
</grid>

---

# Back to Devbook

--

<grid drag="100 10" drop="0 0" align="left" >
### Preparations
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Add **_ID_** to Item:
```go []
type Item struct {  
    ID          int    `json:"id"`  
    Name        string `json:"name"`  
    Description string `json:"description"`  
    URL         string `json:"URL"`  
}
```

Modify store to be based on ID:
```go
type Store struct {  
    nextID int  
    data   map[int]Item  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Preparations - store functions
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
func (store *Store) GetByID(id int) (Item, bool) {  
    item, ok := store.data[id]  
    return item, ok  
}  
  
func (store *Store) GetByName(name string) (Item, bool) {  
    for _, item := range store.data {  
       if item.Name == name {  
          return item, true  
       }  
    }  
  
    return Item{}, false  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Preparations - store functions
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
func (store *Store) Add(item Item) {  
    item.ID = store.nextID  
    store.nextID++  
  
    store.data[item.ID] = item  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Preparations - store functions
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
func (store *Store) DeleteByID(id int) {  
    delete(store.data, id)  
}  
  
func (store *Store) DeleteByName(name string) {  
    for id, item := range store.data {  
       if item.Name == name {  
          delete(store.data, id)  
          return  
       }  
    }  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Preparations - store functions
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
func (store *Store) List() []Item {  
    arr := make([]Item, 0, len(store.data))  
    for _, item := range store.data {  
       arr = append(arr, item)  
    }  
  
    slices.SortFunc(arr, func(i, j Item) int {  
       if i.ID == j.ID {  
          return 0  
       }  
       if i.ID > j.ID {  
          return 1  
       }  
       return -1  
    })  
  
    return arr  
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Http Handler
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
Lets' create new package **_pkg/api_**
```go []
type ItemHandler struct {  
    Store *devbook.Store  
}

func (i *ItemHandler) Register(r *http.ServeMux) {  
    r.HandleFunc("POST /items", i.Create)  
    r.HandleFunc("GET /items", i.List)  
    r.HandleFunc("PUT /items/{id}", i.Update)  
    r.HandleFunc("DELETE /items/{id}", i.Delete)  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Http Handler - List
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
func (i *ItemHandler) List(w http.ResponseWriter, r *http.Request) {  
    items := i.Store.List()  
  
    payload, err := json.Marshal(items)  
    if err != nil {  
       http.Error(w, err.Error(), http.StatusInternalServerError)  
       return  
    }  
  
    w.WriteHeader(http.StatusOK)  
    w.Write(payload)  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Http Handler - Create
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
func (i *ItemHandler) Create(w http.ResponseWriter, r *http.Request) {  
    var item devbook.Item  
    err := json.NewDecoder(r.Body).Decode(&item)  
    if err != nil {  
       http.Error(w, err.Error(), http.StatusBadRequest)  
       return  
    }  
  
    i.Store.Add(item)  
  
    w.WriteHeader(http.StatusNoContent)  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Http Handler - Delete
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go []
func (i *ItemHandler) Delete(w http.ResponseWriter, r *http.Request) {  
    idStr := r.PathValue("id")  
    if idStr == "" {  
       http.Error(w, "Missing id", http.StatusBadRequest)  
       return  
    }  
  
    id, err := strconv.Atoi(idStr)  
    if err != nil {  
       http.Error(w, err.Error(), http.StatusBadRequest)  
       return  
    }  
  
    i.Store.DeleteByID(id)  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Adapting main - getting index.html
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go[]
// Option 1 - embed the index file into the binary

//go:embed static/items.html  
var indexData []byte  

// Option 2 - read the file to memory on program start
func readIndexFile(path string) ([]byte, error) {  
    return os.ReadFile(path)  
}
```
</grid>

--

<grid drag="100 10" drop="0 0" align="left" >
### Adapting main - server setup
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
```go[]
func setupAndRunServer(store *devbook.Store, indexData []byte) error {  
    mux := http.NewServeMux()  
  
    mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {  
       w.Write(indexData)  
    })  
  
    itemHandler := api.ItemHandler{Store: store}  
    itemHandler.Register(mux)  
  
    fmt.Println("Starting server on " + ServerAddress)  
    return http.ListenAndServe(ServerAddress, mux)  
}
```
</grid>

---

# Homework

--

<grid drag="100 10" drop="0 0" align="left" >
### Assignments
</grid>

<grid drag="100 85" drop="0 10" align="left" justify-content="center">
- Add support for editing items on **_"PUT /items/{id}"_**
- Add pagination support to **_List_**
	- Hint: FE is sending **_offset_** and **_limit_** query values
- (optional) Add URL validation to **_Create_** and **_Update_**
    - Should return different error to FE
    - For validation see: [url.Parse](https://pkg.go.dev/net/url#Parse)
- Complete exercises
</grid>

---