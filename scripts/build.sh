# <package-name>_<version>-<release-number>_<architecture>
PACKAGE_NAME="themes_0.0.1-1_x86_64"

BIN="$PWD/bin"
BUILD_PATH="$PWD/bin/$PACKAGE_NAME"
APT_REPOSITORY="$BIN/themes/apt-repo"

sudo apt-get install -y gcc dpkg-dev gpg
mkdir -p bin/themes
mkdir -p bin/themes_0.0.1-1_x86_64/usr/bin/
mkdir -p bin/themes_0.0.1-1_x86_64/DEBIAN/
mkdir -p $APT_REPOSITORY/dists/stable/main/binary-amd64
mkdir -p $APT_REPOSITORY/pool/main/
GOARCH=amd64 GOOS=linux go build -o bin/themes.exe
cp -r  $BIN/themes.exe $BUILD_PATH/usr/bin/themes

echo "Package: themes
Version: 0.0.1
Maintainer: jordan sudo-adduser-jordan@github.com
Depends: libc6, tree, mv, tar, golang-go
Architecture: amd64
Homepage: http://linuxthemes.org
Description: Install and manage themes and configurations" \
> $BUILD_PATH/DEBIAN/control

dpkg --build $BUILD_PATH
dpkg --info $BUILD_PATH.deb
dpkg --contents $BUILD_PATH.deb

cp $BIN/$PACKAGE_NAME.deb $APT_REPOSITORY/pool/main/ 
dpkg-scanpackages --arch amd64 $APT_REPOSITORY/pool/ > $APT_REPOSITORY/dists/stable/main/binary-amd64/Packages
cat $APT_REPOSITORY/dists/stable/main/binary-amd64/Packages | gzip -9 > $APT_REPOSITORY/dists/stable/main/binary-amd64/Packages.gz



