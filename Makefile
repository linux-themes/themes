BINARY_NAME=themes

build:
	mkdir -p  bin
	GOARCH=amd64 GOOS=linux go build  -o bin/${BINARY_NAME}
run:
	./bin/themes.exe
	# ./bin/themes help
	# ./bin/themes list  
	# ./bin/themes list icons 
	# ./bin/themes list themes 
	# ./bin/themes list create 
	# ./bin/themes list build 
	# ./bin/themes list install 
	# ./bin/themes list install icons https://github.com/sudo-adduser-jordan/mint-y-winx/raw/main/mint-y-winx.tar.xz
	# ./bin/themes list install themes https://github.com/sudo-adduser-jordan/mint-y-winx/raw/main/mint-y-winx.tar.xz
	# ,/bin/themes set icons package_1
	# ,/bin/themes set themes package_1
	# ./bin/themes remove icons package_1
	# ./bin/themes remove themes package_1
debian:
	./scripts/build.sh
clean:
	go clean
	rm -rf bin