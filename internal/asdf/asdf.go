package asdf

import (
	"encoding/json"
	"net/http"
	"net/url"

	config "github.com/JesKetchupson/asdf/configs"
	gen "github.com/JesKetchupson/asdf/internal/pkg/generation"
	"github.com/JesKetchupson/asdf/plugins"
	"github.com/JesKetchupson/asdf/storage"
	"golang.org/x/exp/errors/fmt"
)

type AsdfServer struct {
	db storage.DB
}

func Run(conf config.AsdfConfig) error {
	http.HandleFunc("/", ParceNewUri)

	if err := http.ListenAndServe(conf.ServerPort, nil); err != nil {
		return err
	}

	return nil
}

type requestStruct struct {
	URI string `json:"uri"`
}

func ParceNewUri(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var t requestStruct
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	u, err := url.Parse(t.URI)
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	done := make(chan string)

	fmt.Fprintf(w, "Start parcing")

	go func(u *url.URL) {
		//TODO: 1) add protocol discovery
		params := gen.Parametres{
			TypeName: u.Host,
			URL:      u,
		}
		params.GetNewDataType()

		if err := params.GenerateCode(); err != nil {
			fmt.Fprintf(w, "err during generating code %s", err)
			close(done)
			return
		}

		path, err := plugins.Build(params.TypeName)

		if err != nil {
			fmt.Fprintf(w, "err during building plugin %s", err)
			close(done)
			return
		}

		obj, err := plugins.LoadPlugin(path)

		obj.Save(nil)

		done <- params.TypeName
	}(u)

	if tableName, ok := <-done; ok {
		fmt.Fprintf(w, "new table named %s crated", tableName)
	}

}
