#   📃 MY LS
## DESCRIPTION
This project consists on creating your own `ls` command. The `ls` command shows you the files and folders of the directory specified after the command. By exclusion of this directory, it shows the files and folders of the present directory. It must incorporate in your project at least the following flags of the `ls` command:
* ``-l`` : list with long format
* ``-r`` : list in reverse order
* ``-a`` : list all files including hidden file starting with '.'
* ``-R`` : list recursively directory tree
* ``-t`` : sort by time & date


## USAGE
```sh
go run . [OPTIONS] [FILE|DIR]

go run . -l         # ls -l
go run . -r         # ls -r
go run . folder/    # ls folder/
go run . -lraRt     # ls -lraRt
```

##  AUTHORS
+   Bounama Coulibaly
+   Serigne Saliou Mbacké Mbaye
+   Mamoudou Ndiaye
