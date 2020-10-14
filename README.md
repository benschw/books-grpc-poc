## grpc example service

	make


In one terminal run the server:

    ./books-grpc-poc-serve

In another, test it out with the client:

    ./books-grpc-poc-client -cmd add -title "Hello World1" -author "Bob Loblaw"
    2020/10/14 09:58:40 AddBook: id:1 title:"Hello World1" author:"Bob Loblaw"

    ./books-grpc-poc-client -cmd add -title "Hello World2" -author "Bob Loblaw"
    2020/10/14 09:58:43 AddBook: id:2 title:"Hello World2" author:"Bob Loblaw"

    ./books-grpc-poc-client -cmd list
    2020/10/14 09:58:47 FindAllBooks: id:1 title:"Hello World1" author:"Bob Loblaw"
    2020/10/14 09:58:47 FindAllBooks: id:2 title:"Hello World2" author:"Bob Loblaw"
