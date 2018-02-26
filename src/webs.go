package main

import (
  "fmt"
  "io"
  "log"
  "net"
  "net/http"
  "os"
  "strconv"
  "strings"
)

// Used for starting the actual web server
func startWebserver(port int, mainfolder string, filerr map[string]string,
  reportIp bool) {
  http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
    defer r.Body.Close()

    path := r.URL.Path

    // Add index.html onto the end of the path if it ends with /
    if path[len(path) - 1] == '/' {
      path += "index.html"
    }

    for filename, reroute := range(filerr) {
      if filename == path || "/" + filename == path {
        if reroute == "null" {
          http.NotFound(w, r)
          fmt.Println("A user requested " + path + " which was rerouted to null")
          if reportIp {
            ip, _, _ := net.SplitHostPort(r.RemoteAddr)
            fmt.Println("The ip for that request is: " + ip )
          }
          return
        }
        path = reroute
        break
      }
    }

    // Read from the file
    dat, err := os.Open(mainfolder + path)

    // Missing: 404 errors
    if err != nil {
      fmt.Println("User requested missing file (" + path + ")")
      if reportIp {
        ip, _, _ := net.SplitHostPort(r.RemoteAddr)
        fmt.Println("The ip for that request is: " + ip )
      }

      http.NotFound(w, r)
      return
    }

    // Otherwise, get type of file
    var contentType string
    fileExt := strings.Split(path, ".")[len(strings.Split(path, ".")) - 1]

    switch fileExt {
    case "css":
      contentType = "text/css"
    case "html":
      contentType = "text/html"
    case "js":
      contentType = "application/javascript"
    default:
      contentType = "text/plain"
    }

    w.Header().Add("Content-Type", contentType)
    io.Copy(w, dat)

  })

  fmt.Println("Starting server on port", port)
  log.Fatal(http.ListenAndServe(":" + strconv.Itoa(port), nil))
}
