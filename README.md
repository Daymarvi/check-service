# check-service
A small go program to check the Windows service state
# Description

Check Windows service status through the Windows service manager and return a Sensu check state.

| Service state                                                       | Check state  |
|---------------------------------------------------------------------|--------------|
| Running                                                             | OK (0)       |
| Start pending, stop pending, continue pending, pause pending, paused | WARNING (1)  |
| Stopped                                                             | CRITICAL (2) |
| Service not found or service manager unreachable                    | UNKNOWN (3)  |

# Synopis

```
check-service.exe --service MyService
check-service.exe --service "My Service"
```

The service name can also be given through the `CHECK_SERVICE` environment variable.

# Installation

## Requirements

- Go 1.26 or later

## Building from source 

Clone the repository

```
& git clone https://github.com/Daymarvi/check-service
& cd check-service
```

build it on Windows

```
& go build
```

or cross-compile it from Linux/macOS

```
GOOS=windows GOARCH=amd64 go build -o check-service.exe
```

## Usage

```
.\check-service.exe
Usage:
  check-service [flags]
  check-service [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -h, --help             help for check-service
  -s, --service string   Name of the Windows service to check

Use "check-service [command] --help" for more information about a command.

Error executing check-service: error validating input: --service flag or CHECK_SERVICE environment variable is required

.\check-service.exe --service fax
CRITICAL: fax stopped

.\check-service.exe --service wuauserv
CRITICAL: wuauserv stopped

.\check-service.exe --service  Winmgmt
OK: Winmgmt running
```

# Release

Releases are built and published by GitHub Actions with [GoReleaser](https://goreleaser.com) when a version tag is pushed:

```
git checkout main
git pull
git tag v1.0.0
git push origin v1.0.0
```

The release page then contains `check-service_<version>_windows_amd64.tar.gz` and `_windows_386.tar.gz` (with `bin/check-service.exe`, usable as a Sensu asset) and a sha512 checksums file.

# Todo

- Add better command line management
