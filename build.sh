sudo apt-get install -y gcc dpkg-dev gpg


#### Create binary


# <package-name>_<version>-<release-number>_<architecture>
# themes_0.0.1-1_x86_64 
# themes_0.0.1-1_x86_64.deb 
mkdir -p themes_0.0.1-1_x86_64/usr/bin/
mkdir -p themes_0.0.1-1_x86_64/DEBIAN/
cp themes_0.0.1-1_x86_64/usr/bin .

echo "Package: themes
Version: 0.0.1
Maintainer: example <example@example.com>
Depends: libc6
Architecture: amd64
Homepage: http://linuxthemes.org
Description: A program that prints hello" \
> $PWD/themes_0.0.1-1_x86_64/DEBIAN/control

dpkg --build $PWD/themes_0.0.1-1_x86_64
dpkg --info $PWD/themes_0.0.1-1_x86_64.deb
dpkg --contents $PWD/themes_0.0.1-1_x86_64.deb

mkdir -p themes/apt-repo/pool/main/
cp $PWD/themes_0.0.1-1_x86_64.deb $PWD/themes/apt-repo/pool/main/ 

mkdir -p themes/apt-repo/dists/stable/main/binary-amd64/z
dpkg-scanpackages --arch amd64 $PWD/themes/apt-repo/pool/ > $PWD/themes/apt-repo/dists/stable/main/binary-amd64/Packages
cat $PWD/themes/apt-repo/dists/stable/main/binary-amd64/Packages | gzip -9 > $PWD/themes/apt-repo/dists/stable/main/binary-amd64/Packages.gz