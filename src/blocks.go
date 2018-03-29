package main

import (
  "fmt"
  "strconv"
)

// This creates a block on a specific port
func startBlock(args []string) (string, bool, string) {
  if len(args) < 2 || (args[0] != "onport" && args[0] != "save" && args[0] != "load") {
    fmt.Println("Was expecting one of the following the sytax's:")
    fmt.Println("\tonport <port>")
    fmt.Println("\tsave <filename>")
    fmt.Println("\tload <filename>")
    fmt.Println(args[0])
    return "", false, "err"
  }

  if args[0] == "onport" {
    _, err := strconv.Atoi(args[1])
    if err != nil {
      fmt.Println("The format of port was not a number")
      return "", false, "err"
    }
    return "onport", true, args[1]
  }
  if args[0] == "save" || args[0] == "load" {
    if len(args[1]) > 0 {
      return args[0], true, args[1]
    }
    fmt.Println("'" + args[0] + "' requires a valid filename")
  }

  return "", false, "err"
}

// This function adds a string to the block
func addToBlock(block []string, inp string) []string {
  if inp == "" {
    return block
  }

  block = append(block, inp)

  return block
}
