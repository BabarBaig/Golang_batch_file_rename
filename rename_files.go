package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func changeDir() {
	fileSettings := ".settings.txt"
	dirDest := ""
	cwd, err := os.Getwd()
	file, err := os.Open(fileSettings)
	reader := bufio.NewReader(os.Stdin)

	if err == nil { // Successfully opened settings file
		fmt.Println("Found file [" + fileSettings + "] in dir [" + cwd + "]")
		scanner := bufio.NewScanner(file)
		scanner.Scan()
		dirDest = scanner.Text()
		fmt.Print("Do you want to switch to dir [" + dirDest + "]? [y, n, eXit]: ")
		resp, _ := reader.ReadString('\n')
		resp = strings.ToLower(resp)
		if resp[0] == 'x' {
			os.Exit(0)
		} else if resp[0] == 'y' {
			fmt.Println("cd [" + dirDest + "]")
			os.Chdir(dirDest)
			return
		}
	}

	file.Close()
	fmt.Print("Enter directory:\t")
	dirDest, _ = reader.ReadString('\n')
	dirDest = dirDest[:len(dirDest)-2]
	os.Chdir(dirDest)
	curDir, _ := os.Getwd()
	fmt.Println("cd [" + curDir + "]")
}

func renameFiles() {
	files, _ := ioutil.ReadDir("./")
	fnameCur := ""
	fnameNew := ""

	for _, f := range files {
		fnameCur = f.Name()
		fmt.Println(fnameCur)
	}

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
		if strings.Index(fnameCur, srchStr) == -1 {
			continue
		}
		fnameNew = strings.Replace(fnameCur, srchStr, replStr, 1)
		fmt.Println("Rename\n[" + fnameCur + "]\n" + "[" + fnameNew + "]")
		os.Rename(fnameCur, fnameNew)
	}
}

func main() {
	fmt.Println("This program looks for a search string in file names and replaces it " +
		"with a replacement string.")
	changeDir()
	reader := bufio.NewReader(os.Stdin)
	for {
		renameFiles()
		fmt.Println("Press any key to continue, CTRL-c to eXit: ")
		reader.ReadString('\n')
	}
}
