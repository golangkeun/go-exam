// This is the name of our package
// Everything with this package name can see everything
// else inside the same package, regardless of the file they are in
package main

// These are the libraries we are going to use
import (
	"fmt"
	"log"
	"os"
	"text/template"
)

// Declare type pointer to a template
var tmpl *template.Template

type StrukturData struct {
	id   int
	Nama string
}

// Using the init function to make sure the template is only parsed once in the program
func init() {
	// template.Must takes the reponse of template.ParseFiles and does error checking
	tmpl = template.Must(template.ParseGlob("templates/*"))
}

// Parse all the files in a certain directory
func main() {
	fmt.Println("Hello, World!\n")
	// Execute myName into the template and print to Stdout
	data := []StrukturData{
		StrukturData{
			id:   1,
			Nama: "Udin",
		},
		StrukturData{
			id:   2,
			Nama: "Sasmita",
		},
		StrukturData{
			id:   3,
			Nama: "Ujang",
		},
	}
	err := tmpl.ExecuteTemplate(os.Stdout, "index.gohtml", &data)
	if err != nil {
		log.Fatalln(err)
	}
} // end main
