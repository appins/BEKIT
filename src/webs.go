package main

import (
  "net/http"
)

func startWebserver(port int, mainfolder string, filerr map[string]string) {
  http.Handle("/", func (w http.ResponseWriter, r *http.Request) {

  })
}
