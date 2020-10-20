## grpc example service

	make


In one terminal run the server:

    ./books-grpc-poc-server

In another, test it out with the client:

    ./books-grpc-poc-client -cmd add -title "Hello World" -author "Bob Loblaw"
    Book Added: id:1  title:"Hello World"  author:"Bob Loblaw"

    ./books-grpc-poc-client -cmd add -title "Hello World 2" -author "Bob Loblaw"
    Book Added: id:2  title:"Hello World 2"  author:"Bob Loblaw"

