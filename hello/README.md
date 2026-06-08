# Day 1 — packages, `main`, and modules

**Every package has a name.** A file in a folder called `hello/` declares
`package hello` — the convention is that the package name matches the folder name.

**Runnable vs. library packages.** To *run* a program you need a package named
`main` **and** a `func main()` inside it. `main()` is the entry point. Anything
else (like my `hello` package) is a library valid Go, but it can't run on its
own; it exists to be imported.

```go
package main

import "fmt"

func main() {        // entry point — runs when the program starts
    fmt.Println("Hello, World!")
}
```

**`go.mod` and `go mod init`.** The `go.mod` file isn't created automatically
I make it once per project with:

```bash
go mod init go-learning-journey
```

The name I pass becomes the import path prefix for the whole project. It lives at
the root and covers every subfolder (the `hello/` folder does *not* need its own
`go.mod`). If I run `go mod init` with no name, Go tries to guess from my git
remote and errors out if it can't so passing the name explicitly is the safe bet.

**Importing my own package.** To use a function from another package I qualify it
with the package name:

```go
import "go-learning-journey/hello"

func main() {
    hello.Greet()    // package name . function name
}
```

- `hello` (before the dot) = the **package** name
- `Greet` (after the dot) = the **function** name inside it

Same pattern as the standard library: `fmt.Println`, `strings.Split`, etc.

**Exported names start with a capital letter.** I can only call `hello.Greet()`
because `Greet` is capitalized, which makes it *exported* (visible to other
packages). A lowercase `func greet()` would be private to its own package and
unreachable from `main`.