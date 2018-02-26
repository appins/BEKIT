// Alex Anderson (c) 2018. This file interprets blocks of code
package main

import (
  "fmt"
  "strings"
)

// Run a block of code on a certain port
func run(block []string, port int) {
  mainFolder := ""
  var filerr map[string]string
  filerr = make(map[string]string)
  reportIp := false
  ignoreFakes := false
  ignoreErrors := false

  for line, comm := range(block) {
    args := strings.Split(comm, " ")
    switch args[0] {

    // Set up the main, or root, folder for the project (must exist)
    case "set-main":
      fallthrough
    case "main":
      fallthrough
    case "root":
      if len(args) == 1 {
        errReport("main folder was not stated.", line)
        if !ignoreErrors {
          return
        }
        break
      }

      argument := strings.Join(args[1:len(args)], " ")

      // Make sure file exists (and contains index.html)
      if !fileOrFolderExists(argument) {
        errReport("the main folder that was specified does not exist.", line)
        if !ignoreErrors {
          return
        }
        break
      }
      if !fileOrFolderExists(argument + "/index.html") {
        errReport("main folder must contain 'index.html' to run.", line)
        if !ignoreErrors {
          return
        }
        break
      }

      mainFolder = argument
      break
    case "filerr":
      fallthrough
    case "rrfile":
      if len(args) == 1 {
        errReport("filerr needs an argument.", line)
        if !ignoreErrors {
          return
        }
        break
      }

      argument := strings.Join(args[1:len(args)], " ")
      files := strings.Split(argument, "->")

      if len(files) != 2 {
        errReport("filerr only takes 2 files separated by a '->'", line)
        if !ignoreErrors {
          return
        }
        break
      }

      filerr[files[0]] = files[1]
      break
    case "log-ip":
      fallthrough
    case "logip":
      reportIp = true
      break
    case "force":
      ignoreErrors = true
      ignoreFakes = true
      fmt.Println("Ignoring errors and commands that do not exist")
    case "force-lite":
      ignoreFakes = true
      fmt.Println("Ignoring commands that do not exist")
    // If the command is not found, report it as an error, unless `force` or `force-lite`
    default:
      errReport("command not found.", line)
      if !ignoreFakes {
        return
      }
    }
  }

  if mainFolder == "" {
    fmt.Println("You must set a main folder to run your server!")
    return
  }

  startWebserver(port, mainFolder, filerr, reportIp)
}
