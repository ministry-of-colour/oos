package api

import (
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

// Http is an implementation of the API over HTTP
type Http struct {
	log *logrus.Logger
}

// NewHttp returns a new HttpAPI
func NewHttp(l *logrus.Logger) *Http {
	return &Http{log: l}
}

// Version returns the current version number as a string
func (h *Http) Version() string {
	return "v1"
}

func (h *Http) Log(r *http.Request, txt string) {
	h.log.WithFields(logrus.Fields{
		"addr":   r.RemoteAddr,
		"proto":  r.Proto,
		"method": r.Method,
		"URI":    r.RequestURI,
	}).Println(txt)
}

// Default does not much at all yet
func (h *Http) Default(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Nothing to see here\n"))
	h.Log(r, "Default")
}

// Hello writes hello world as a response
func (h *Http) Hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!" + "\n"))
	h.Log(r, "Hello")
}

// Stock writes a stock update in JSON
func (h *Http) Stock(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	type stock struct {
		Item  string
		Descr string
		Qty   string
	}

	type stocks struct {
		Note  string
		Items []stock
	}

	// TODO - get this from the DB
	s := stocks{
		Note: "We have plenty left in stock",
		Items: []stock{
			stock{Item: "abc", Descr: "the abc item", Qty: "10"},
			stock{Item: "def", Descr: "the def item", Qty: "20"},
			stock{Item: "xyz", Descr: "the xyz item", Qty: "6"},
		},
	}

	json.NewEncoder(w).Encode(s)
	h.Log(r, "Stock")
}
