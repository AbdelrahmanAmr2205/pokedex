# Pokedex CLI

A command-line Pokedex application written in Go. This project is built as part of the backend learning path on [Boot.dev](https://boot.dev). It interacts with the [PokeAPI](https://pokeapi.co/) to navigate the Pokemon world, explore locations, and catch Pokemon directly from your terminal!

## Features

* **Interactive REPL:** A clean Read-Eval-Print Loop for continuous interaction.
* **API Integration:** Fetches real-time data from the PokeAPI.
* **In-Memory Caching:** Implements a custom concurrency-safe caching system with a background reap loop to minimize redundant network requests.
* **Pagination:** Seamlessly navigate forward and backward through Pokemon location areas.
* **Catching Mechanics:** Probability-based catching system that scales with the Pokemon's base experience.

## Prerequisites

* [Go](https://golang.org/doc/install) 1.26 or higher

## Installation

1. Clone the repository:
    ```bash
   git clone [https://github.com/AbdelrahmanAmr2205/pokedex.git](https://github.com/AbdelrahmanAmr2205/pokedex.git)
   cd pokedex
   ```

2. Build the executable:
   ```bash
   go build -o pokedex
   ```

3. Run the application:
   ```bash
   ./pokedex
   ```

## Usage & Commands

Once you start the application, you will be greeted by the `Pokedex >` prompt. You can use the following commands:

* `help`: Displays a help message with all available commands.
* `map`: Displays the next 20 location areas in the Pokemon world.
* `mapb`: Displays the previous 20 location areas.
* `explore <location_name>`: Lists all the Pokemon that can be found in a specific location area.
* `catch <pokemon_name>`: Throws a Pokeball to attempt to catch a Pokemon. Catch success relies on the Pokemon's base experience.
* `inspect <pokemon_name>`: View the name, height, weight, stats, and types of a Pokemon you have already caught.
* `pokedex`: Lists all the Pokemon you have successfully caught and added to your Pokedex.
* `exit`: Safely closes the application.

## Project Structure

* `main.go` & `repl.go`: CLI entrypoint and command processing loop.
* `command_*.go`: Handlers for each individual CLI command.
* `internal/pokeapi/`: Handles all HTTP requests to the PokeAPI, including JSON parsing and response structuring.
* `internal/pokecache/`: A custom caching package using `sync.RWMutex` to ensure thread-safe read/write operations and automatic background expiration of stale data.

## Author

* **Abdelrahman Amr**
