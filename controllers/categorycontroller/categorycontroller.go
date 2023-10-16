package categorycontroller

import (
  "html/template"
  "net/http"
  "strconv"
  "time"
  
  "github.com/Rizkyyullah/golang-crud/entities"
  "github.com/Rizkyyullah/golang-crud/models/categorymodel"
)

type M map[string]any

func Index(w http.ResponseWriter, r *http.Request) {
  categories := categorymodel.GetAll()
  data := M{
    "title": "Golang CRUD | Categories",
    "categories": categories,
  }
  
  temp, _ := template.ParseFiles("views/category/index.html", "views/layout/index.html")
  temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case http.MethodGet: {
      data := M{
        "title": "Golang CRUD | Categories - Add",
      }
      
      temp, _ := template.ParseFiles("views/category/add.html", "views/layout/index.html")
      temp.Execute(w, data)
    }
    case http.MethodPost: {
      category := entities.Category{
        Name: r.FormValue("name"),
      }
      
      if success := categorymodel.Create(category); !success {
        data := M{
          "title": "Golang CRUD | Categories - Add",
          "error": "Gagal membuat category baru",
        }
        
        temp, _ := template.ParseFiles("views/category/add.html", "views/layout/index.html")
        temp.Execute(w, data)
        return
      }
      
      http.Redirect(w, r, "/categories", http.StatusMovedPermanently)
    }
  }
}

func Edit(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case http.MethodGet: {
      id := r.URL.Query().Get("id")
      category := categorymodel.Detail(id)
      
      data := M{
        "title": "Golang CRUD | Categories - Edit",
        "category": category,
      }
      
      temp, _ := template.ParseFiles("views/category/edit.html", "views/layout/index.html")
      temp.Execute(w, data)
    }
    case http.MethodPost: {
      idString := r.FormValue("id")
      id, _ := strconv.Atoi(idString)
      
      category := entities.Category{
        ID: uint(id),
        Name: r.FormValue("name"),
        UpdatedAt: time.Now(),
      }
      
      if success := categorymodel.Edit(category); !success {
        data := M{
          "title": "Golang CRUD | Categories - Edit",
          "error": "Gagal mengubah category",
        }
        
        temp, _ := template.ParseFiles("views/category/add.html", "views/layout/index.html")
        temp.Execute(w, data)
        return
      }
      
      http.Redirect(w, r, "/categories", http.StatusMovedPermanently)
    }
  }
}

func Delete(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query().Get("id")
  categorymodel.Delete(id)
  http.Redirect(w, r, "/categories", http.StatusMovedPermanently)
}
