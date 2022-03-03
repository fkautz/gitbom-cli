# gitbom-cli

## What is GitBOM?

To quote [the GitBOM website](https://gitbom.dev/):

```
GitBOM is a minimalistic scheme for build tools to:
1. Build a compact artifact tree, tracking every source code file incorporated into each build artifact
2. Embed a unique, content addressable reference for that artifact tree, the GitBOM identifier, into the artifact at build time
```

For information, see [the website](https://gitbom.dev/) and the [list of GitBOM resources](https://gitbom.dev/resources/)

## Using gitbom-cli

### Pre-requisites

You will need go 1.18 installed on your workstation

As of this writing (March 2022),  go 1.18 is still unstable. To install, head to https://go.dev/dl/

## Set-up

```bash
$ git clone https://github.com/fkautz/gitbom-go.git
$ cd gitbom-go
$ go build -o gitbom cmd/gitbom/main.go 
```

This will give you a gitbom binary within the directory you run the build from.

## Running

To run this binary, run this command:

```
$ ./gitbom

NAME
       gitbom (v0.0.1) - Generate gitboms from files

USAGE
       gitbom artifact-tree [files]
       gitbom bom [artifact-file] [artifact-tree-files [artifact-tree files...]]

       gitbom will create a .bom/ directory in the current working
       directory and store generated gitboms in .bom/

LEGAL
       gitbom (v0.0.1) Copyright 2022 gitbom-go contributors
       SPDX-License-Identifier: Apache-2.0
```

We can use this binary to create a gitoid of any file like this:

```bash
$ ./gitbom bom [FILE NAME]
```

For example:

```bash
$ ./gitbom bom go.mod
cae3ffe42751b963dc04bb007e7bade25d7b330ccebabfeea3d44871f71e409c
```

We can also create an artifact-tree gitoid like this

```bash
$ ./gitbom artifact-tree [DIRECTORY]
```

For example:

```bash
$ ./gitbom artifact-tree .
29439dbacda23ec9221a238b5f73c4583ad6f7ad38b5e8db242ef7842dfcc73b
```

When you run these command, gitbom will create a .bom/ directory in the current working directory and store generated gitboms in .bom/