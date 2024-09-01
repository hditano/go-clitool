# Go-CLITool

![Go](https://img.shields.io/badge/Go-1.16%2B-blue.svg) ![License](https://img.shields.io/badge/license-MIT-green.svg)

## Description

`clitool` is a command-line interface (CLI) tool written in Go that provides various functionalities organized into "palettes." Currently, it includes commands for performing network-related tasks and retrieving disk usage information.

## Features

- **Network Palette**:
  - `ping`: Sends a ping request to a specified URL and returns the response status code.

- **Information Palette**:
  - `diskusage`: Displays information about disk usage, including total, free, and used space in gigabytes.

## Installation

To install `clitool`, follow these steps:

1. **Clone the repository**:

   ```bash
   git clone https://github.com/hditano/go-clitool.git
   cd go-clitool
   ```

2. **Build the project**:

   Make sure you have Go installed on your system. Then, run:

   ```bash
   go build -o clitool
   ```

3. **Run the tool**:

   ```bash
   ./clitool
   ```

## Usage

Once you have installed `clitool`, you can use it from the command line. Here are some examples of how to use the available commands:

### Network Commands

- **Ping**:

  ```bash
  clitool net ping --url <your-url>
  ```

  Replace `<your-url>` with the URL you want to ping.

### Information Commands

- **Disk Usage**:

  ```bash
  clitool info diskusage
  ```

## TODO

- [ ] Add more information commands (e.g., CPU usage, memory usage).
- [ ] Create unit tests for all functionalities.
- [ ] Add support for additional output formats (e.g., JSON, XML).
- [ ] Implement a help command to list all available commands and their usage.
- [ ] Implement Azure/Aws Storage solutions
- [ ] Implement VMs/EC2 Solution usage

## Modules Used

This project utilizes the following Go modules:

- [Cobra](https://github.com/spf13/cobra): For creating the CLI.
- [golang.org/x/sys/unix](https://pkg.go.dev/golang.org/x/sys/unix): For retrieving system information, such as disk usage.

## Contributions

Contributions are welcome! If you would like to contribute, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Contact

If you have any questions or comments, feel free to reach out to me at [hditano@gmail.com](mailto:hditano@gmail.com).

