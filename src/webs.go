package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

// Used for starting the actual web server
func startWebserver(port string, mainfolder string, filerr map[string]string,
	reportIp bool, f_in []string, f_out []string) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		path := r.URL.Path

		// Add index.html onto the end of the path if it ends with /
		if path[len(path)-1] == '/' {
			path += "index.html"
		}

		for filename, reroute := range filerr {
			if filename == path || "/"+filename == path {
				if reroute == "null" {
					http.NotFound(w, r)
					fmt.Println("A user requested " + path + " which was rerouted to null")
					if reportIp {
						ip, _, _ := net.SplitHostPort(r.RemoteAddr)
						fmt.Println("The ip for that request is: " + ip)
					}
					return
				}
				path = reroute
				break
			}
		}

		// Read from the file
		dat, err := os.Open(mainfolder + path)

		// Handle all user created input-output functions
		for i, _ := range f_in {
			input_parts := strings.Split(f_in[i], ":")
			in_dat := ""
			output_parts := strings.Split(f_out[i], ":")
			writeToOutput := false

			switch input_parts[0] {
			// Handle both the request and 404 input keywords
			case "request":
				fallthrough
			case "404":
				if err != nil || input_parts[0] == "request" {
					writeToOutput = true
					if input_parts[1] == "ip" {
						ip, _, _ := net.SplitHostPort(r.RemoteAddr)
						in_dat = string(ip)
						break
					}
					if input_parts[1] == "file" {
						in_dat = string(path)
						break
					}
					if input_parts[1] == "is404" {
						if err != nil {
							in_dat = "true"
							break
						}
						in_dat = "false"
						break
					}
					writeToOutput = false
				}
				break
			// Handles form values
			case "form":
				in_dat = r.FormValue(strings.Join(input_parts[1:len(input_parts)], ":"))
				if in_dat != "" {
					writeToOutput = true
				}
			// Handles static text from the CLI
			case "text":
				writeToOutput = true
				if input_parts[1] == "newline" {
					in_dat = "\n"
					break
				}
				in_dat = strings.Join(input_parts[1:len(input_parts)], ":")
			}
			// Check if we should write to output
			if writeToOutput {
				// Write to console (stdout)
				if output_parts[0] == "console" {
					fmt.Print(in_dat)
					if len(output_parts) == 1 || output_parts[1] != "nonewline" {
						fmt.Println()
					}
				}
				// Write to a file
				if output_parts[0] == "write" {
					if len(output_parts) > 1 {
						filename := strings.Join(output_parts[1:len(output_parts)], ":")
						if !fileOrFolderExists(filename) {
							_, err := os.Create(filename)
							if err != nil {
								panic(err)
							}
						}
						f, err1 := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
						defer f.Close()
						_, err2 := f.WriteString(in_dat)
						if err1 != nil || err2 != nil {
							panic(err1)
							panic(err2)
						}
						f.Close()
					}
				}
			}
		}

		// Missing: 404 errors
		if err != nil {
			fmt.Println("User requested missing file (" + path + ")")
			if reportIp {
				ip, _, _ := net.SplitHostPort(r.RemoteAddr)
				fmt.Println("The ip for that request is: " + ip)
			}

			http.NotFound(w, r)
			return
		}

		// Otherwise, get type of file
		var contentType string
		fileExt := strings.Split(path, ".")[len(strings.Split(path, "."))-1]

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
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
