package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestNoSurf(t *testing.T) {
	var myh myHandler
	h := NoSurf(&myh)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("Type mismatch: Expected http.handler, got %T", v))

	}
}

func TestSessionLoad(t *testing.T) {
	var myh myHandler
	h := SessionLoad(&myh)

	switch v := h.(type) {
	case http.Handler:
	default:
		t.Error(fmt.Sprintf("Type mismatch: Expected http.handler, got %T", v))

	}
}
