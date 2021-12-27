package main

import (
	"html/template"
	"net/http"
	"time"
)

func addHeaderAndFooterFiles(tFiles []string) []string {
	//工作日使用普通头尾，双休使用节日专用头尾
	switch time.Now().Weekday() {
	case time.Saturday:
	case time.Sunday:
		tFiles = append(tFiles, "header_week.html", "footer_week.html")
	default:
		tFiles = append(tFiles, "header.html", "footer.html")
	}
	return tFiles
}

func index(rw http.ResponseWriter, r *http.Request) {
	today := time.Now().Format("2006-01-02")
	tFiles := []string{"index.html"}
	tFiles = addHeaderAndFooterFiles(tFiles)
	t := template.Must(template.ParseFiles(tFiles...))
	t.ExecuteTemplate(rw, "html", today)
}

func main() {
	http.HandleFunc("/index", index)
	http.ListenAndServe(":8080", nil)
}
