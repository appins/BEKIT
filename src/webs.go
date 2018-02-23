package main

import (
  "net/http"
  "strconv"
)

func startWebserver(port int, mainfolder string, filerr map[string]string) {
  http.Handle("/", func (w http.ResponseWriter, r *http.Request) {

  })

  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
