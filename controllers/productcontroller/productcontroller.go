package productcontroller

import (
  "html/template"
  "net/http"
  "strconv"
  "time"
  
  "github.com/Rizkyyullah/golang-crud/models/productmodel"
  "github.com/Rizkyyullah/golang-crud/entities"
)

type M map[string]any

func Index(w http.ResponseWriter, r *http.Request) {
  products := productmodel.GetAll()
  data := M{
    "title": "Golang CRUD | Products ",
    "products": products,
  }
  
  temp, _ := template.ParseFiles("views/product/index.html", "views/layout/index.html")
  temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case http.MethodGet: {
      categories := productmodel.GetCategories()
      data := M{
        "title": "Golang CRUD | Products - Add",
        "categories": categories,
      }
      
      temp, _ := template.ParseFiles("views/product/add.html", "views/layout/index.html")
      temp.Execute(w, data)
    }
    case http.MethodPost: {
      categoryIdStr := r.FormValue("category")
      stockStr := r.FormValue("stock")
      priceStr := r.FormValue("price")
      
      categoryId, _ := strconv.Atoi(categoryIdStr)
      stock, _ := strconv.Atoi(stockStr)
      price, _ := strconv.Atoi(priceStr)
      
      product := entities.Product{
        Name: r.FormValue("name"),
        Category: entities.Category{
          ID: uint(categoryId),
        },
        Stock: uint(stock),
        Price: uint(price),
        Description: r.FormValue("description"),
      }
      
      if success := productmodel.Create(product); !success {
        data := M{
          "title": "Golang CRUD | Products - Add",
          "error": "Proses membuat product gagal!",
        }
        
        temp, _ := template.ParseFiles("views/product/add.html", "views/layout/index.html")
        temp.Execute(w, data)
      }
      
      http.Redirect(w, r, "/products", http.StatusMovedPermanently)
    }
  }
}

func Detail(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query().Get("id")
  product := productmodel.Detail(id)
  data := M{
    "title": "Golang CRUD | Products - Detail",
    "product": product,
  }
  
  temp, _ := template.ParseFiles("views/product/detail.html", "views/layout/index.html")
  temp.Execute(w, data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
  switch r.Method {
    case http.MethodGet: {
      id := r.URL.Query().Get("id")
      product := productmodel.Detail(id)
      categories := productmodel.GetCategories()
      data := M{
        "title": "Golang CRUD | Products - Edit",
        "product": product,
        "categories": categories,
      }
      
      temp, _ := template.ParseFiles("views/product/edit.html", "views/layout/index.html")
      temp.Execute(w, data)
    }
    case http.MethodPost: {
      id := r.FormValue("id")

      categoryIdStr := r.FormValue("category")
      stockStr := r.FormValue("stock")
      priceStr := r.FormValue("price")
      
      categoryId, _ := strconv.Atoi(categoryIdStr)
      stock, _ := strconv.Atoi(stockStr)
      price, _ := strconv.Atoi(priceStr)
      
      product := entities.Product{
        Name: r.FormValue("name"),
        Category: entities.Category{
          ID: uint(categoryId),
        },
        Stock: uint(stock),
        Price: uint(price),
        Description: r.FormValue("description"),
        UpdatedAt: time.Now(),
      }
      
      if success := productmodel.Edit(product, id); !success {
        data := M{
          "title": "Golang CRUD | Products - Edit",
          "error": "Gagal memperbarui product",
        }
        
        temp, _ := template.ParseFiles("views/product/edit.html", "views/layout/index.html")
        temp.Execute(w, data)
        return
      }
      
      http.Redirect(w, r, "/products", http.StatusMovedPermanently)
    }
  }
}

func Delete(w http.ResponseWriter, r *http.Request) {
  id := r.URL.Query().Get("id")
  productmodel.Delete(id)
  http.Redirect(w, r, "/products", http.StatusMovedPermanently)
}
