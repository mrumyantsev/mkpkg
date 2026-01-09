# Mkpkg

[![ru](https://img.shields.io/badge/lang-ru-red.svg)](./README.ru.md)

**Mkpkg** is a lightweight code generation tool for automated Go package creation. Depending on specified parameters it creates the Go files with `structs` or `interfaces` and its methods. If the file has been created already than the tool will append only inexistent code blocks to its end.

## Installation

1. Run the commands in the Terminal.

```bash
make
sudo make install
```

## Usage

### Example 1: Creating a Simple Struct

The command:

```bash
mkpkg ./internal/config
```

Will create a Go file by path `./internal/config/config.go` with the following content:

```go
package config

type Config struct {
}

func New() *Config {
	return &Config{}
}

```

### Example 2: Creating a Customized Struct

The command:

```bash
mkpkg -m 'Load(fpath string); do(smth string) (int, error); ParseCli(args []string) error' ./internal/config
```

Will create a Go file by path `./internal/config/config.go` with the following content:

```go
package config

type Config struct {
}

func New() *Config {
	return &Config{}
}

func (c *Config) Load(fpath string) {
}

func (c *Config) ParseCli(args []string) error {
	return nil
}

func (c *Config) do(smth string) (int, error) {
	return 0, nil
}

```

### Example 3: Creating an Interface

The command:

```bash
mkpkg --iface -n IConfig -m 'Load(fpath string); do(smth string) (int, error); ParseCli(args []string) error' ./internal/config
```

Will create a Go file by path `./internal/config/config.go` with the following content:

```go
package config

type IConfig interface {
	Load(fpath string)
	ParseCli(args []string) error
	do(smth string) (int, error)
}

```

### Get Help

To show full information about flags you can provide `-h` or `--help` flag like so:

```bash
mkpkg -h
```

## System Requirements

- Windows/Linux/macOS
- Make
- Go (v1.15 or higher)

## License

[MIT License](./LICENSE)
