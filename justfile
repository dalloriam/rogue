NAME := "rogue"
PKG := "github.com/dalloriam/rogue"
BUILDDIR := "./bin"

clean:
    rm -r {{BUILDDIR}}

build:
    go build -o {{BUILDDIR}}/{{NAME}} ./cmd/main.go
