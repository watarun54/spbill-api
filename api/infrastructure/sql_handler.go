package infrastructure

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"

	"github.com/watarun54/serverless-skill-manager/server/interfaces/database"
)

type SqlHandler struct {
	Conn *gorm.DB
	Asso *gorm.Association
}

func NewSqlHandler() database.SqlHandler {
	isLambda := os.Getenv("LAMBDA")
	if isLambda != "TRUE" {
		envPath := fmt.Sprintf("%s.env", os.Getenv("GO_ENV"))
		err := godotenv.Load(os.ExpandEnv(envPath))
		if err != nil {
			panic(err)
		}
	}
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbURL := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	conn, err := gorm.Open("mysql", dbURL)
	if err != nil {
		panic(err)
	}
	log.Println("connected to", dbHost)
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *SqlHandler) Exec(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(sql, values...)
}

func (handler *SqlHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *SqlHandler) Take(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Take(out, where...)
}

func (handler *SqlHandler) Raw(sql string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(sql, values...)
}

func (handler *SqlHandler) Create(value interface{}) *gorm.DB {
	return handler.Conn.Create(value)
}

func (handler *SqlHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Save(value)
}

func (handler *SqlHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *SqlHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}

func (handler *SqlHandler) Preload(column string, conditions ...interface{}) *gorm.DB {
	return handler.Conn.Preload(column, conditions)
}

func (handler *SqlHandler) Set(name string, value interface{}) *gorm.DB {
	return handler.Conn.Set(name, value)
}

func (handler *SqlHandler) Scan(dest interface{}) *gorm.DB {
	return handler.Conn.Scan(dest)
}

func (handler *SqlHandler) Association(column string) *gorm.Association {
	return handler.Conn.Association(column)
}

func (handler *SqlHandler) Replace(values ...interface{}) *gorm.Association {
	return handler.Asso.Replace(values)
}

func (handler *SqlHandler) Debug() *gorm.DB {
	return handler.Conn.Debug()
}

func (handler *SqlHandler) Model(value interface{}) *gorm.DB {
	return handler.Conn.Model(value)
}
