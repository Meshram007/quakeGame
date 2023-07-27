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


