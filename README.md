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
        |  52  | Saint-Malo                       |
        +------+----------------------------------+
        |  68  | Trébeurden                       |
        +------+----------------------------------+
        |  69  | Locquirec                        |
        +------+----------------------------------+
        |  84  | Morgat                           |
        [...]

* Describe harbor:

        $ +-------------+-------------------------------------------------------------------------------------------------+
        | INFORMATION |                                              VALUE                                              |
        +-------------+-------------------------------------------------------------------------------------------------+
        | latitude    |                                                                                           44.67 |
        +-------------+-------------------------------------------------------------------------------------------------+
        | longitude   |                                                                                           -1.17 |
        +-------------+-------------------------------------------------------------------------------------------------+
        | name        | Arcachon (Jetée d'Eyrac)                                                                        |
        +-------------+-------------------------------------------------------------------------------------------------+
        | tides       | PM: 00h43 BM: 06h44 PM: 13h21 BM: 19h24 PM: 3,63m BM: 1,05m PM: 3,54m BM: 1,08m PM: 55 PM: 55   |
        |             | PM: 01h56 BM: 07h58 PM: 14h35 BM: 20h36 PM: 3,66m BM: 0,99m PM: 3,67m BM: 0,93m PM: 56 PM: 59   |
        |             | PM: 03h09 BM: 09h10 PM: 15h43 BM: 21h47 PM: 3,79m BM: 0,86m PM: 3,89m BM: 0,74m PM: 63 PM: 68   |
        |             | PM: 04h15 BM: 10h17 PM: 16h44 BM: 22h53 PM: 3,97m BM: 0,70m PM: 4,13m BM: 0,53m PM: 74 PM: 79   |
        |             | PM: 05h13 BM: 11h20 PM: 17h38 BM: 23h53 PM: 4,16m BM: 0,53m PM: 4,34m BM: 0,33m PM: 85 PM: 90   |
        |             | PM: 06h06 BM: 12h16 PM: 18h29 BM: 4,30m PM: 0,39m BM: 4,50m PM: 94 PM: 98                       |
        |             | PM: 00h47 BM: 06h56 PM: 13h06 BM: 19h18 PM: 0,20m BM: 4,38m PM: 0,31m BM: 4,56m BM: 100 BM: 101 |
        +-------------+-------------------------------------------------------------------------------------------------+


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
