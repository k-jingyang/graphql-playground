package main

import (
	"encoding/csv"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

var (
	queryType   *graphql.Object
	digimonType *graphql.Object
	digimons    []Digimon
)

type Digimon struct {
	Number    int
	Name      string
	Stage     string
	Type      string
	Attribute string
	HP        int
	Attack    int
}

func init() {

	digimonType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Digimon",
		Fields: graphql.Fields{
			"Number": &graphql.Field{
				Type: graphql.Int,
			},
			"Name": &graphql.Field{
				Type: graphql.String,
			},
			"Stage": &graphql.Field{
				Type: graphql.String,
			},
			"Type": &graphql.Field{
				Type: graphql.String,
			},
			"Attribute": &graphql.Field{
				Type: graphql.String,
			},
			"HP": &graphql.Field{
				Type: graphql.String,
			},
			"Attack": &graphql.Field{
				Type: graphql.Int,
			},
		},
	})

	queryType = graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"Digimon": &graphql.Field{
				Type: graphql.NewList(digimonType),
				Args: graphql.FieldConfigArgument{
					"stage": &graphql.ArgumentConfig{
						Description:  "Stage of digimon",
						Type:         graphql.String,
						DefaultValue: "",
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {

					var filtered []Digimon
					stageFilter := p.Args["stage"]
					if stageFilter == "" {
						return digimons, nil
					}
					for i := 0; i < len(digimons); i++ {
						if digimons[i].Stage == stageFilter {
							filtered = append(filtered, digimons[i])
						}
					}
					return filtered, nil
				},
			},
		},
	})

	loadDigimonInfo()
}

func loadDigimonInfo() {
	csvfile, err := os.Open("data/DigiDB_digimonlist.csv")

	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	digimonData := csv.NewReader(csvfile)
	// Removes header
	digimonData.Read()

	for {
		record, err := digimonData.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		number, err := strconv.Atoi(record[0])
		hp, err := strconv.Atoi(record[7])
		atk, err := strconv.Atoi(record[9])
		name := record[1]
		stage := record[2]
		dType := record[3]
		attribute := record[4]

		digimons = append(digimons, Digimon{number, name, stage, dType, attribute, hp, atk})
	}
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query: queryType,
	})

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/", h)

	log.Print("Starting listener at port ", port)
	http.ListenAndServe(":"+port, nil)

}
