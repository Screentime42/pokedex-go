# pokedex-go 🐾

A simple Pokédex CLI written in Go.  
This project lets you explore Pokémon data directly from your terminal, with an interactive REPL (read–eval–print loop) for smooth command handling.

---

## Features
- Interactive REPL interface for exploring Pokémon data
- Built entirely in Go with idiomatic patterns
- Modular design for easy extension
- Lightweight and fast — no external dependencies beyond Go’s standard library

---

## Installation
Clone the repository and build the binary:

    git clone https://github.com/Screentime42/pokedex-go.git
    cd pokedex-go
    go build -o pokedex

Run the CLI:

    ./pokedex

---

## Usage
Once inside the REPL, you can type commands to interact with the Pokédex.

Example commands:

    help            # Show available commands
    exit            # Quit the REPL
    catch <name>    # Attempt to catch a Pokémon by name
    inspect <name>  # View details of a caught Pokémon
    map             # Explore the Pokémon world map

(Commands may evolve as the project grows — check main.go and repl_test.go for the latest behavior.)

---

## Project Structure
- main.go → Entry point for the CLI
- pokedex-go/ → Core logic and REPL implementation
- repl_test.go → Tests for REPL functionality
- go.mod → Module definition

---

## Development
Run tests with:

    go test ./...

Format and tidy dependencies:

    go fmt ./...
    go mod tidy
