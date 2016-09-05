# Buile x86 and x86_64 version
#
VER:=$(shell cat VERSION)
BINARY=synoindex
BINARY_SERVER=synoindex-server
INIFILE=simple-synoindex-server.ini

preinit:
	mkdir -p dist/x86
	mkdir -p dist/x86_64
	mkdir -p dist/darwin

x86_64: preinit clean
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY) simple-synoindex-client.go simple-synoindex-utils.go
	env GOOS=linux GOARCH=amd64 go build -o $(BINARY_SERVER) simple-synoindex-server.go simple-synoindex-utils.go
	mv $(BINARY) dist/x86_64
	mv $(BINARY_SERVER) dist/x86_64
	cp $(INIFILE) dist/x86_64

x86: preinit clean
	env GOOS=linux GOARCH=386 go build -o $(BINARY) simple-synoindex-client.go simple-synoindex-utils.go
	env GOOS=linux GOARCH=386 go build -o $(BINARY_SERVER) simple-synoindex-server.go simple-synoindex-utils.go
	mv $(BINARY) dist/x86
	mv $(BINARY_SERVER) dist/x86
	cp $(INIFILE) dist/x86

darwin: preinit clean
	env GOOS=darwin GOARCH=amd64 go build -o $(BINARY) simple-synoindex-client.go simple-synoindex-utils.go
	env GOOS=darwin GOARCH=amd64 go build -o $(BINARY_SERVER) simple-synoindex-server.go simple-synoindex-utils.go
	mv $(BINARY) dist/darwin
	mv $(BINARY_SERVER) dist/darwin
	cp $(INIFILE) dist/darwin

package_x86:
	cd dist/x86 && rm -f *.zip && \
		zip simple-synoindex-server-x86-$(VER).zip $(BINARY) $(BINARY_SERVER) $(INIFILE)

package_x86_64:
	cd dist/x86_64 && rm -f *.zip && \
		zip simple-synoindex-server-x86_64-$(VER).zip $(BINARY) $(BINARY_SERVER) $(INIFILE)

clean:
	rm -f $(BINARY)
	rm -f $(BINARY_SERVER)

all: x86 x86_64 package_x86 package_x86_64

# All are .PHONY for now because dependencyness is hard
.PHONY: all
