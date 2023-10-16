package productmodel

import (
  "github.com/Rizkyyullah/golang-crud/config"
  "github.com/Rizkyyullah/golang-crud/entities"
)

func GetAll() []entities.Product {
  query := `
    SELECT 
      p.id,
      p.name,
      c.name,
      p.stock,
      p.price,
      p.description,
      p.created_at,
      p.updated_at
    FROM 
      products p
    LEFT JOIN
      categories c
    ON
      p.category_id = c.id
    ORDER BY id;
  `
  rows, err := config.DB.Query(query)
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  
  var products []entities.Product
  for rows.Next() {
    product := entities.Product{}
    if err := rows.Scan(&product.ID, &product.Name, &product.Category.Name, &product.Stock, &product.Price, &product.Description, &product.CreatedAt, &product.UpdatedAt); err != nil {
      panic(err)
    }
    
    products = append(products, product)
  }
  
  if len(products) > 0 {
    return products
  }
  
  return nil
}

func GetCategories() []entities.Category {
  query := `
    SELECT 
      id,
      name
    FROM 
      categories
    ORDER BY
      id;
  `
  rows, err := config.DB.Query(query)
  if err != nil {
    panic(err)
  }
  defer rows.Close()
  
  var categories []entities.Category
  for rows.Next() {
    category := entities.Category{}
    if err := rows.Scan(&category.ID, &category.Name); err != nil {
      panic(err)
    }
    
    categories = append(categories, category)
  }
  
  if len(categories) > 0 {
    return categories
  }
  
  return nil
}

func Create(p entities.Product) bool {
  query := "INSERT INTO products(name, category_id, stock, price, description) VALUES($1, $2, $3, $4, $5);"
  result, err := config.DB.Exec(query,
    p.Name,
    p.Category.ID,
    p.Stock,
    p.Price,
    p.Description,
  )
  if err != nil {
    return false
  }
  
  rowsAffected, _ := result.RowsAffected()
  return rowsAffected > 0
}

func Detail(id string) entities.Product {
  product := entities.Product{}
  query := `
    SELECT
      p.id,
      p.name,
      c.id,
      c.name,
      p.stock,
      p.price,
      p.description,
      p.updated_at
    FROM 
      products p
    LEFT JOIN
      categories c
    ON 
      p.category_id = c.id
    WHERE
      p.id = $1;`
  if err := config.DB.QueryRow(query, id).Scan(
    &product.ID,
    &product.Name,
    &product.Category.ID,
    &product.Category.Name,
    &product.Stock,
    &product.Price,
    &product.Description,
    &product.UpdatedAt,
  ); err != nil {
    panic(err)
  }
  
  return product
}

func Edit(p entities.Product, id string) bool {
  query := `
    UPDATE
      products
    SET
      name = $2,
      category_id = $3,
      stock = $4,
      price = $5,
      description = $6,
      updated_at = $7
    WHERE
      id = $1;
  `
  result, err := config.DB.Exec(query,
    id,
    p.Name,
    p.Category.ID,
    p.Stock,
    p.Price,
    p.Description,
    p.UpdatedAt,
  )
  if err != nil {
    return false
  }
  
  rowsAffected, _ := result.RowsAffected()
  return rowsAffected > 0
}

func Delete(id string) {
  query := "DELETE FROM products WHERE id = $1;"
  if _, err := config.DB.Exec(query, id); err != nil {
    panic(err)
  }
}
