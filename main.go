package main

import (
  "fmt"
  "log"
  "net/http"
  
  "github.com/Rizkyyullah/golang-crud/config"
  "github.com/Rizkyyullah/golang-crud/controllers/categorycontroller"
  "github.com/Rizkyyullah/golang-crud/controllers/homecontroller"
  "github.com/Rizkyyullah/golang-crud/controllers/productcontroller"
)

func main() {
  config.LoadConfig()
  config.ConnectDB()
  
  // Homepage
  http.HandleFunc("/", homecontroller.Index)
  
  // Categoriespage
  http.HandleFunc("/categories", categorycontroller.Index)
  http.HandleFunc("/categories/add", categorycontroller.Add)
  http.HandleFunc("/categories/edit", categorycontroller.Edit)
  http.HandleFunc("/categories/delete", categorycontroller.Delete)
  
  // Productspage
  http.HandleFunc("/products", productcontroller.Index)
  http.HandleFunc("/products/add", productcontroller.Add)
  http.HandleFunc("/products/edit", productcontroller.Edit)
  http.HandleFunc("/products/detail", productcontroller.Detail)
  http.HandleFunc("/products/delete", productcontroller.Delete)
  
  log.Println(fmt.Sprintf("Server berjalan di http://localhost:%d", config.ENV.APP_PORT))
  http.ListenAndServe(fmt.Sprintf(":%d", config.ENV.APP_PORT), nil)
}
