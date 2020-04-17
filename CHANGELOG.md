# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

Upcoming release

### Added

- `make release`: Use ghr to publish binaries to GitHub
- `make tools`: Install build and release tooling (gox and ghr)
- `make debug`: Display generated vars used in the Makefile
- Commit vendored modules to the repo
- `.go-version` file set to 1.14.0
- Include sha256 checksum file for binaries
- `make release-draft`: Publish a draft release to GitHub

### Changed

- `make build_all`: Use `go mod tidy` and `go mod vendor`
- Determine version from Git describe output
- Built binaries now include the version string

### Removed

- Removes LDFLAG for main.buildTime to support reproducible builds

## [0.1.0]

Initial release with bash-completion and API commands. Supports configuration
using CLI flags, environment variables, or YAML/JSON files.

### Added

- Initial project boilerplate (docs, license, contributing, etc)
- Bash completion is via the `bash-completion` command

Added the following commands that interact with the Signal Sciences API:

- get-corp
- list-agents
- list-corps
- list-suspicious-ips
- list-top-attacks

<!--
This section should be updated with every release. It contains a sequence of
links to GitHub that show the full Git diff between each release. The brackets
allow us to render the version headers as links by adding brackets to any
matching headers. Any commits that don't yet belong to a Git tag will show the
Git diff from the last tag to the master branch HEAD.
-->
[Unreleased]: https://github.com/timoguin/sigcli/compare/v0.1.0..HEAD
[0.1.0]: https://github.com/timoguin/sigcli/tree/v0.1.0
