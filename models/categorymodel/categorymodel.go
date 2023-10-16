package categorymodel

import (
  "github.com/Rizkyyullah/golang-crud/config"
  "github.com/Rizkyyullah/golang-crud/entities"
)

func GetAll() []entities.Category {
  query := "SELECT id, name, created_at, updated_at FROM categories;"
  rows, err := config.DB.Query(query)
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  
  categories := []entities.Category{}
  for rows.Next() {
    category := entities.Category{}
    if err := rows.Scan(&category.ID, &category.Name, &category.CreatedAt, &category.UpdatedAt); err != nil {
      panic(err)
    }
    
    categories = append(categories, category)
  }
  
  if len(categories) > 0 {
    return categories
  }
  
  return nil
}

func Create(c entities.Category) bool {
  query := "INSERT INTO categories(name) VALUES($1);"
  result, err := config.DB.Exec(query, c.Name)
  if err != nil {
    return false
  }
  
  rowsAffected, _ := result.RowsAffected()
  return rowsAffected > 0
}

func Detail(id string) entities.Category {
  var category entities.Category
  query := "SELECT id, name FROM categories WHERE id = $1 ORDER BY id;"
  if err := config.DB.QueryRow(query, id).Scan(&category.ID, &category.Name); err != nil {
    panic(err)
  }
  
  return category
}

func Edit(c entities.Category) bool {
  query := "UPDATE categories SET name = $2, updated_at = $3 WHERE id = $1;"
  result, err := config.DB.Exec(query, c.ID, c.Name, c.UpdatedAt)
  if err != nil {
    return false
  }
  
  rowsAffected, _ := result.RowsAffected()
  return rowsAffected > 0
}

func Delete(id string) {
  query := "DELETE FROM categories WHERE id = $1;"
  if _, err := config.DB.Exec(query, id); err != nil {
    panic(err)
  }
}
