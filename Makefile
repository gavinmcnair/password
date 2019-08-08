password: *.go
	go get
	go build -o password *.go

install: password
	sudo install -m 755 password /usr/local/bin
