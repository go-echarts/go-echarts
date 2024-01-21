<p align="center">
	<img src="https://user-images.githubusercontent.com/19553554/52535979-c0d0e680-2d8f-11e9-85c8-2e9f659e7c6f.png" width=300 height=300 />
</p>

<h1 align="center">go-echarts</h1>
<p align="center">
    <em>🎨 The adorable charts library for Golang.</em>
</p>

<p align="center">
	<a href="https://github.com/go-echarts/go-echarts/pulls">
        <img src="https://img.shields.io/badge/contributions-welcome-brightgreen.svg?style=flat" alt="Contributions welcome">
    </a>
    <a href="https://opensource.org/licenses/MIT">
        <img src="https://img.shields.io/badge/License-MIT-brightgreen.svg" alt="MIT License">
    </a>
    <a href="https://pkg.go.dev/github.com/go-echarts/go-echarts/v2">
        <img src="https://godoc.org/github.com/go-echarts/go-echarts?status.svg" alt="GoDoc">
    </a>
    <a href="https://echarts.apache.org/">
        <img src="https://img.shields.io/badge/echarts-_v5.4.3-orange" alt="echartsVersion">
    </a>
</p>

# 🔰 Installation

## gomod

```shell
$ go get -u github.com/go-echarts/go-echarts/v2/...
```

OR

```shell
# go.mod

require github.com/go-echarts/go-echarts/v2
```

## classic to get go-echarts

```shell
# this may be a crude way to use v2 go-echarts without gomod(GO111MODULE=off)

$ go get -u github.com/go-echarts/go-echarts/...
$ cd $go-echarts-project
$ mkdir v2 && mv charts components datasets opts render templates types v2
```

# ⏳ Version Control

The go-echarts project is being developed under v2 version and the active codebase is on the master branch.

v1 and v2 are incompatible which means that you cannot upgrade go-echarts from v1 to v2 smoothly. But I think it is
worth trying that new version.

Especially, when there contains mino changes (usually in enhancement), we will release the `rc` version before a
standard
release.
So, if you upgrade your projects cross the `rc` versions, maybe
need little adjust.

# ✨ Features

* Clean and comprehensive API.
* Visualize your data in 25+ different ways.
* Highly configurable chart options.
* Detailed documentation and a rich collection of examples.
* Visualize your geographical data with 400+ maps.

# 💡 Contribution

go-echarts is an open source project and built on the top of other open-source projects.
Welcome all the kinds of contributions. No matter it is for typo fix, bug fix or big new features.
Please do not hesitate to ask a question or send a pull request.

We strongly value documentation and integration with other projects, so we are very glad to accept improvements for
these aspects.

