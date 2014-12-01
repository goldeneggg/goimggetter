goimggetter [![Build Status](https://drone.io/github.com/goldeneggg/goimggetter/status.png)](https://drone.io/github.com/goldeneggg/goimggetter/latest) [![GoDoc](https://godoc.org/github.com/goldeneggg/goimggetter?status.png)](https://godoc.org/github.com/goldeneggg/goimggetter) [![MIT License](http://img.shields.io/badge/license-MIT-lightgrey.svg)](https://github.com/goldeneggg/goimggetter/blob/master/LICENSE)
===========

__goimggetter__ is image download tool by Go.


## Features
* You can download images from website using search query keyword.
* Download target directory is `$PWD/imgs/` (not need to mkdir $PWD/imgs)
* If search result was already downloaded, it will be skipped.


## Install

```
% go get github.com/goldeneggg/goimggetter
```

## Binary download link
* [linux_amd64](https://drone.io/github.com/goldeneggg/goimggetter/files/artifacts/bin/linux_amd64/goimggetter)
* [linux_386](https://drone.io/github.com/goldeneggg/goimggetter/files/artifacts/bin/linux_386/goimggetter)
* [darwin_amd64](https://drone.io/github.com/goldeneggg/goimggetter/files/artifacts/bin/darwin_amd64/goimggetter)
* [darwin_386](https://drone.io/github.com/goldeneggg/goimggetter/files/artifacts/bin/darwin_386/goimggetter)
* [windows_amd64](https://drone.io/github.com/goldeneggg/goimggetter/files/artifacts/bin/windows_amd64/goimggetter.exe)
* [windows_386](https://drone.io/github.com/goldeneggg/goimggetter/files/artifacts/bin/windows_386/goimggetter.exe)


## Usage

```
Usage:
  goimggetter -s <SITE> [OTHER OPTIONS] <QUERY>

Application Options:
  -s, --site=        Target site (ex. flickr, and more...)
  -o, --offset=      Offset index of search result *default = 1
  -l, --limitpage=   Limit page of search result *default = 1
  -c, --concurrency= Concurrency count for save img *default = 1
  -d, --debug        Debug detail information
  -v, --version      Print version

Help Options:
  -h, --help     Show this help message
```

* Get from flickr using search query "tokyo"

```
% goimggetter -s flickr tokyo

% goimggetter -s flickr -l 3 tokyo  # page limit is 3

% goimggetter -s flickr -c 3 tokyo  # concurrency of saving img is 3
```

* Get from flickr using multibyte search query "東京"

```
% goimggetter -s flickr %E6%9D%B1%E4%BA%AC  # query need to url-encode utf8
```


## Target Sites
* flickr : using `-s flickr` option


## ChangeLog
[CHANGELOG](CHANGELOG) file for details.


## License

[LICENSE](LICENSE) file for details.

## Author
[goldeneggg](https://github.com/goldeneggg)
