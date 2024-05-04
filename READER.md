# Telnet Client

A simple command-line Telnet client written in Go (Golang). This tool allows you to connect to Telnet servers, send commands, and view responses directly from your terminal.

## Features

- Connects to a specified IP address and port.
- Validates IP address and port range (1-65535) before attempting connection.
- Supports concurrent reading and writing to the Telnet server.
- Gracefully handles disconnections and input errors.
- Exits upon receiving a Ctrl+C command.

## Usage

1. **Prerequisites**: Ensure you have [Go](https://golang.org/doc/install) installed on your system.

2. **Building the Application**:
    - Clone this repository to your local machine.
   ```bash
   git clone https://github.com/wuchuhengtools/telnet.git
   ```
    - Navigate to the project directory.
   ```bash
   cd YOUR_REPOSITORY
   ```
    - Build the Telnet client binary.
   ```bash
   go build -o telnet-client main.go
   ```

3. **Running the Telnet Client**:
   Open a terminal and run the compiled binary with the target IP address and port number.
   ```bash
   ./telnet-client <address> <port>
   ```
   Example:
   ```bash
   ./telnet-client 192.168.1.1 23
   ```

## Command Line Arguments

- `<address>`: The IP address of the Telnet server (e.g., `192.168.1.1`). Must be a valid IPv4 address.
- `<port>`: The port number of the Telnet service (e.g., `23`). Must be within the range of 1 to 65535.

## Development

Contributions are welcome! If you find any bugs or have feature requests, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.