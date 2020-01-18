# chi-router-example

Starter pack for chi

## Setup

Clone this repo, and rename the directory for your own project. Remove the .git
repo, and start over.

```
cd myserver
rm -rf .git
git init
```

Go modules requires that we declare a module path, so you'll need to invoke
`go mod init` with your own repo path. For example.

```
go mod init github.com/me/myserver
```

Then you can build. Dependencies will be fetched automatically.

```
go build
```

## Running

We just have a single `func main` declared in a single file, so just

```
go run main.go
```

If you have [`entr`](https://eradman.com/entrproject/) installed, you can watch 
for changes to files and automatically re-run the server.

```
ls *.go | entr -r go run main.go
```

## Cross Compiling

As long as we don't have any cgo dependencies, we can cross-compile for Linux.

```
GOOS=linux GOARCH=amd64 go build
```


## References

* [chi router library](https://github.com/go-chi/chi)
* [Go standard library package: net/http](https://golang.org/pkg/net/http/)

