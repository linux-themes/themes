# <package-name>_<version>-<release-number>_<architecture>

$PACKAGE_NAME="hello-world"
$VERSION=
$RELEASE_NUMBER=
$ARCHITECTURE=
$FILE_NAME=$PACKAGE_NAME_$VERSION-$RELEASE_NUMBER_$ARCHITECTURE

mkdir -p ~/example/$FILE_NAME
cd ~/example/$FILE_NAME

mkdir -p usr/bin
cp ~/example/$PACKAGE_NAME-program/$PACKAGE_NAME usr/bin/.

mkdir -p ~/example/$FILE_NAME/DEBIAN

echo "Package: hello-world
Version: 0.0.1
Maintainer: example <example@example.com>
Depends: libc6
Architecture: amd64
Homepage: http://example.com
Description: A program that prints hello" \
> ~/example/$FILE_NAME/DEBIAN/control
# Expected
# ~/example/$FILE_NAME/usr/bin/hello-world
# ~/example/$FILE_NAME/DEBIAN/control

# Build
dpkg --build ~/example/$FILE_NAME
# Inspect 
# dpkg-deb --info ~/example/hello-world_0.0.1.deb
# new Debian package, version 2.0.
# size 2832 bytes: control archive=336 bytes.
#     182 bytes,     7 lines      control
# Package: hello-world
# Version: 0.0.1
# Maintainer: example <example@example.com>
# Depends: libc6
# Architecture: amd64
# Homepage: http://example.com
# Description: A program that prints hello

# View Content
# dpkg-deb --contents ~/example/hello-world_0.0.1.deb


# This package can then be installed using the -f option under apt-get install:
sudo apt-get install -f ~/example/hello-world_0.0.1-1_amd64.deb
# Then once installed, you can verify it works with commands like:
which hello-world
# and
hello-world
# which should output /usr/bin/hello-world and hello packaged world respectively.
# Finally, if you want to remove it, you can run:
sudo apt-get remove hello-world