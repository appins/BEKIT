// Alex Anderson (c) 2018. This file interprets blocks of code
package main

import (
  "fmt"
  "io/ioutil"
  "os"
  "strconv"
  "strings"
)

// Run a block of code on a certain port
func run(block []string, port string, mode string, filename string) {
  mainFolder := ""
  var filerr map[string]string
  filerr = make(map[string]string)
  var f_in []string
  var f_out []string
  reportIp := false
  ignoreFakes := false
  ignoreErrors := false

  for line, comm := range(block) {
    args := cleanSplit(comm)

    // Loading from a file shifts all data down
    if mode == "load" {
      line++
    }

    if len(args) == 0 {
      continue
    }

    switch args[0] {
    case "":
    case " ":
    // Handle comments
    case "//":
    case "#":

    // Set up the root folder for the project (must contain index.html)
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

      // Refer to misc.go
      argument := wholeArgument(comm)

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

      argument := wholeArgument(comm)
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
    case "f":
      argument := wholeArgument(comm)
      pieces := strings.Split(argument, "->")

      if len(pieces) != 2 {
        errReport("f takes one input and one output seperated by a '->'", line)
        if !ignoreErrors {
          return
        }
        break
      }

      f_in = append(f_in, pieces[0])
      f_out = append(f_out, pieces[1])

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

  // Handle all three modes
  if mode == "onport" || mode == "load" {
    startWebserver(port, mainFolder, filerr, reportIp, f_in, f_out)
  }
  if mode == "save" {
    fildat := "port " + port + "\n"
    fildat += strings.Join(block, "\n")

    f, err1 := os.OpenFile(filename, os.O_WRONLY, 0644)
    defer f.Close()
    _, err2 := f.WriteString(fildat)

    if err1 != nil || err2 != nil {
      panic(err1)
      panic(err2)
    }
  }

}

// Load a file, then pass it into the interpreter
func loadFile(filename string) {
  if !fileOrFolderExists(filename) {
    fmt.Println("File does not exist!")
    return
  }

  dat, err := ioutil.ReadFile(filename)
  if err != nil {
    panic(err)
  }

  block := strings.Split(string(dat), "\n")
  // Check if the port is still valid
  port := cleanSplit(block[0])
  if len(port) != 2 || port[0] != "port" {
    fmt.Println("Loaded file's port config is incorrect!")
    return
  }

  _, err = strconv.Atoi(port[1])
  if err != nil {
    fmt.Println("Loaded file's port config is incorrect!")
    return
  }

  run(block[1:len(block)], port[1], "load", "")

}
