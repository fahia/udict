# udict
Urban Dictionary from your command line.

# Usage
Say you were curious about what the word `kk` meant, so you head over to your terminal and type the following:
`udict define -word kk`
![udict-pic](https://user-images.githubusercontent.com/22937651/44891183-f1bed280-acab-11e8-9ed1-03a52092afb5.png)

Just like that, you can stay up to date with the latest, bleeding-edge phrases all the young kids are using these days.

# Installation

### Via Go

`go get github.com/fahia/udict`

### Binary

#### Download
```shell
$ curl -fSL "https://github.com/fahia/udict/releases/download/0.1.0/udict" -o "/usr/local/bin/udict" \
	&& chmod a+x "/usr/local/bin/udict"
```
#### Run it!
```shell
$ udict define -word any-word-you-want
```
# Contributing

When contributing to this repository, please first discuss the change you wish to make via a Github issue.
