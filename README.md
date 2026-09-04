# minesweeper-go

A terminal-based Minesweeper game written in Go, using [tcell](https://github.com/gdamore/tcell) for rendering.

## Features

- Three difficulty presets (beginner, normal, expert) or a custom board size / mine count
- Keyboard-driven navigation with flood-fill reveal of empty cells
- Flag placement with automatic win detection
- Elapsed-time tracking

## Installation

### From source

```sh
go build -o minesweeper-go .
```

### Arch Linux (AUR)

```sh
yay -S minesweeper-go-git
```

## Usage

```sh
./minesweeper-go
```

At startup, type `begginer`, `normal`, or `expert` for a preset board, or press enter to configure a custom mine count and map size.

### Controls

| Key            | Action                  |
|----------------|-------------------------|
| Arrow keys     | Move cursor             |
| Space          | Reveal cell             |
| `F`            | Flag / unflag cell      |
| `R`            | Restart the game        |
| `Esc` / Ctrl+C | Quit                    |

## License

Released into the public domain under the [Unlicense](LICENSE).
