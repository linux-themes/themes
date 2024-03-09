BINARY_NAME=themes

build:
	mkdir -p  bin
	GOARCH=amd64 GOOS=linux go build  -o bin/themes.exe
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
	sudo nala remove themes -y
	sudo nala install ./bin/themes_0.0.1-1_x86_64.deb -y
clean:
	go clean
	rm -rf bin