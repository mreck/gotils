package httptils_test

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mreck/gotils/httptils"

	"github.com/stretchr/testify/assert"
)

func TestJSONHandler(t *testing.T) {
	t.Run("handle success", func(t *testing.T) {
		h := http.NewServeMux()
		j := httptils.NewJSONHandler(context.TODO())

		h.HandleFunc("/", j.H(func(ctx context.Context, r *http.Request) (int, any, error) {
			return http.StatusOK, "ok", nil
		}))

		s := httptest.NewServer(h)
		defer s.Close()

		r, err := http.Get(s.URL)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		defer r.Body.Close()

		var rd httptils.JSONResponse[string]
		err = json.NewDecoder(r.Body).Decode(&rd)
		assert.Nil(t, err)
		assert.Equal(t, httptils.JSONResponse[string]{Data: "ok"}, rd)
	})

	t.Run("handle failure", func(t *testing.T) {
		h := http.NewServeMux()
		j := httptils.NewJSONHandler(context.TODO())
		e := errors.New("test error")

		h.HandleFunc("/", j.H(func(ctx context.Context, r *http.Request) (int, any, error) {
			return http.StatusNotFound, nil, e
		}))

		s := httptest.NewServer(h)
		defer s.Close()

		r, err := http.Get(s.URL)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, r.StatusCode)
		defer r.Body.Close()

		var rd httptils.JSONResponse[any]
		err = json.NewDecoder(r.Body).Decode(&rd)
		assert.Nil(t, err)
		assert.Equal(t, httptils.JSONResponse[any]{Error: e.Error()}, rd)
	})
}
