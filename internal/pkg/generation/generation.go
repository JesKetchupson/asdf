package generation

import (
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"
	"text/template"

	jtg "github.com/JesKetchupson/asdf/third_party/json_to_go_struct"
)

type Parametres struct {
	TypeName string
	NewType  string
	URL      *url.URL
}

const pathToTmpl = "internal/pkg/code_gen/templates/"

func loadTempl(name string) *template.Template {
	b, err := ioutil.ReadFile(pathToTmpl + "struct.tmpl")

	if err != nil {
		panic(err)
	}

	t := template.Must(template.New(name).Parse(string(b)))
	return t
}

func (p Parametres) GenerateCode() error {
	fmt.Printf("Generating %s\n", p.TypeName)

	file, err := os.Create(p.TypeName + ".go")
	tmpl := loadTempl(p.TypeName)
	p.GetNewDataType()

	err = tmpl.Execute(file, p)

	if err != nil {
		return err
	}
	return err
}
func (p *Parametres) GetNewDataType() {
	p.NewType = jtg.Parce(p.URL)
	strings.Replace(p.NewType, "AutoGenerated", p.TypeName, 1)
	strings.Replace(p.NewType, "[]struct", "struct", 1)
}
