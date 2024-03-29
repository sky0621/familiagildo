package main

import (
	_ "embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/spf13/viper"
)

const (
	idPackage  = "github.com/sky0621/familiagildo/adapter/controller/model."
	outputFile = "../../../src/backend/adapter/controller/model/scalar_%s.go"
)

//go:embed template/scalarIDGo.tmpl
var scalarIDGo string

func init() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("gqlgen")
	viper.AddConfigPath("../../../src/backend")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
}

func main() {
	t := template.Must(template.New("").Parse(scalarIDGo))

	for _, id := range targetIDPackages() {
		if err := exec(id, t); err != nil {
			log.Printf("failed to create for id:%s ... %+v", id, err)
			continue
		}
	}
}

func targetIDPackages() []string {
	idPackages := viper.Sub("models").Sub("id").GetStringSlice("model")

	var results []string
	for _, idPkg := range idPackages {
		if strings.Contains(idPkg, idPackage) {
			results = append(results, strings.Trim(idPkg, idPackage))
		}
	}

	return results
}

func exec(id string, t *template.Template) error {
	f, err := os.OpenFile(fmt.Sprintf(outputFile, id), os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer func() {
		if f != nil {
			if err := f.Close(); err != nil {
				log.Printf("failed to close file: %+v", err)
			}
		}
	}()

	if err := t.Execute(f, id); err != nil {
		return err
	}

	return nil
}
