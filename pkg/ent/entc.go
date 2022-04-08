//+build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{},entc.Extensions(&RestExtension{}))
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

type RestExtension struct {
	entc.DefaultExtension
}

func (*RestExtension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("rest").ParseFiles("../template/rest.tmpl")),
	}
}