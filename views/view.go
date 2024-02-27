package views

import (
	"os"
	"text/template"
	"time"
)

// Tmpl is a template.
var Tmpl *template.Template

func date(t *time.Time) string {
	return t.Local().Format("January 2, 2006 15:04:05")
}

func shortDate(t *time.Time) string {
	return t.Local().Format("January 2, 2006")
}

func truncate(s string) string {
	result := s

	if len(s) > 160 {
		result = s[0:160] + "..."
	}

	return result
}

var functions = template.FuncMap{
	"date":      date,
	"shortDate": shortDate,
	"truncate":  truncate,
}

func getDir() string {
	dir := ""

	return dir
}

func init() {
	dir, _ := os.Getwd()
	Tmpl = template.Must(template.New(dir + "./views/*.gohtml").Funcs(functions).ParseGlob(dir + "./views/*.gohtml"))
}

// func init() {
// 	var err error
// 	dir := getDir() + "./views/*.gohtml"

// 	Tmpl, err = template.New("views").Funcs(functions).ParseGlob(dir)
// 	if err != nil {
// 		log.Fatalf("Error parsing templates: %v", err)
// 	}
// }
