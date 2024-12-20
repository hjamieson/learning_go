package main

import (
	"fmt"
	"io"
	"sort"
)

type Team struct {
	name    string
	players []string
}

type League struct {
	Teams []Team
	Wins  map[string]int
}

type Record struct {
	team string
	wins int
}

func (l *League) MatchResult(t1 string, score1 int, t2 string, score2 int) {
	if score1 > score2 {
		l.Wins[t1]++
	}
	if score2 > score1 {
		l.Wins[t2]++
	}
}

func NewLeague() League {

	return League{
		Teams: []Team{
			{name: "Tampa", players: []string{"Bull", "Bob", "Bjorn"}},
			{name: "Orlando", players: []string{"Oscar", "Owen", "Omar"}},
			{name: "Miami", players: []string{"Manny", "Moe", "Mike"}},
		},
		Wins: map[string]int{"Tampa": 0, "Orlando": 0, "Miami": 0},
	}
}

func (l League) Ranking() []Record {
	var ranking []Record

	for k, v := range l.Wins {
		ranking = append(ranking, Record{k, v,})
	}
	sort.Slice(ranking, func(i, j int) bool{
		return ranking[i].wins > ranking[j].wins
	})

	return ranking
}

func game(l League, t1 string, score1 int, t2 string, score2 int) {
	l.MatchResult(t1, score1, t2, score2)
	fmt.Println(l.Ranking())
}

type Ranker interface {
	Ranking()[]string
}
func RankPrinter(r Ranker, w io.Writer) {
	
}

func main() {
	league := NewLeague()
	ranking := league.Ranking()
	fmt.Println(ranking)
	game(league, "Tampa", 3, "Orlando", 2)
	game(league,"Orlando", 0, "Miami", 2)
	game(league, "Orlando", 1, "Tampa", 2)
	game(league, "Orlando", 1, "Tampa", 0)
	game(league, "Orlando", 1, "Miami", 0)
	game(league, "Orlando", 1, "Miami", 0)

}
