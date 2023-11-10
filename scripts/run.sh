
#!/bin/sh
srcPath="cmd"
pkgFile="main.go"
app="nucleo-template"
src="$srcPath/$app/$pkgFile"

printf "\nStart running: $app\n"
time go run $src
printf "\nStopped running: $app\n\n"