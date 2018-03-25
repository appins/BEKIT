// Alex Anderson (c) 2018. This file operates other functions.
package main

import (
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  // Each block has a port associated with it
  var blocks [][]string
  var ports []string

  // Variable to tell if the user is writing inside of outside of a block
  inblock := false
  mode := ""
  filename := ""

  fmt.Println("Backend Kit " + version())
  for {
    if inblock == true {
      fmt.Print("    { ")
    } else {
      fmt.Print("> ")
    }

    inp, err := readInput()
    if err != nil {
      fmt.Println("Error reading input! Exiting program.")
      return
    }

    // Split up the input string by spaces
    args := strings.Split(inp, " ")

    if inp == "exit" || inp == "stop" || inp == "quit" {
      fmt.Println("Goodbye!")
      return
    }

    if inblock == false {
      typ, blockstart, dat := startBlock(args)
      if blockstart == true {
        inblock = true
        if typ == "onport" {
          ports = append(ports, dat)
          var emptyarr []string
          blocks = append(blocks, emptyarr)
          mode = "onport"
          continue
        }
        if typ == "save" {
          filename = dat
          if !fileOrFolderExists(filename) {
            _, err := os.Create(filename)
            if err != nil {
              panic(err)
            }
          }
          var emptyarr []string
          blocks = append(blocks, emptyarr)
          var port string
          for {
            fmt.Print("Port number: ")
            port, _ := readInput()
            _, err := strconv.Atoi(port)
            if err == nil {
              break
            }
          }
          ports = append(ports, port)
          mode = "save"
        }
        if typ == "load" {
          inblock = false
          loadFile(dat)
        }

      }
      continue
    }

    if inp == "end" || inp == "}" {
      inblock = false
      // NOTE: This line is just for testing and easy execution
      run(blocks[len(blocks) - 1], ports[len(ports) - 1], mode, filename)
    }

    blocks[len(blocks) - 1] = addToBlock(blocks[len(blocks) - 1], inp)

  }
}
