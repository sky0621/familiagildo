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
	var sqlParseResults []*sqlparser.SQLParseResult
	if err := filepath.WalkDir(filepath.Join("cmd", "testdata"), func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() {
			return nil
		}

		fl, err := os.Open(path)
		if err != nil {
			panic(err)
		}
		defer func() {
			if fl != nil {
				if err := fl.Close(); err != nil {
					fmt.Println(err)
				}
			}
		}()

		stat, err := fl.Stat()
		if err != nil {
			panic(err)
		}

		sqlFileName := stat.Name()

		sqlName := ""
		sql := strings.Builder{}

		sc := bufio.NewScanner(fl)
		for sc.Scan() {
			line := sc.Text()

			if isBlankLine(line) {
				continue
			}

			if isSQLNameLine(line) {
				sqlName = getSQLName(line)
				continue
			}

			sql.WriteString(line + " ")

			if isEndSQL(line) {
				res, err := sqlparser.NewSQLParser().Parse(sqlName, sqlFileName, sql.String())
				if err != nil {
					panic(err)
				}
				sqlParseResults = append(sqlParseResults, res)

				sqlName = ""
				sql.Reset()
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}

	for _, pr := range sqlParseResults {
		fmt.Printf("[%s (%s)]\n", pr.SQLName, pr.SQLFileName)
		for _, x := range pr.TableNameWithCRUDSlice {
			fmt.Printf("%s : %s\n", x.TableName, x.CRUD.ToShortName())
		}
		fmt.Println("===============")

	}
}

func isBlankLine(line string) bool {
	return len(strings.Trim(line, " ")) == 0
}

func isSQLNameLine(line string) bool {
	return strings.HasPrefix(strings.Trim(line, " "), "-- name")
}

func isEndSQL(line string) bool {
	return strings.HasSuffix(strings.Trim(line, " "), ";")
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
