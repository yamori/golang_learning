# Learning GoLang

Enumerated `*.go` files, demonstrating functionality as posted in [gowebexamples.com](https://gowebexamples.com), but might differ slightly per my own tinkering.

## Warnings

- Each file is `package main` with its own function `main()`.  VSCode complains, but keeping them seperate with the `Makefile` works for the purposes of this repo, but not a design to be repeated.

## Commands

### Building this repo

`go mod init yamori/go_learnings`

### How to invoke each enumerated file

Web:

```
 # HTTP Server
make 02buildrun
http://localhost/static/blah.txt

 # Routing
make 03buildrun
http://localhost/books/SirensOfTitan/page/42

 # Embedded Assets
make 05buildrun
http://localhost:80
```

DB:

```
 # Connect to mySQL DB
docker pull mysql
docker run --name some-mysql -e MYSQL_ROOT_PASSWORD=mysecretpw -e MYSQL_DATABASE=root -d -p 127.0.0.1:3306:3306 mysql
make 04buildrun
```

## What I Learned

- Template and Assets: it's easy enough to inject dynamic values into templates (`template.Execute(w, data)`), and then serve static assets as `fs := http.FileServer(http.Dir("assets/"))`, but the problem is for each asset that's another call to the DB.  For `06_assets.go` I attempted to `embed` a CSS file at compile-time, and use the string value to be injected into the template.  But this results in a value of `ZgotmplZ` which turns out is a special value of GO, that unsafe content had been detected and was substituted.  I did not get far enough into investigating how to escape this, rather I fell back on offering the static asset as first described as a seperate request to the server.  After all, this is a browser's job, but I do wonder what the performance trade-offs would be and I'm sure it would be dependent on the hosting infrastructure/topology.
- Standalone executable (binary) go apps must be declared with `package main`, otherwise `go install` creates an archive file `*.a`
- Imports using the blank identifier: `_ "github.com/go-sql-driver/mysql"`.  Imports a package solely for its side-effects (initialization).  Can then be used by a standard library for example `sql.Open("mysql", "root:mysecretpw@(127.0.0.1:3306)/root")`

### Embed

As of Go `1.16.?`, there's an `embed` module to attach static assets to the ultimate compiled Go binary.  Useful when creating webapp if it's only a few assets (html, css js, etc.).  Probably not a good idea to cram in a lot of assets, but if going Serverless this could be useful.  This demonstrated in `05_templates.go`.

```
 # not within a function, note the regex pattern for matching what is to be embedded
//go:embed templates/*
var assetData embed.FS

 # within the function, grab the particular asset from the object
t, err := template.ParseFS(assetData, "templates/layout.html")

 # in the http.HandleFunc
t.Execute(of io.Writer, data struct instance)
```