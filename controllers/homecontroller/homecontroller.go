package homecontroller

import (
  "html/template"
  "net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    http.NotFound(w, r)
    return
  }

  temp, _ := template.ParseFiles("views/home/index.html", "views/layout/index.html")
  temp.Execute(w, map[string]string{"title": "Golang CRUD | Home"})
}
