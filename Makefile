hello:
	echo "hello"

01buildrun:
	go build -o bin/01_hello_world 01_hello_world.go
	./bin/01_hello_world

02buildrun:
	go build -o bin/02_http_server 02_http_server.go
	./bin/02_http_server

03buildrun:
	go build -o bin/03_mux_routing 03_mux_routing.go
	./bin/03_mux_routing

04buildrun:
	go build -o bin/04_mysql_db 04_mysql_db.go
	./bin/04_mysql_db

clean:
	rm bin/*