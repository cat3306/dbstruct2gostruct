package main

import (
	"fmt"
	"github.com/cat3306/dbstruct2gostruct"
)

func main() {
	t2t := dbstruct2gostruct.NewTable2Struct()

	err := t2t.
		SavePath("./t.go").
		Dsn(&dbstruct2gostruct.DsnConf{
			Ip:       "127.0.0.1",
			Port:     3306,
			DataBase: "demo",
			User:     "root",
			Pwd:      "12345678",
		}).
		Table("student_grade").DateToTime(true).TagKey("gorm").EnableJsonTag().PackageName("main").
		Run()
	fmt.Println(err)
}
