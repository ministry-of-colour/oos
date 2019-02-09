package api

import "net/http"

type NullAPI struct{}

// Version returns the version string
func (h *NullAPI) Version() string { return "v1null" }

// Hello does not much at all - its a NULL driver
func (h *NullAPI) Hello(http.ResponseWriter, http.Request) {}

// Stock does not much at all - its a NULL driver
func (h *NullAPI) Stock(http.ResponseWriter, http.Request) {}
