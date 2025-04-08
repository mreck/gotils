package httptils

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type JSONResponse[T any] struct {
	Data  T      `json:"data"`
	Error string `json:"error"`
}

type JSONFunc func(context.Context, *http.Request) (int, any, error)

type JSONHandler struct {
	ctx     context.Context
	slogger *slog.Logger
}

func NewJSONHandler(ctx context.Context) *JSONHandler {
	return &JSONHandler{
		ctx:     ctx,
		slogger: slog.Default().With(slog.String("httptils", "json")),
	}
}

func (h *JSONHandler) SetSlogger(logger *slog.Logger) { h.slogger = logger }

func (h *JSONHandler) H(f JSONFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp JSONResponse[any]

		code, data, err := f(h.ctx, r)
		if err != nil {
			resp = JSONResponse[any]{Error: err.Error()}
		} else {
			resp = JSONResponse[any]{Data: data}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			h.slogger.Error("encoding data failed", "error", err)

			w.WriteHeader(http.StatusInternalServerError)
			resp = JSONResponse[any]{Error: http.StatusText(http.StatusInternalServerError)}
			_ = json.NewEncoder(w).Encode(resp)
		}
	}
}
