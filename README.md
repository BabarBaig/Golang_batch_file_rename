# Golang_batch_rename_file

This program looks for a search_string in filenames and replaces it with a repl_string.<br>
It's helpful where a large number of files have to be renamed, making manual renaming impractical.<br>
Written in Go / Golang<br>
This allows for convenient multi-step renaming of files, progressively renaming them until they have the desired format.
<br>
To generate a WINDOWS 11, 64 bit executable on replit, use the following command (this fails sometimes):

    GOOS=windows GOARCH=amd64 go build -o 000_rename_files.exe
