package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"

	"github.com/sky0621/familiagildo/tool/sqlparser"
)

//go:embed testdata/*
var testdata embed.FS

func main() {
	execMain()
}

func execMain() {
	fsys, err := fs.Sub(testdata, "testdata")
	if err != nil {
		panic(err)
	}
	tpl := template.Must(template.ParseFS(fsys))
	fmt.Println(tpl)
}

func execTest() {
	for _, sql := range sqls {
		res, err := sqlparser.NewSQLParser().Parse(sql[0], sql[1])
		if err != nil {
			panic(err)
		}

		fmt.Println(res.SQLName)
		for _, x := range res.TableNameWithCRUDSlice {
			fmt.Println(x.TableName)
			fmt.Println(x.CRUD.ToShortName())
		}
	}
}
