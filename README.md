## books api
- https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/
- https://github.com/rahmanfadhil/gin-bookstore

### examples

	$ curl -X POST localhost:8080/books -d '{"title": "hello world", "author": "ben schwartz"}'
	{"data":{"id":1,"title":"hello world","author":"ben schwartz"}}
	
	$ curl -X POST localhost:8080/books -d '{"title": "hello world", "author": "ben schwartz"}'
	{"data":{"id":2,"title":"hello world","author":"ben schwartz"}}
	
	$ curl -X POST localhost:8080/books -d '{"title": "hello world", "author": "ben schwartz"}'
	{"data":{"id":3,"title":"hello world","author":"ben schwartz"}
	
	$ curl -s localhost:8080/books | jq .
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

## db
- https://hashinteractive.com/blog/docker-compose-up-with-postgres-quick-tips/



copy config

	docker run -i --rm postgres cat /usr/share/postgresql/postgresql.conf.sample > postgres.conf

start container

	docker-compose up -d

connect with cli

	psql -h localhost -p 5400 -U docker postgres



