package plugins

import (
	"os/exec"
	"plugin"

	"github.com/JesKetchupson/asdf/storage"
	"golang.org/x/exp/errors/fmt"
)

const pathToSo = "internal/pkg/code_gen/shared_objects/"

type SaveablePluginObject interface {
	Save(storage.DB)
}

//TODO: build versioning
func Build(filename string) (PathToPlugin string, err error) {
	cmd := exec.Command("go", "build", "-o", pathToSo, "-buildmode=plugin", filename+".go")
	errOut, err := cmd.CombinedOutput() //errOut on _
	if err != nil {
		fmt.Println(string(errOut))
		return PathToPlugin, err
	}

	PathToPlugin = pathToSo + filename + ".so"

	return PathToPlugin, nil
}
func LoadPlugin(PathToPlugin string) (SaveablePluginObject, error) {
	p, err := plugin.Open(PathToPlugin)
	if err != nil {
		return nil, err
	}
	ps, err := p.Lookup("NewType")
	if err != nil {
		return nil, err
	}
	return ps.(SaveablePluginObject), err
}
