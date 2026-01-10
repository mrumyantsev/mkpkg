# Mkpkg

[![en](https://img.shields.io/badge/lang-en-blue.svg)](./README.md)

**Mkpkg** &mdash; это легковесный генератор кода для автоматизированного создания Go-пакетов. В зависимости от указанных параметров он создает Go-файлы со структурами (`struct`) или интерфейсами (`interface`) и их методами. Если файл уже был создан, то генератор добавит только несуществующие блоки кода в конец.

## Установка

1. Клонируйте репозиторий на свой компьютер.

```bash
git clone https://github.com/mrumyantsev/mkpkg.git
```

2. Запустите команды в терминале.

```bash
make
sudo make install
```

## Применение

### Пример 1: Создание простой структуры

Данная команда:

```bash
mkpkg ./internal/config
```

Создаст Go-файл по пути `./internal/config/config.go` со следующим содержимым:

```go
package config

type Config struct {
}

func New() *Config {
	return &Config{}
}

```

### Пример 2: Создание кастомизированной структуры

Данная команда:

```bash
mkpkg -m 'Load(fpath string); do(smth string) (int, error); ParseCli(args []string) error' ./internal/config
```

Создаст Go-файл по пути `./internal/config/config.go` со следующим содержимым:

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

### Пример 3: Создание интерфейса

Данная команда:

```bash
mkpkg --iface -n IConfig -m 'Load(fpath string); do(smth string) (int, error); ParseCli(args []string) error' ./internal/config
```

Создаст Go-файл по пути `./internal/config/config.go` со следующим содержимым:

```go
package config

type IConfig interface {
	Load(fpath string)
	ParseCli(args []string) error
	do(smth string) (int, error)
}

```

### Получить помощь

Чтобы получить полную информацию о флагах вы можете добавить флаг `-h` или `--help`, как в примере:

```bash
mkpkg -h
```

## Системные требования

- Windows/Linux/macOS
- Make
- Go (версии 1.15 или выше)

## Лицензия

[MIT License](./LICENSE)
