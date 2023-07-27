package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
)

// GameStats represents the statistics for each game.
type GameStats struct {
	TotalKills   int            `json:"total_kills"`   // Total number of kills in the game.
	Players      []string       `json:"players"`      // List of players in the game.
	Kills        map[string]int `json:"kills"`        // Number of kills for each player.
	PlayerRank   map[int]string `json:"player_ranking"` // Ranking of players based on kills.
	KillsByMeans map[string]int `json:"kills_by_means"` // Number of kills by means (weapon).
}

// readLogFile reads the log file and returns the content as a slice of strings (lines).
func readLogFile(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

// convertToMeansOfDeath calculates the total kills by each means of death across all games.
func convertToMeansOfDeath(games map[string]GameStats) map[string]map[string]int {
	meansOfDeath := make(map[string]map[string]int)

	for _, gameStats := range games {
		for weapon, count := range gameStats.KillsByMeans {
			if meansOfDeath[weapon] == nil {
				meansOfDeath[weapon] = make(map[string]int)
			}
			meansOfDeath[weapon]["kills"] += count
		}
	}

	return meansOfDeath
}

// extractUsefulInfo extracts useful information from the log file and generates game statistics.
func extractUsefulInfo(lines []string) map[string]GameStats {
	reKill := regexp.MustCompile(`Kill: (\d+) (\d+) (\d+): (.+) killed (.+) by (.+)`)

	// Initialize game stats for each game
	games := make(map[string]GameStats)

	currentGame := "game_01" // Initialize the first game
	games[currentGame] = GameStats{
		Players:      []string{},
		Kills:        map[string]int{},
		PlayerRank:   map[int]string{},
		KillsByMeans: map[string]int{},
	}

	for _, line := range lines {
		// Check if the line contains a kill event
		if match := reKill.FindStringSubmatch(line); match != nil {
			killer := match[4]
			victim := match[5]
			weapon := match[6]

			// If the victim is not "<world>", update the kills
			if victim != "<world>" {
				// Increment the killer's kill count
				games[currentGame].Kills[killer]++

				// Decrement the victim's kill count
				games[currentGame].Kills[victim]--
			}

			// Increment the kill count by means (weapon)
			games[currentGame].KillsByMeans[weapon]++
		} else if match := regexp.MustCompile(`InitGame`).FindStringSubmatch(line); match != nil {
			// If a new game starts, create a new gameStats struct for the new game
			currentGame = fmt.Sprintf("game_%02d", len(games)+1)
			games[currentGame] = GameStats{
				Players:      []string{},
				Kills:        map[string]int{},
				PlayerRank:   map[int]string{},
				KillsByMeans: map[string]int{},
			}
		}
	}

	// Calculate the total kills and update player ranking for each game
	for key, gameStats := range games {
		// Calculate the total kills for each game
		for _, kills := range gameStats.Kills {
			gameStats.TotalKills += kills
		}

		// Filter out "<world>" from the players list
		filteredPlayers := []string{}
		for _, player := range gameStats.Players {
			if player != "<world>" {
				filteredPlayers = append(filteredPlayers, player)
			}
		}
		gameStats.Players = filteredPlayers

		// Filter out "<world>" from the kills map
		filteredKills := make(map[string]int)
		for player, kills := range gameStats.Kills {
			if player != "<world>" {
				filteredKills[player] = kills
			}
		}
		gameStats.Kills = filteredKills

		// Update player ranking after filtering
		playerRank := make([]string, len(gameStats.Kills))
		for player := range gameStats.Kills {
			gameStats.Players = append(gameStats.Players, player)
		}
		copy(playerRank, gameStats.Players)

		sort.Slice(playerRank, func(i, j int) bool {
			return gameStats.Kills[playerRank[i]] > gameStats.Kills[playerRank[j]]
		})

		for i, player := range playerRank {
			gameStats.PlayerRank[i+1] = player
		}

		games[key] = gameStats
	}

	return games
}

func main() {
	filename := "assets/logs/qgames.log"
	lines := readLogFile(filename)
	games := extractUsefulInfo(lines)

	// Convert games to JSON
	jsonData, err := json.MarshalIndent(games, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling JSON:", err)
	}

	// Print JSON data to the terminal
	fmt.Println(string(jsonData))
}
