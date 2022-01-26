package database

import "github.com/jinzhu/gorm"

type SqlHandler interface {
	Exec(string, ...interface{}) *gorm.DB
	Find(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
	Take(interface{}, ...interface{}) *gorm.DB
	Raw(string, ...interface{}) *gorm.DB
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Delete(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	Preload(string, ...interface{}) *gorm.DB
	Set(string, interface{}) *gorm.DB
	Scan(interface{}) *gorm.DB
	Association(string) *gorm.Association
	Replace(...interface{}) *gorm.Association
	Debug() *gorm.DB
}
