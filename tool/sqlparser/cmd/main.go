package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/sky0621/familiagildo/tool/sqlparser"
)

func main() {
	execMain()
}

func execMain() {
	if err := filepath.WalkDir(filepath.Join("cmd", "testdata"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		fmt.Println(path)
		fl, err := os.Open(path)
		if err != nil {
			panic(err)
		}

		sqlName := ""
		sql := strings.Builder{}

		sc := bufio.NewScanner(fl)
		for sc.Scan() {
			line := sc.Text()
			if isBlankLine(line) {
				fmt.Println("is blank line...")
				continue
			}
			if isSQLNameLine(line) {
				sqlName = getSQLName(line)
				fmt.Println(sqlName)
				continue
			}
			sql.WriteString(line + " ")

		}
		return nil
	}); err != nil {
		panic(err)
	}
}

func isBlankLine(line string) bool {
	return len(strings.Trim(line, " ")) == 0
}

func isSQLNameLine(line string) bool {
	fmt.Println(line)
	return strings.HasPrefix(strings.Trim(line, " "), "--name")
}

func getSQLName(line string) string {
	// 形式　-- name: CreateGuestToken :one
	tLine := strings.Trim(line, " ")
	tpLine := strings.TrimPrefix(tLine, "--")
	tokens := strings.Split(tpLine, ":")
	if len(tokens) != 3 {
		return ""
	}
	return tokens[1]
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
