package graphql

import (
	"io/ioutil"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/graph-gophers/graphql-go"
	"github.com/graph-gophers/graphql-go/relay"
	"google.golang.org/grpc/resolver"
)

type Schema struct {
	schemaString string
	Schema       *graphql.Schema
}

func NewSchema() Schema {
	return Schema{}
}

func (sch *Schema) Config(resolver resolver.Resolver) {
	sch.readSchemaFiles()
	sch.parseSchema(resolver)
}

func (sch *Schema) readSchemaFiles() {
	schemaPath := "./graphql/"
	typesPath := "./graphql/types/"

	schema, err := ioutil.ReadFile(schemaPath + "schema.graphql")
	if err != nil {
		panic(err)
	}
	schema = append(schema, []byte("\n")...)

	files, err := ioutil.ReadDir(typesPath)
	if err != nil {
		log.Fatalln(err)
	}

	for _, file := range files {
		temp, err := ioutil.ReadFile(typesPath + file.Name())
		if err != nil {
			log.Fatalln(err)
		}
		schema = append(schema, temp...)
		schema = append(schema, []byte("\n")...)
	}
	sch.schemaString = string(schema)
}

func (sch *Schema) parseSchema(resolver resolver.Resolver) {
	sch.Schema = graphql.MustParseSchema(sch.schemaString, &resolver)
}

func (sch *Schema) SchemaHandle() *relay.Handler {
	return &relay.Handler{Schema: sch.Schema}
}

func (sch *Schema) GinHandler() gin.HandlerFunc {
	return gin.WrapH(&relay.Handler{Schema: sch.Schema})
}
