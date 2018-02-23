package main

import (
  "fmt"
  "strconv"
)

// This creates a block on a specific port
func startBlock(args []string) (bool, int) {
  if len(args) != 2 || args[0] != "onport" {
    fmt.Println("Was expecting an opening following the sytax:")
    fmt.Println("\tonport <port>")
    fmt.Println(args[0])
    return false, -1
  }

  port, err := strconv.Atoi(args[1])
  if err != nil {
    fmt.Println("The format of port was not a number")
    return false, -1
  }

  return true, port
}

// This function adds a string to the block
func addToBlock(block []string, inp string) []string {
  if inp == "" {
    return block
  }

  block = append(block, inp)

  return block
}
