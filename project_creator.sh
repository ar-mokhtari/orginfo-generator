#!/bin/bash

# Read the user input
#gnome-terminal --tab --active --title="..::|| orginfo-generator ||::.." -- sh -c  "./12.sh"
#gnome-terminal -- bash -c "./12.sh; exec bash"
# shellcheck disable=SC2016
gnome-terminal --tab --active --title="..::|| orginfo-generator ||::.." -- sh -c '
echo please insert project name;
read project_name;
echo "--------start---------\n";
cd ~/go/src;
echo "--------cd ~/go/src---------\n";
mkdir -p -m777 $project_name/cli/generator;
echo "--------mkdir -p -m777 $project_name/cli/generator---------\n";
echo "package main

import \"github.com/ar-mokhtari/orginfo-generator\"

func main() {
   generator.New()
}

" > $project_name/cli/generator/main.go;
echo "--------$project_name/cli/generator/main.go;---------\n";
echo "package main

import (
	\"$project_name/adapter/storage\"
	\"$project_name/delivery/http\"
)

func init() {
	//initialise D.B.M.S. (Database management system)
	storage.Init()
	//initialise new Echo (web framework)
	http.Init()
}

func main() {

}" > $project_name/main.go;
echo "-------- $project_name/main.go;---------\n";
echo "module $project_name

go 1.19

require (
	github.com/ar-mokhtari/orginfo-generator v0.0.0-20221011095741-3289a676bde8
	github.com/dlclark/regexp2 v1.7.0
	github.com/go-ozzo/ozzo-validation v3.6.0+incompatible
	github.com/labstack/echo/v4 v4.9.0
	golang.org/x/crypto v0.0.0-20221010152910-d6f0a8c073c2
	gorm.io/driver/mysql v1.3.6
	gorm.io/gorm v1.23.8
)

require (
	github.com/asaskevich/govalidator v0.0.0-20210307081110-f21760c49a8d // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/labstack/gommon v0.3.1 // indirect
	github.com/mattn/go-colorable v0.1.11 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.1 // indirect
	golang.org/x/net v0.0.0-20211112202133-69e39bad7dc2 // indirect
	golang.org/x/sys v0.0.0-20211103235746-7861aae1554b // indirect
	golang.org/x/text v0.3.7 // indirect
)" > $project_name/go.mod && cd $project_name/cli/generator;
echo "--------$project_name/go.mod && cd $project_name/cli/generator---------\n";
go mod tidy;
echo "--------go mod tidy---------\n";
go get -u github.com/ar-mokhtari/orginfo-generator;
echo "--------go get -u github.com/ar-mokhtari/orginfo-generator---------\n";
go build;
echo "--------go build;---------\n";
export GO111MODULE="off";
go get ./.;
export GO111MODULE="auto";
go mod tidy;
read -p "Do you want to create a new domain too? [y,n]" doit
case $doit in  y|Y|yes|YES) cd ~/go/src/$project_name/cli/generator;
echo "insert domain name";
read domain_name;
echo "insert fields flag - read documents-readme.rm";
read fields_flag;
./generator new -domain_name=$domain_name -fields=$fields_flag;
echo "--------./generator new -domain_name=$domain_name -fields=$fields_flag---------\n";
go mod tidy;
cd ../..;
go get ./. ;;
esac;
cd ~/go/src/$project_name;
find . -type d | sed -e "s/[^-][^\/]*\//  |/g" -e "s/|\([^ ]\)/| - \1/";
echo "-------end (will close in 45sec) -------\n";
sleep 45;'
