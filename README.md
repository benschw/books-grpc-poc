

## db
Start a postgresql container with docker-compose

	docker-compose up -d


## Run the UI 
https://github.com/benschw/books-ui-poc
(this will proxy API requests to localhost:8080 to match the below default service usage)

	git clone git@github.com:benschw/books-ui-poc.git
	cd books-ui-poc
	yarn start

## Run the service
build & run the app

	go build
	./books-poc

## Interact with the API

	$ curl -X POST localhost:8080/book -d '{"title": "hello world", "author": "ben schwartz"}'
	{"data":{"id":1,"title":"hello world","author":"ben schwartz"}}
	
	$ curl -X POST localhost:8080/book -d '{"title": "hello world", "author": "ben schwartz"}'
	{"data":{"id":2,"title":"hello world","author":"ben schwartz"}}
	
	$ curl -X POST localhost:8080/book -d '{"title": "hello world", "author": "ben schwartz"}'
	{"data":{"id":3,"title":"hello world","author":"ben schwartz"}
	
	$ curl -s localhost:8080/book | jq .
	{
	  "data": [
	    {
	      "id": 1,
	      "title": "hello world",
	      "author": "ben schwartz"
	    },
	    {
	      "id": 2,
	      "title": "hello galaxy",
	      "author": "ben schwartz"
	    },
	    {
	      "id": 3,
	      "title": "hello world",
	      "author": "ben schwartz"
	    }
	  ]
	}}

	$ curl -X PUT localhost:8080/book/1 -d '{"title": "hello galaxy", "author": "benjamin schwartz"}'
	{"data":{"id":1,"title":"hello galaxy","author":"benjamin schwartz"}}

	$ curl -X DELETE localhost:8080/book/2

	$ curl -s localhost:8080/book | jq .
	{
	  "data": [
	    {
	      "id": 1,
	      "title": "hello galaxy",
	      "author": "benjamin schwartz"
	    },
	    {
	      "id": 3,
	      "title": "hello world",
	      "author": "ben schwartz"
	    }
	  ]
	}}



## Notes
- postgres.conf
	- postgresql config harvest with `docker run -i --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > postgres.conf`
- schema.sql
	- initial schema to work with - (implement migrations?)
- once running...
	- connect with cli: `psql -h localhost -p 5400 -U docker postgres`

- https://hashinteractive.com/blog/docker-compose-up-with-postgres-quick-tips/
- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
- https://github.com/rahmanfadhil/gin-bookstore
- https://github.com/jackc/pgx/blob/master/examples/todo/main.go

