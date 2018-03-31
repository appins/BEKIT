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

		// If the user is not in a block, check if they used one of the commmands
		// defined in blocks.go
		if inblock == false {
			typ, blockstart, dat := startBlock(args)
			if blockstart == true {
				inblock = true
				// Handle the onport block command
				if typ == "onport" {
					ports = append(ports, dat)
					var emptyarr []string
					blocks = append(blocks, emptyarr)
					mode = "onport"
					continue
				}
				// Handle the save block command
				if typ == "save" {
					// Insert a .beks if the file path does not contain a dot
					if strings.Contains(dat, ".") {
						filename = dat
					} else {
						filename = dat + ".beks"
					}

					if !fileOrFolderExists(filename) {
						_, err := os.Create(filename)
						if err != nil {
							panic(err)
						}
					} else {
						var ow bool
						fmt.Println("File already exists...")
						for {
							fmt.Print("Overwrite? [y/n] ")
							overwrite, err := readInput()
							if err != nil {
								panic(err)
							}
							if overwrite == "y" || overwrite == "Y" || overwrite == "yes" {
								ow = true
								break
							}
							if overwrite == "n" || overwrite == "N" || overwrite == "no" {
								ow = false
								break
							}
						}
						if ow == false {
							inblock = false
							continue
						}
					}
					var emptyarr []string
					blocks = append(blocks, emptyarr)
					var port string
					for {
						fmt.Print("Port number: ")
						port, _ = readInput()
						_, err := strconv.Atoi(port)
						if err == nil {
							break
						}
					}
					ports = append(ports, port)
					mode = "save"

				}
				// Handle the load block command
				if typ == "load" {
					inblock = false
					// Load the file name if it contains a dot and load the name ending in
					// beks if it does not.
					if strings.Contains(dat, ".") {
						loadFile(dat)
						continue
					}
					loadFile(dat + ".beks")

				}

			}
			continue
		}

		if inp == "end" || inp == "}" {
			inblock = false
			// NOTE: This line is just for testing and easy execution
			run(blocks[len(blocks)-1], ports[len(ports)-1], mode, filename)
		}

		blocks[len(blocks)-1] = addToBlock(blocks[len(blocks)-1], inp)

	}
}
