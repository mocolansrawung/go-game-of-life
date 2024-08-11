# Game of Life in Go

This is a Go implementation of Conway's Game of Life. It includes functionality for creating and evolving a grid of cells according to the rules of the game.

## Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)
- [License](#license)

## Overview

Conway's Game of Life is a cellular automaton devised by mathematician John Conway. The game consists of a grid of cells that evolve through generations based on the following rules:
- **Any live cell with fewer than two live neighbours dies as if caused by under-population.**
- **Any live cell with two or three live neighbours lives on to the next generation.**
- **Any live cell with more than three live neighbours dies, as if by over-population.**
- **Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.**

## Features

- Create and initialize a grid.
- Add patterns such as the Glider Gun to the grid.
- Evolve the grid through multiple generations.
- Display the grid in the console with live and dead cells.

## Installation

To run this project, you need to have Go installed on your system. You can download and install Go from the [official Go website](https://golang.org/dl/).

1. Clone the repository:

    ```bash
    git clone https://github.com/mocolansrawung/go-game-of-life.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-game-of-life-go
    ```

3. Build the project:

    ```bash
    go build
    ```

## Usage

To run the Game of Life simulation:

1. Navigate to the project directory.

2. Execute the compiled binary:

    ```bash
    ./go-game-of-life
    ```

3. The program will display the grid and evolve it over multiple generations in the console.

### Example

When running the program, you should see a grid of cells where `⬛` represents live cells and `⬜` represents dead cells. The grid will update to show the next generation of cells.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
