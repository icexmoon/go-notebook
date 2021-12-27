package main

import (
	"html/template"
	"net/http"
)

func index(rw http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")
	fm := template.FuncMap{"score_pass": func(score float64) bool { return score > 60 }}
	t.Funcs(fm)
	t.ParseFiles("index.html")
	scores := make(map[string][]float64)
	scores["Li lei"] = []float64{55, 60, 80}
	scores["Han Meimei"] = []float64{90, 60, 33}
	scores["Jack Chen"] = []float64{80, 100, 95}
	t.Execute(rw, scores)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", http.DefaultServeMux)
}
