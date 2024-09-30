package main

import (
	"fmt"
	"reflect"

	"github.com/go-metaverse/zeri/tag"
)

type Entity struct {
	Name  string `gorm:"primaryKey;column:user_id"`
	Age   int    `gorm:"column:age"`
	Email string `gorm:"column:email"`
}

func main() {
	objT := reflect.TypeOf(Entity{})
	for i := 0; i < objT.NumField(); i++ {
		tagGorm := tag.ParseTag(objT.Field(i).Tag.Get("gorm"), ";")
		fmt.Println(tagGorm)
	}
}
