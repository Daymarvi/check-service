# check-service
A small go program to check the Windows service state
# Description

Check Windows service status by using the syscall and return sensu.checkstate.

# Synopis

```
check-service.exe --service MyService
check-service.exe --service "My Service"
```

# Installation

## Building from source 

Clone the repository

```
& git clone https://github.com/Daymarvi/sensu-go-service-check
& cd sensu-go-service-check
```

build it

```
& go build
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
  -s, --service string   Expected service status

Use "check-service [command] --help" for more information about a command.

Error executing check-service: error validating input: --service environment variable is required

.\check-service.exe --service fax
CRITICAL: fax stopped

.\check-service.exe --service wuauserv
CRITICAL: wuauserv stopped

.\check-service.exe --service  Winmgmt
OK: Winmgmt Running.
```

# Todo

- Add better command line management
- Add more state status like "pending start"