/*
<!--
Copyright (c) 2019 Christoph Berger. Some rights reserved.

Use of the text in this file is governed by a Creative Commons Attribution Non-Commercial
Share-Alike License that can be found in the LICENSE.txt file.

Use of the code in this file is governed by a BSD 3-clause license that can be found
in the LICENSE.txt file.

The source code contained in this file may import third-party source code
whose licenses are provided in the respective license files.
-->

<!--
NOTE: The comments in this file are NOT godoc compliant. This is not an oversight.

Comments and code in this file are used for describing and explaining a particular topic to the reader. While this file is a syntactically valid Go source file, its main purpose is to get converted into a blog article. The comments were created for learning and not for code documentation.
-->

+++
title = "CLI tools FTW (or: how to release your CLI tools with goreleaser)"
description = "How to use goreleaser for publishing Go CLI tools beyond go get"
author = "Christoph Berger"
email = "chris@appliedgo.net"
date = "2021-04-23"
draft = "true"
categories = ["Go Ecosystem"]
tags = ["deployment", "packaging", "distribution"]
articletypes = ["Tutorial"]
+++

"go get" is a super-simple way of installing Go binaries, but not everyone has a Go compiler installed. If you want to make your CLI tools and apps available to the world, have a look at goreleaser.

<!--more-->

When I browse through Go repos on GitHub, more often than not I see this installation instruction:

```sh
go get github.com/owner/repo
```

Granted, `go get` is intriguingly simple to use, and if your CLI tool is targeted towards Go developers, it is absolutely ok to use built-in tools for sharing your work.

If your CLI tool is not restricted (or even related) to Go development, you will want to enable your target audience to get your tools the way they are familiar with.


## Meet goreleaser

Luckily, you do not need to become a package manager expert. There is a tool for this: `goreleaser`. Mostly known for use inside CI/CD pipelines, it also does a great service if you want to publish binaries for manual installation.

CI/CD pipelines usually include `goreleaser` in their workflows. In this context, using `goreleaser` is super simple. Just push a Git tag that consists of a (semantic!) version number, and the CI/CD workflow does the rest.

If you develop your Go CLI tool without a full-blown CI/CD pipeline, `goreleaser` is almost as easy to use as a command-line tool. In the following sections, I present a short walkthrough of setting up and using `goreleaser` for publishing a binary on GitHub and for Homebrew as an example of an OS package manager. (`goreleaser` provides a couple more to choose from).

As a sample project, I use `goman`, the "missing man pages" tool for Go that turns the repo README into an ad-hoc help page. (See [this post](https://appliedgo.net/goman/) for an intro to `goman`.)


## Using goreleaser on the command line

Setting up `goreleaser` consists of a couple of simple steps.

1. Install `goreleaser`
2. Run `goreleaser init`
3. Edit the configuration file and:
    * Add or modify prebuild hooks
    * Add a GitHub/GitLab/Gitea token

There are a couple of optional steps available, depending on which features you want to use. Here, I'll focus on two of these.

1. How to release to Homebrew
2. How to sign your release

Finally, I will show a fancy bonus feature: How to have your binary print the current version, without having to manually maintain a version string in your code.


### Installing goreleaser

Not surprisingly, `goreleaser` is published using `goreleaser`, so for installing you can pick your favorite installation method. Head over to the [Install - GoReleaser](https://goreleaser.com/install/) document and follow the steps for the desired installation method.


### Initial setup

After having installed `goreleaser`, cd into your Go CLI project and run

```sh
goreleaser init
```

to generate a default config file. The configuration uses YAML syntax. If you are not familiar with YAML, you will want to keep the [YAML reference]([YAML Ain’t Markup Language (YAML™) Version 1.2](https://yaml.org/spec/1.2/spec.html)) at hand. Don't let the large TOC scare you off. The spec comes with clear examples for each and every data type representation.

The default file looks like this:

```yaml
before:
  hooks:
    - go mod download
    - go generate ./...
builds:
  - env:
      - CGO_ENABLED=0
    goos:
      - linux
      - windows
      - darwin
archives:
  - replacements:
      darwin: Darwin
      linux: Linux
      windows: Windows
      386: i386
      amd64: x86_64
checksum:
  name_template: 'checksums.txt'
snapshot:
  name_template: "{{ .Tag }}-next"
changelog:
  sort: asc
  filters:
    exclude:
      - '^docs:'
      - '^test:'
```

The file already contains a couple of build actions and release targets.

* The "before" hooks declare commans to run prior to building the binary
* The "builds" settings describe environment variables to use and target operating systems to compile for. (The provided defaults disable CGO and compile to linux, windows and darwin)
* The "archives" target produces tarred and gzipped binaries. The "replacement" settings beneath controls how the resulting archive files are named.
* The "checksum" target generates (surprise, surprise) a checksum file
* The "snapshot" setting allows creating a test release "between" two versions
* The "changelog" target generates a change log for the release page on the Git host.


### A few basic customizations

The first change I am going to apply is a modification of the before hooks. `goman` needs no `go generate`, and I also want to tidy my `go.mod` file at this occasion.

```yaml
before:
  hooks:
    - go mod tidy
    - go mod download
```

Hooks are executed in the order of appearance, so no surprises here.

This setup already allows a first test. `goreleaser` can run in a test mode, without publishing anything.

Call

```sh
goreleaser --snapshot --skip-publish
```

The command creates a `dist` directory, compiles binaries for all targets provided, and tars an gzips them into archives, based on the `archives` section in the above config.

If you run this step yourself, you may notice a message saying, "gomod.proxy is disabled". This is indeed always true for snapshot builds, but is also the default for release builds. So my next change to the config file is to enable the Go proxy, in order to ensure verifiable builds. There is a special `gomod` option available for this:

```yaml
gomod:
  proxy: true
```

Also, for Windows, the archive format should be "zip". A simple format override inside the "archives" section does the trick.

And while we are at it, let's also add the README and LICENSE files, and wrap it all into a single directory inside the archive.

```yaml
archives:
  -
    replacements:
      ...
    format_overrides:
    - goos: windows
      format: zip
    files:
      - README.md
      - LICENSE.txt
    wrap_in_directory: true

```

As a result, the generated archive contains the binary and the additional files inside a subdirectory with the archiv file's base name:

![archive contents (png)](gomantarzip.png)




## The code
*/

// ## Imports and globals
package main

/*
## How to get and run the code

Step 1: `go get` the code. Note the `-d` flag that prevents auto-installing
the binary into `$GOPATH/bin`.

    go get -d github.com/appliedgo/TODO:

Step 2: `cd` to the source code directory.

    cd $GOPATH/src/github.com/appliedgo/TODO:

Step 3. Run the binary.

    go run TODO:.go


## Odds and ends
## Some remarks
## Tips
## Links


**Happy coding!**

*/
