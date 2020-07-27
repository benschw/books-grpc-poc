## db

- postgres.conf - postgresql config 
	- harvest with `docker run -i --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > postgres.conf`
- schema.sql - initial schema to work with

Start a postgresql container with docker-compose


	# start the container
	docker-compose up -d

	# stop the container
	docker-compose down

	# connect to the container
	psql -h localhost -p 5400 -U docker postgres


## Run the service

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



## Reference Links
- https://hashinteractive.com/blog/docker-compose-up-with-postgres-quick-tips/
- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
- https://github.com/rahmanfadhil/gin-bookstore

