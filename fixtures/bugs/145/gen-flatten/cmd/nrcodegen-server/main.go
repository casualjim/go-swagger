package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/loads/fmts"
	flags "github.com/jessevdk/go-flags"

	"github.com/go-swagger/go-swagger/fixtures/bugs/145/gen-flatten/restapi"
	"github.com/go-swagger/go-swagger/fixtures/bugs/145/gen-flatten/restapi/operations"
)

// This file was generated by the swagger tool.
// Make sure not to overwrite this file after you generated it because all your edits would be lost!

func init() {
	loads.AddLoader(fmts.YAMLMatcher, fmts.YAMLDoc)
}

func main() {

	server := restapi.NewServer(nil)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "To-do Demo"
	parser.LongDescription = "This OAS2 (Swagger 2) file represents a real API that lives at http://todos.stoplight.io.\n\nFor authentication information, click the apikey security scheme in the editor sidebar."

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}
	log.Println("loading spec at:", server.Spec)
	swaggerSpec, err := loads.Spec(string(server.Spec))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("spec loaded with path", swaggerSpec.SpecFilePath())

	api := operations.NewNrcodegenAPI(swaggerSpec)
	server.SetAPI(api)
	defer server.Shutdown()

	server.ConfigureAPI()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
