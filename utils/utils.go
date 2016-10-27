package utils

import (
	"fmt"
	"html/template"
)

// Log msg ...
func Log(msg string, d ...interface{}) {
	fmt.Printf(fmt.Sprintf("\n%s\n", msg), d...)
}

// LoadTemplate load template safe and easy control
func LoadTemplate(path ...string) *template.Template {
	t := template.Must(template.ParseFiles(path...))
	return t
}
