package docs

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

//go:embed swagger-ui
var SwaggerUI embed.FS

// RegisterSwaggerAPI provides a common function which registers swagger route with API Server
func RegisterSwaggerAPI(_ client.Context, rtr *mux.Router) error {

	root, err := fs.Sub(SwaggerUI, "swagger-ui")
	if err != nil {
		return err
	}

	staticServer := http.FileServer(http.FS(root))
	rtr.PathPrefix("/sixchain/swagger/").Handler(http.StripPrefix("/sixchain/swagger/", staticServer))

	return nil
}
