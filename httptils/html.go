package httptils

import (
	"context"
	"html/template"
	"log/slog"
	"net/http"

	"github.com/mreck/gotils"
)

type D map[string]any

type HTMLFunc func(context.Context, *http.Request) (int, string, D, error)

type HTMLHandler struct {
	ctx         context.Context
	tmpl        *template.Template
	errTmplName string
	slogger     *slog.Logger
}

func NewHTMLHandler(ctx context.Context, tmpl *template.Template) *HTMLHandler {
	return &HTMLHandler{
		ctx:         ctx,
		tmpl:        tmpl,
		errTmplName: "error.html",
		slogger:     slog.Default().With(slog.String("httptils", "html")),
	}
}

func (h *HTMLHandler) SetSLogger(logger *slog.Logger)   { h.slogger = logger }
func (h *HTMLHandler) SetErrorTemplateName(name string) { h.errTmplName = name }

func (h *HTMLHandler) H(f HTMLFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code, name, data, err := f(h.ctx, r)
		if err != nil {
			if data == nil {
				data = D{}
			}
			gotils.ExtendMap(data, D{"Error": err})
			name = h.errTmplName
		}

		w.WriteHeader(code)

		err = h.tmpl.ExecuteTemplate(w, name, data)
		if err != nil {
			h.slogger.Error("executing template failed", "error", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
