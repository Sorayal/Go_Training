# Commands

### Init new module

`Go go mod init url/modulename `

### Execute Go package in memory (not running a single file) in current directory

`Go go run . `

Compilation of Go Source happens in memory and will be executed from there.

### Execute in sub directory, example

`Go go run ./Random/main.go `

### Build Go Compilation (builds the actual exe / binary files). Name follows the module name.

`Go go build . `

### Run compiled programme, call via module name

`Go ./modulename`
