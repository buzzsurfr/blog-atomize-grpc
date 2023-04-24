# blog-atomize-grpc
Blog showing how to break a monolith into microservices using gRPC

## Demo

### Monolith

The monolith application is in `cmd/monolith` and contains all the functions and services within that folder. The loyalty and orders backends are in _loyalty.go_ and _orders.go_ (respectively) and the main API server as well as function handlers are all in _main.go_.

To start the API server, run:

```
go run cmd/monolith/*.go
```

Then test using postman.

### Microservices

The microservices application is separated into:

* a loyalty service, in `cmd/microloyalty`,
* an orders service, in `cmd/microorders`, and
* the main API server, in `cmd/microlith`.

So that we can monitor all of the microservices at once, open up 3 terminals (windows or tabs).

Start the loyalty service.

```
go run cmd/microloyalty/main.go
```

Start the orders service.

```
go run cmd/microorders/main.go
```

Finally, start the main API server.

```
go run cmd/microlith/main.go
```

Then test using postman.

## Test using Postman

The Postman collection is available in [Atomize-gRPC.postman_collection.json](Atomize-gRPC.postman_collection.json).

Run the monolith or microservices demo, and test the two included services using Postman. No changes are required to test using Postman.

### Monitoring the Monolith

### Monitoring the microservices