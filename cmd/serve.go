package cmd

import (
	"context"
	"net/http"

	"github.com/lesomnus/otx"
	"github.com/lesomnus/otx/otxhttp"
	"github.com/lesomnus/vend/internal/vend"
	"github.com/lesomnus/xli"
)

func NewCmdServe() *xli.Command {
	return &xli.Command{
		Name:  "serve",
		Brief: "start the vanity import server",

		Handler: xli.OnRun(func(ctx context.Context, cmd *xli.Command, next xli.Next) error {
			c := use_config.Must(ctx)
			s := vend.NewServer(c.Packages)

			x := otx.From(ctx)
			h := otxhttp.NewHandler(x, otxhttp.BoundaryLogger()(s), "vend")

			mux := http.NewServeMux()
			mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
			})
			mux.Handle("/", h)

			return http.ListenAndServe(c.Server.Addr, mux)
		}),
	}
}
