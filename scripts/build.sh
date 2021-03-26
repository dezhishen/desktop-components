path=$1
go mod tidy
go mod vendor
qtdeploy build desktop $path