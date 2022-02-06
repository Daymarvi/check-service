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

Clone the repository, then

```
& git clone https://github.com/Daymarvi/sensu-go-service-check
& cd sensu-go-service-check
```
To build it

```
& go build
```

## Usage

```
.\check-service.exe
Usage:
  sensu-go-service-check [flags]
  sensu-go-service-check [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -h, --help             help for sensu-go-service-check
  -s, --service string   Expected service status

Use "sensu-go-service-check [command] --help" for more information about a command.

Error executing sensu-go-service-check: error validating input: --service environment variable is required
.\check-service.exe --service fa
Usage:
  sensu-go-service-check [flags]
  sensu-go-service-check [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -h, --help             help for sensu-go-service-check
  -s, --service string   Expected service status

Use "sensu-go-service-check [command] --help" for more information about a command.

Error executing sensu-go-service-check: error executing check: Could not access service: The specified service does not exist as an installed service.
.\check-service.exe --service fax
CRITICAL: fax stopped
.\check-service.exe --service wuauserv
CRITICAL: wuauserv stopped
.\check-service.exe --service  Winmgmt
OK: Winmgmt Running.
```
