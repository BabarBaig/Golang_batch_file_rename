# Golang_batch_rename_file

This program looks for a search_string in filenames and replaces it with a dest_string.<br>
It's helpful where a large number of files have to be renamed, making manual renaming impractical.<br>
Written in Go / Golang<br>
It saves the last directory where files were renamed.  So user doesn't have to enter directory path on every iteration.<br>
This allows for convenient multi-step renaming of files, progressively renaming them until they have the desired format.
<br>
To generate a WINDOWS 11, 64 bit executable on replit, use the following command:

    GOOS=windows GOARCH=amd64 go build -o rename_files.exe
