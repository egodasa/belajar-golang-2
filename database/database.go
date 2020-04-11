package database
import (
  _ "github.com/go-sql-driver/mysql"
  "github.com/jmoiron/sqlx"
  "github.com/joho/godotenv"
  "os"
)

func Connect() *sqlx.DB {
  err := godotenv.Load()
  if err != nil {
    panic("Error loading .env file")
  }

  DB_CONNECTION := os.Getenv("DB_CONNECTION")
  DB_STRING := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") + "@(" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + ")" + "/" + os.Getenv("DB_DATABASE")
  
  // konek ke database
  db, err := sqlx.Connect(DB_CONNECTION, DB_STRING)
  // cek error konek ke database
  if err != nil {
    panic(err);
  }
  return db
}
