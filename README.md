## grpc example service

	make


In one terminal run the server:

    ./books-grpc-poc-server

In another, test it out with the client:

    ./books-grpc-poc-client -cmd add -title "Hello World" -author "Bob Loblaw"
    Book Added: id:1  title:"Hello World"  author:"Bob Loblaw"

    ./books-grpc-poc-client -cmd add -title "Hello World 2" -author "Bob Loblaw"
    Book Added: id:2  title:"Hello World 2"  author:"Bob Loblaw"

    ./books-grpc-poc-client -cmd list
    id:1  title:"Hello World"  author:"Bob Loblaw"
    id:2  title:"Hello World 2"  author:"Bob Loblaw"
    
    
Download the sample RDF catalog from [Project Gutenberg](https://www.gutenberg.org/ebooks/offline_catalogs.html)

    wget -qO- "https://www.gutenberg.org/cache/epub/feeds/catalog.rdf.zip" | tar xOvz -
    
    
Bulk load them into the books service

    go run cmd/bulk_load_client/main.go -input ./catalog.rdf
    # client will output a summary of how many books by each author were added