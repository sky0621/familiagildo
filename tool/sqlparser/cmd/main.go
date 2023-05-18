package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/xuri/excelize/v2"

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

	sortedUniqueTableNames := collectTableNames(sqlParseResults)
	fmt.Println(sortedUniqueTableNames)

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	sheetName := "CRUD"
	index, err := f.NewSheet(sheetName)
	if err != nil {
		panic(err)
	}

	if err := f.SetCellStr(sheetName, "A2", "No"); err != nil {
		panic(err)
	}
	if err := f.SetCellStr(sheetName, "B2", "SQL関数名"); err != nil {
		panic(err)
	}
	if err := f.SetCellStr(sheetName, "C2", "SQLファイル名"); err != nil {
		panic(err)
	}
	for i, tableName := range sortedUniqueTableNames {
		if err := f.SetCellStr(sheetName, fmt.Sprintf("%s2", tableColSet[i]), tableName); err != nil {
			panic(err)
		}
	}

	for i, x := range sqlParseResults {
		if err := f.SetCellInt(sheetName, fmt.Sprintf("A%d", i+3), i+1); err != nil {
			panic(err)
		}
		if err := f.SetCellStr(sheetName, fmt.Sprintf("B%d", i+3), x.SQLName.ToString()); err != nil {
			panic(err)
		}
		if err := f.SetCellStr(sheetName, fmt.Sprintf("C%d", i+3), x.SQLFileName.ToString()); err != nil {
			panic(err)
		}
		// FIXME:
	}

	f.SetActiveSheet(index)

	if err := f.SaveAs("CRUD.xlsx"); err != nil {
		fmt.Println(err)
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

func collectTableNames(sqlParseResults []*sqlparser.SQLParseResult) []string {
	m := map[string]struct{}{}
	for _, x := range sqlParseResults {
		for _, y := range x.TableNameWithCRUDSlice {
			m[y.TableName.ToString()] = struct{}{}
		}
	}
	var r []string
	for k, _ := range m {
		r = append(r, k)
	}
	sort.Strings(r)
	return r
}

var tableColSet = []string{"D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U"}
