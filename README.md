In order to run: `go run hello.go`

In order to build: `go build hello.go`

Install it under `%GOPATH%/bin`: `go install first-go-sample\hello` This will compile the code into an executable file.

## Go Workspace Folder Convention

<pre>

├───bin
├───pkg
│   └───windows_amd64
│       └───github.com
│           └───codegangsta
└───src
    ├───first-go-sample
    │   ├───cli-sample
    │   ├───hello
    │   └───http-sample
    └───github.com
        └───codegangsta
            └───cli
                └───autocomplete
</pre>