package main

import (
	"fmt"
	"log"
	"os"

	loads "github.com/go-openapi/loads"
	flag "github.com/spf13/pflag"

	"github.com/eatigo/go-swagger/examples/todo-list/restapi"
	"github.com/eatigo/go-swagger/examples/todo-list/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func main() {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	var server *restapi.Server // make sure init is called

	flag.Usage = func() {
		fmt.Fprint(os.Stderr, "Usage:\n")
		fmt.Fprint(os.Stderr, "  todo-list-server [OPTIONS]\n\n")

		title := "Simple To Do List API"
		fmt.Fprint(os.Stderr, title+"\n\n")
		desc := swaggerSpec.Spec().Info.Description
		if desc != "" {
			fmt.Fprintf(os.Stderr, desc+"\n\n")
		}
		fmt.Fprintln(os.Stderr, flag.CommandLine.FlagUsages())
	}
	// parse the CLI flags
	flag.Parse()

	api := operations.NewTodoListAPI(swaggerSpec)
	// get server with flag values filled out
	server = restapi.NewServer(api)

	defer server.Shutdown()

	server.ConfigureAPI()
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
