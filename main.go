package main
/*
go build -o rename_files.exe main.go
~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
To build on Replit: (fails sometimes)
GOOS=windows GOARCH=amd64 go run   rename_files.go
GOOS=windows GOARCH=amd64 go build rename_files.go
*/
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("2023-10-31: This program looks for a [search_string] in file names and replaces it " +
		"with a [replacement_string].")
	rename_files1()
}

func rename_files1() {
	resp := ""
	for {
		fmt.Println("Press any key to continue, q to quit: ")
		fmt.Scanln(&resp)
		if resp == "q" { break }
		rename_files2()
	}
}

func rename_files2() {
	files, _ := os.ReadDir("./")
	fnameCur := ""
	fnameNew := ""

	fmt.Print("Enter srchStr:\t")
	reader := bufio.NewReader(os.Stdin)
	srchStr, _ := reader.ReadString('\n')
	srchStr = srchStr[:len(srchStr)-2]

	fmt.Print("Enter replStr:\t")
	replStr, _ := reader.ReadString('\n')
	replStr = replStr[:len(replStr)-2]

	rename_all := false
	for _, f := range files {
		fnameCur = f.Name()
	    // if current file isn't a match, move to next file in directory
		if strings.Index(fnameCur, srchStr) == -1 { continue }

		// found a match
		fnameNew = strings.Replace(fnameCur, srchStr, replStr, 1)
		if !rename_all {
			fmt.Println("[" + fnameCur + "] ->\n" + "[" + fnameNew + "]")
			fmt.Print("<Enter> y(es) n(o) q(uit) a(ll)\t")
			srchStr, _ := reader.ReadString('\n')
			if (srchStr[0] == 'n'){ continue }
			if (srchStr[0] == 'q'){ break }
			if (srchStr[0] == 'a'){ rename_all = true }
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
