# Poseidon

[![License Apache 2][badge-license]](LICENSE)
[![GitHub version](https://badge.fury.io/gh/nlamirault%2Fposeidon.svg)](https://badge.fury.io/gh/nlamirault%2Fposeidon)

* Master : [![Circle CI](https://circleci.com/gh/nlamirault/poseidon/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/poseidon/tree/master)
* Develop : [![Circle CI](https://circleci.com/gh/nlamirault/poseidon/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/poseidon/tree/develop)

This tool is a tool for tides.


## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/poseidon-0.1.0_netbsd_arm) ]


## Usage

* CLI help:

        $ poseidon help

* List harbors :

        $ poseidon harbor list
        +------+----------------------------------+
        |  ID  |               NAME               |
        +------+----------------------------------+
        | /52  | Saint-Malo                       |
        +------+----------------------------------+
        | /68  | Trébeurden                       |
        +------+----------------------------------+
        | /69  | Locquirec                        |
        +------+----------------------------------+
        | /84  | Morgat                           |
        [...]
        +------+----------------------------------+



## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>

[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat
