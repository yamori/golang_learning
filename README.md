# Learning GoLan

Enumerated `*.go` files, demonstrating functionality as posted in [gowebexamples.com](https://gowebexamples.com), but might differ slightly per my own tinkering.

## Warnings

- Each file is `package main` with its own function `main()`.  VSCode complains, but keeping them seperate with the `Makefile` works for the purposes of this repo, but not a design to be repeated.

## Commands

`go mod init yamori/go_learnings`

```
make 02buildrun
http://localhost/static/blah.txt

make 03buildrun
http://localhost/books/SirensOfTitan/page/42
```

## What I Learned

- Standalone executable (binary) go apps must be declared with `package main`, otherwise `go install` creates an archive file `*.a`