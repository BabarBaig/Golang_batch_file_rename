# Golang_batch_rename_file

This program is written in Go/Golang.  It looks for a search_string in filenames in the current folder on Windows and <br>
replaces it with a repl_string.  This is helpful when a large number of files have to be renamed, making manual renaming <br>
tedious or impractical.<br><br>
It also allows for iteratively renaming files, progressively renaming them until they have the desired format.<br>
To generate a WINDOWS 11, 64 bit executable on replit, use the following command (this fails sometimes):

    GOOS=windows GOARCH=amd64 go build -o 000_rename_files.exe
