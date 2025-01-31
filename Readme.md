# PingMon

## Overview
**PingMon** is a command-line utility that continuously pings a specified host at a given interval and plays an alert sound when the host is up or down. This is useful for network monitoring and system administration.

## Features
- Supports Linux, macOS, and Windows
- Customizable ping interval
- Sound alerts when the host is up or down
- Simple command-line interface

## Prerequisites
- Go 1.16 or later installed
- `aplay` for Linux or `afplay` for macOS (or PowerShell for Windows)

## Installation
Clone the repository and navigate to the project directory:
```sh
 git clone <repository_url>
 cd pingmon
```
Initialize the Go module:
```sh
go mod init github.com/yourusername/pingmon
go mod tidy
```

## Usage
Run the tool with the required arguments:
```sh
go run main.go --host <hostname> --time <seconds> --alert <up/down>
```

### Example:
Ping Google every 3 seconds and play a sound when it is up:
```sh
go run main.go --host google.com --time 3 --alert up
```

Ping a local server every 5 seconds and play a sound when it is down:
```sh
go run main.go --host 192.168.1.1 --time 5 --alert down
```

## Building an Executable
To create a standalone binary:
```sh
go build -o pingmon
```
Run it:
```sh
./pingmon --host google.com --time 3 --alert up
```

## License
This project is open-source and available under the MIT License.

## Contributing
Feel free to submit issues or pull requests to improve this tool!

