package main

import (
	"bytes"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/benschw/books-grpc-poc/pkg/pb/books"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"log"
	"os"
)
type RDF struct {
	XMLName xml.Name `xml:"RDF"`
	EText   []EText   `xml:"etext"`
}


type EText struct {
	XMLName xml.Name `xml:"etext"`
	Title string `xml:"title"`
	Creator string `xml:"creator"`
}
var (
	addr = flag.String("addr", "localhost:9000", "The server address in the format of host:port")
	input = flag.String("input", "catalog.rdf", "Path to input rdf file")
)

func main() {

	flag.Parse()

	// connect to grpc server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := books.NewBookServiceClient(conn)

	entries, err := LoadCatalog(*input)

	if err := BulkLoad(c, entries); err != nil {
		fmt.Errorf("%s", err)
	}
}

func LoadCatalog(input string) (RDF, error) {
	var entries RDF

	xmlFile, err := os.Open(input)
	if err != nil {
		return entries, err
	}
	defer xmlFile.Close()

	b, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return entries, err
	}

	d := xml.NewDecoder(bytes.NewReader(b))
	d.Entity = map[string]string{
		"pg": "(noun)",
		"lic": "(noun)",
		"f": "(noun)",
	}
	err = d.Decode(&entries)

	return entries, err
}

func BulkLoad(client books.BookServiceClient, entries RDF) error {

	stream, err := client.BulkAddBooks(context.Background())
	if err != nil {
		return err
	}

	done := make(chan bool)

	go func() {
		m := make(map[string]int)
		for {
			c, err := stream.Recv()
			if err == io.EOF {
				for k,v := range m {
					fmt.Printf("%d books by %s\n", v, k)
				}
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
				continue
			}
			if _, ok := m[c.Author]; !ok {
				m[c.Author] = 0;
			}
			m[c.Author]++
		}
	}()

	for _, entry := range(entries.EText) {
		if err := stream.Send(&books.Book{Author: entry.Creator, Title: entry.Title}); err != nil {
			log.Fatalf("can not send %v", err)
		}
	}

	err = stream.CloseSend()
	<-done
	return err
}

