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
date = "2020-12-11"
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

If you develop your Go CLI tool without a full-blown CI/CD pipeline, `goreleaser` is almost as easy to use as a command-line tool. In the following sections, I present a short walkthrough of setting up and using `goreleaser` for publishing a binary on GitHub and for Homebrew (as an example of an OS package manager. `goreleaser` provides a couple more to choose from).


## Using goreleaser on the command line

Setting up `goreleaser` consists of a couple of simple steps.

1. Install `goreleaser`
2. Run `goreleaser init`
3. Edit the configuration file and:
    * Add or modify prebuild hooks
    * Add a GitHub/GitLab/Gitea token

There are a couple of optional steps available, depending on which features you want to use. Here, we focus on two of these.

1. How to release to Homebrew
2. How to sign your release

Finally, we will add a cool bonus feature: The ability to have your binary print the current version, without having to manually maintain a version string in your code.


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
# This is an example goreleaser.yaml file with some sane defaults.
# Make sure to check the documentation at http://goreleaser.com
before:
  hooks:
    # You may remove this if you don't use go modules.
    - go mod download
    # you may remove this if you don't need go generate
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

* "before" hooks that run prior to building the binary
* build settings (defaulting to disabling CGO and compiling to linux, windows and darwin)
* an archives target that produces tarred and gzipped binaries
* a checksum target that generates a checksum file
* a snapshot setting, if you want to create a test release "between" two versions
* and a changelog target for the release page on the Git host.

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
