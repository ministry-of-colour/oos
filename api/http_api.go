package api

import (
	"encoding/json"
	"net/http"
	"time"

	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

// Http is an implementation of the API over HTTP
type Http struct {
	log *logrus.Logger
	db  *pgDB
}

// NewHttp returns a new HttpAPI
func NewHttp(l *logrus.Logger) *Http {
	return &Http{log: l, db: newDB(l)}
}

// Version returns the current version number as a string
func (h *Http) Version() string {
	return "v1"
}

func (h *Http) Log(r *http.Request, t time.Time, txt string) {
	h.log.WithFields(logrus.Fields{
		"Duration": time.Since(t).String(),
		"addr":     r.RemoteAddr,
		"proto":    r.Proto,
		"method":   r.Method,
		"URI":      r.RequestURI,
	}).Println(txt)
}

// Default does not much at all yet
func (h *Http) Default(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Nothing to see here\n"))
	h.Log(r, t1, "Default")
}

// Hello writes hello world as a response
func (h *Http) Hello(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!" + "\n"))
	h.Log(r, t1, "Hello")
}

// Stock writes a stock update in JSON
func (h *Http) Stock(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	items, err := h.db.StockReport()
	if err != nil {
		h.log.WithError(err).Warn("error fetching all the items")
	}
	json.NewEncoder(w).Encode(items)
	h.Log(r, t1, "Stock")
}
