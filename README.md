# Themes

![Build Status](https://github.com/linux-themes/themes/actions/workflows/themes.yml/badge.svg)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/linux-themes/themes)
![Repo Size](https://img.shields.io/github/repo-size/linux-themes/themes)
![GitHub Release](https://img.shields.io/github/v/release/linux-themes/themes)
![GitHub Tag](https://img.shields.io/github/v/tag/linux-themes/themes)



Manage your linux desktop themes

## Install

```sh
go install github.com/linux-themes/themes@latest
```

Install go from https://go.dev/doc/install

Or

Install go from apt

```sh
sudo apt install nala -y
sudo nala install golang -y
```

Check go bin is on path
```sh
echo $PATH | tr ':' '\n' | grep 'go/bin'
```

If not add to .bashrc 
```sh
echo -n 'export PATH="~/go/bin:$PATH"' >> ~/.bashrc
```
```sh
source ~/.bashrc
```

## Use

```sh
themes list all
themes install https://github.com/sudo-adduser-jordan/mint-y-winx/raw/main/mint-y-winx.tar.xz
themes remove 
```

## Contributing

Feel free to dive in! [Open an issue](https://github.com/RichardLitt/standard-readme/issues/new) or submit PRs.

```sh
git clone https://github.com/linux-themes/themes  
```


### Contributors

This project exists thanks to all the people who contribute. 
<!-- <a href="https://github.com/RichardLitt/standard-readme/graphs/contributors"><img src="https://opencollective.com/standard-readme/contributors.svg?width=890&button=false" /></a> -->
