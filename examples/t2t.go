package main

import (
	"fmt"
	"github.com/cat3306/dbstruct2gostruct"
)

func main() {
	t2t := dbstruct2gostruct.NewTable2Struct()

	err := t2t.
		SavePath("./t.go").
		Dsn("root:12345678@tcp(127.0.0.1:3306)/demo?charset=utf8").
		Table("student_grade").DateToTime(true).TagKey("gorm").EnableJsonTag().
		Run()
	fmt.Println(err)
}
