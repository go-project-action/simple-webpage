package http

import (
	"log"
	"net/http"
	"text/template"
)

func isOdd(num int) bool {
	return (num % 2) == 0
}

// Home will handle / page.
func (a *AppServer) Home() func(http.ResponseWriter, *http.Request) {
	type Profile struct {
		Name    string
		Work    string
		Amateur string
	}

	var items = []Profile{
		{
			Name:    "Heather",
			Work:    "PHP",
			Amateur: "Go",
		},
		{
			Name:    "Albert",
			Work:    "Java",
			Amateur: "PHP",
		},
		{
			Name:    "Charlotte",
			Work:    "C++",
			Amateur: "Ruby",
		},
		{
			Name:    "Helen",
			Work:    "C#",
			Amateur: "JavaScript",
		},
		{
			Name:    "Pete",
			Work:    "Node.js",
			Amateur: "JavaScript",
		}}

	templates := []string{
		"web/template/layout/base.html",
		"web/template/layout/footer.html",
		"web/template/layout/header.html",
		"web/template/index.html",
	}

	t, err := template.New("index.html").
		Funcs(template.FuncMap{"isOdd": isOdd}).
		ParseFiles(templates...)
	if err != nil {
		log.Println("parse template file error:", err.Error())
		panic(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		t.ExecuteTemplate(w, "base", items)
		if err != nil {
			log.Println("template execute error:", err.Error())
			return
		}
	}
}
