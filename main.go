package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	fmt.Println("This program looks for a [search_string] in file names and replaces it " +
		"with a [replacement_string].")
	changeDir()
	reader := bufio.NewReader(os.Stdin)
	for {
		renameFiles()
		fmt.Println("Press any key to continue, CTRL-c to eXit: ")
		reader.ReadString('\n')
	}
}

func changeDir() {
	dirDest := ""
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter directory:\t")
	dirDest, _ = reader.ReadString('\n')
	dirDest    = dirDest[:len(dirDest)-2]
	os.Chdir(dirDest)
	curDir, _ := os.Getwd()
	fmt.Println("cd [" + curDir + "]")
}

func renameFiles() {
	files, _ := ioutil.ReadDir("./")
	fnameCur := ""
	fnameNew := ""

	fmt.Println("Replace a search string [srchStr], with replacement string [replStr]")
	fmt.Print("Enter srchStr:\t")
	reader := bufio.NewReader(os.Stdin)
	srchStr, _ := reader.ReadString('\n')
	srchStr = srchStr[:len(srchStr)-2]

	fmt.Print("Enter replStr:\t")
	replStr, _ := reader.ReadString('\n')
	replStr = replStr[:len(replStr)-2]

	for _, f := range files {
		fnameCur = f.Name()
    // if current file isn't a match, move to next file in directory
		if strings.Index(fnameCur, srchStr) == -1 {
			continue
		}
		fnameNew = strings.Replace(fnameCur, srchStr, replStr, 1)
		fmt.Println("[" + fnameCur + "] ->\n" + "[" + fnameNew + "]")
		fmt.Print("Press 'n' to skip, <Enter> to rename file, CTRL-c to eXit: ")
		srchStr, _ := reader.ReadString('\n')
		if (srchStr[0] == 'n'){  // don't rename current file
			continue
		}

    err := os.Rename(fnameCur, fnameNew)
    // If fnameNew already exists, access-denied error is generated.
    if err != nil {
      e := err.(*os.LinkError)
      fmt.Println("Err: ", e.Err)
      fmt.Println("Does file [" + fnameNew + "] already exist?")
    }
	}
}
