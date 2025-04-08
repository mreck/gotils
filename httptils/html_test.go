package httptils_test

import (
	"context"
	"errors"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mreck/gotils/httptils"

	"github.com/stretchr/testify/assert"
)

func TestHTMLHandler(t *testing.T) {
	t.Run("handle success", func(t *testing.T) {
		h := http.NewServeMux()
		tmpl, _ := template.New("test").Parse(`a {{ index . "a" }}`)
		j := httptils.NewHTMLHandler(context.TODO(), tmpl)

		h.HandleFunc("/", j.H(func(ctx context.Context, r *http.Request) (int, string, httptils.D, error) {
			return http.StatusOK, "test", httptils.D{"a": 1}, nil
		}))

		s := httptest.NewServer(h)
		defer s.Close()

		r, err := http.Get(s.URL)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, r.StatusCode)
		defer r.Body.Close()

		b, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Equal(t, []byte("a 1"), b)
	})

	t.Run("handle failure", func(t *testing.T) {
		h := http.NewServeMux()
		tmpl, _ := template.New("error.html").Parse(`error: {{ index . "Error" }}`)
		j := httptils.NewHTMLHandler(context.TODO(), tmpl)
		e := errors.New("test error")

		h.HandleFunc("/", j.H(func(ctx context.Context, r *http.Request) (int, string, httptils.D, error) {
			return http.StatusNotFound, "", nil, e
		}))

		s := httptest.NewServer(h)
		defer s.Close()

		r, err := http.Get(s.URL)
		assert.Nil(t, err)
		assert.Equal(t, http.StatusNotFound, r.StatusCode)
		defer r.Body.Close()

		b, err := io.ReadAll(r.Body)
		assert.Nil(t, err)
		assert.Equal(t, []byte("error: test error"), b)
	})
}
