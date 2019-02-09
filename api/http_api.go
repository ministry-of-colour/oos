package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/husobee/vestigo"
	"github.com/sirupsen/logrus"
)

// HTTP is an implementation of the API over HTTP
type HTTP struct {
	log *logrus.Logger
	db  *pgDB
}

// NewHTTP returns a new HTTP API
func NewHTTP(l *logrus.Logger) *HTTP {
	return &HTTP{log: l, db: newDB(l)}
}

// Version returns the current version number as a string
func (h *HTTP) Version() string {
	return "v1"
}

func (h *HTTP) doLog(r *http.Request, t time.Time, txt string) {
	h.log.WithFields(logrus.Fields{
		"Duration": time.Since(t).String(),
		"addr":     r.RemoteAddr,
		"proto":    r.Proto,
		"method":   r.Method,
		"URI":      r.RequestURI,
	}).Println(txt)
}

// Default does not much at all yet
func (h *HTTP) Default(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Nothing to see here\n"))
	h.doLog(r, t1, "Default")
}

// Hello writes hello world as a response
func (h *HTTP) Hello(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, World!" + "\n"))
	h.doLog(r, t1, "Hello")
}

// StockReport writes a stock report in JSON
func (h *HTTP) StockReport(w http.ResponseWriter, r *http.Request) {
	t1 := time.Now()
	brand := vestigo.Param(r, "brand")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	items, err := h.db.StockReport(brand)
	if err != nil {
		h.log.WithError(err).Warn("error fetching all the items")
	}
	json.NewEncoder(w).Encode(items)
	h.doLog(r, t1, "Stock")
}
