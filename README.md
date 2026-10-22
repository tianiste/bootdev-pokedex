# Pokédex CLI

A simple Go command-line Pokédex built as a Boot.dev learning project using the PokeAPI.

## Run

With Go installed, run from the project root:

```sh
go run .
```

## Commands

| Command | What it does |
| --- | --- |
| `help` | Show available commands. |
| `map` | Show the next 20 location areas. |
| `mapb` | Show the previous page of location areas. |
| `explore <location>` | List Pokémon found in a location area. |
| `catch <pokemon>` | Attempt to catch a Pokémon. |
| `inspect <pokemon>` | Show a caught Pokémon's height, weight, stats, and types. |
| `pokedex` | List caught Pokémon. |
| `exit` | Close the CLI. |

Location responses are cached in memory, with a mutex protecting concurrent access and periodic cleanup of old entries. Catching uses a random chance based on base experience. Caught Pokémon are kept in memory for the current session.

An internet connection is needed for API requests.
