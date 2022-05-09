package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"cyoa"
)

func main() {
	port := flag.Int("port", 3000, "the port to start the CYOA web application on")
	filename := flag.String("file", "gopher.json", "the JSON file with the story")
	flag.Parse()
	fmt.Printf("Using the story %s\n", *filename)

	f, err := os.Open(*filename)
	if err != nil {
		panic(err)
	}

	story, err := cyoa.JsonStory(f)
	if err != nil {
		panic(err)
	}

	// tpl := template.Must(template.New("").Parse("Custom template test"))
	// h := cyoa.NewHandler(story, cyoa.WithTemplate(tpl))

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the application on port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), h))
}
