package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
)

func version() string {
  return "0.1.1 ALPHA"
}

func readInput() (string, error) {
  reader := bufio.NewReader(os.Stdin)
  text, err := reader.ReadString('\n')
  if err != nil {
    return "", err
  }
  text = text[0:len(text) - 1]
  return text, nil
}

// Check if a file exists
func fileOrFolderExists(filename string) bool {
  _, err := os.Stat(filename)
  if err == nil {
    return true
  }
  return false
}

// Report an error on a certain line
func errReport(err string, line int) {
  fmt.Println("Error on line " + strconv.Itoa(line + 1) + ": " + err)
}
