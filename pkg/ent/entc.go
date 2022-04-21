//+build ignore

package main

import (
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"fmt"
	"log"
	"text/template"
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
		gen.MustParse(gen.NewTemplate("rest").Funcs(template.FuncMap{"echo": echo}).ParseFiles("../template/rest.tmpl")),
	}
}

func echo(vs ...interface{}) string  {
	for _,v := range vs {
		fmt.Println(v)
	}
	return ""
}