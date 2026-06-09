# Go Learning Journey

## How to run

From the project root:

```bash
go run .
```

This builds and runs the `main` package. (See the Learning Log for why one file
must be `package main` with a `func main()`.)

---

## Project structure

```
go-learning-journey/
├── go.mod          # module definition (created with `go mod init`)
├── main.go         # package main, func main()  ← the runnable entry point
└── hello/
   └── hello.go    # package hello              ← a library, imported by main
   └── README.md   
└── variable/
    └── variable.go # package variable           ← a library, imported by main
    └── README.md
```