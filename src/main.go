// Alex Anderson (c) 2018. This file operates other functions.
package main

import (
  "fmt"
  "strings"
)

func main() {
  // Each block has a port associated with it
  var blocks [][]string
  var ports []int

  // Variable to tell if the user is writing inside of outside of a block
  inblock := false

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

    if inblock == false {
      blockstart, port := startBlock(args)
      if blockstart == true {
        inblock = true
        ports = append(ports, port)
        var emptyarr []string
        blocks = append(blocks, emptyarr)
      }
      continue
    }

    if inp == "end" || inp == "}" {
      inblock = false
      // NOTE: This line is just for testing and easy execution
      run(blocks[len(blocks) - 1], ports[len(ports) - 1])
    }

    blocks[len(blocks) - 1] = addToBlock(blocks[len(blocks) - 1], inp)

  }
}
