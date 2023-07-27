# Quake Log Parser

This project is a Quake log parser implemented in Go, developed specifically for solving [this assignment](https://gist.github.com/cloudwalk-tests/704a555a0fe475ae0284ad9088e203f1). It allows you to read a Quake log file, extract game data for each match, collect kill data, and generate reports based on the parsed information.

## Requirements

To use this project, you need the following:

- Go programming language
- Git

## Getting Started

Follow these steps to get started with the Quake game log parser:

1. Clone the repository:

   ```bash
   git clone https://github.com/Meshram007/quakeGame.git
   ```

2. Run the project:

   ```bash
   go run main.go
   ```

3. Output the project:

   ```bash
   "game_22": {
   "total_kills": 0,
   "players": [
      "Isgalamido",
      "Zeh",
      "Oootsimo",
      "Mal",
      "Assasinu Credi",
      "Dono da Bola"
   ],
   "kills": {
      "Assasinu Credi": -8,
      "Dono da Bola": -3,
      "Isgalamido": 0,
      "Mal": -18,
      "Oootsimo": 6,
      "Zeh": 6
   },
   "player_ranking": {
      "1": "Zeh",
      "2": "Oootsimo",
      "3": "Isgalamido",
      "4": "Dono da Bola",
      "5": "Assasinu Credi",
      "6": "Mal"
   },
   "kills_by_means": {
      "MOD_FALLING": 3,
      "MOD_MACHINEGUN": 4,
      "MOD_RAILGUN": 9,
      "MOD_ROCKET": 37,
      "MOD_ROCKET_SPLASH": 60,
      "MOD_SHOTGUN": 4,
      "MOD_TRIGGER_HURT": 14
    }
   }
   ```
