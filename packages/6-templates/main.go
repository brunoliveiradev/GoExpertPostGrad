package main

import (
	"html/template"
	"net/http"
	"os"
	"strings"
)

type Course struct {
	Name  string
	Hours int
}

type Courses []Course

func main() {
	ExecuteFirstTemplateExample()
	ExecuteSecondTemplateExample()
	ExecuteThirdTemplateExample()
	ExecuteFourthTemplateExample()
	//ExecuteFifthTemplateExample()
}

// ExecuteFirstTemplateExample demonstrates the basic usage of text/template package.
// It creates a new template, parses a string into it, and then executes the template.
func ExecuteFirstTemplateExample() {
	course := Course{Name: "Go", Hours: 20}
	templ := template.New("CourseTemplate")
	templ, _ = templ.Parse("Course: {{.Name}} ({{.Hours}} Hours)\n")

	err := templ.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}

// ExecuteSecondTemplateExample demonstrates the usage of template.Must function.
// It creates a new template, parses a string into it, and then executes the template,
// writing the output to os.Stdout. If there is an error during the parsing, the program will panic.
func ExecuteSecondTemplateExample() {
	course := Course{Name: "Golang", Hours: 220}
	tmp := template.Must(template.New("CourseTemplate").Parse("Course: {{.Name}} ({{.Hours}} Hours)\n"))

	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}

// ExecuteThirdTemplateExample demonstrates the usage of template.Must function with ParseFiles.
// It creates a new template, parses a file into it, and then executes the template,
// writing the output to os.Stdout. If there is an error during the parsing, the program will panic.
func ExecuteThirdTemplateExample() {
	tmp := template.Must(template.New("template.html").ParseFiles("packages/6-templates/template.html"))

	err := tmp.Execute(os.Stdout, getCourses())
	if err != nil {
		panic(err)
	}
}

// ExecuteFourthTemplateExample starts a simple HTTP server that responds to requests with a rendered template.
// It creates a new template, parses a file into it, and then executes the template,
// handleRequest is the handler function for HTTP requests. It renders a template and writes it to the response.
func ExecuteFourthTemplateExample() {
	http.HandleFunc("/", handleRequest)
	http.ListenAndServe(":8080", nil)
}

func handleRequest(w http.ResponseWriter, r *http.Request) {
	tmp := template.Must(template.ParseFiles("packages/6-templates/template.html"))
	err := tmp.Execute(w, getCourses())
	if err != nil {
		panic(err)
	}
}

func getCourses() Courses {
	return Courses{
		Course{Name: "Golang", Hours: 220},
		Course{Name: "Python", Hours: 40},
		Course{Name: "Java", Hours: 180},
	}
}

func ExecuteFifthTemplateExample() {
	tmpt := template.New("template.html")

	tmpt.Funcs(template.FuncMap{"ToUpper": ToUpper})
	// Após ter mapeado a cima, tem que ter dinâmico no HTML <td>{{ .Name | ToUpper}}</td>

	tmpt = template.Must(tmpt.ParseFiles("packages/6-templates/template.html"))

	err := tmpt.Execute(os.Stdout, getCourses())
	if err != nil {
		panic(err)
	}
}

func ToUpper(s string) string {
	return strings.ToUpper(s)
}
