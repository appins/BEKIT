package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
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

// Seperate a string by words and remove every empty entry (whitespace remover)
func cleanSplit(line string) []string {
  var cleanWords []string
  // Replace tabs with spaces
  noTabs := strings.Replace(line, "\t", " ", -1)
  regSplit := strings.Split(noTabs, " ")
  for _, cont := range(regSplit) {
    if cont != "" {
      cleanWords = append(cleanWords, cont)
    }
  }

  return cleanWords
}

// When we do a clean split, we should be able to still extract the whole arg,
// expecially when several spaces may make a big difference in formatting
func wholeArgument(line string) string {
  var arg string
  reachedComm := false
  // Replace tabs with spaces
  noTabs := strings.Replace(line, "\t", " ", -1)
  regSplit := strings.Split(noTabs, " ")
  // Look for first reconized keyword (first object in clean array) and then
  // look for where the argument starts and gather the whole argument, including
  // spaces. You can test this with f text:  t->console and confirm they are
  // two spaces before the t
  for i, cont := range(regSplit) {
    if reachedComm == false && cont != "" {
      reachedComm = true
      continue
    }
    // If the first command has already occured and the second has started
    if reachedComm && cont != "" {
      arg = strings.Join(regSplit[i:len(regSplit)], " ")
      break
    }
  }
  return arg
}
