package parser

import (
	"os"
	"bufio"
	"github.com/parserlog/model"
)

func Parser(bc *model.BodyCount) error {

	// init kills
	bc.Kills = make(map[string]int)

	// open file
	file, err := os.Open(bc.FilePath)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		// get line
		line := scanner.Text()

		// match kill
		if match_kill, _ := kill.MatchString(line); match_kill {

			bc.TotalKills++

			// match killer
			match_killer, _ := killer.FindStringMatch(line)

			// killer or killed?
			if match_killer.String() == "<world>" {

				// match killed
				match_killed, _ := killed.FindStringMatch(line)

				// decrease kills
				bc.Kills[match_killed.String()]--

			} else {

				// increase kills
				bc.Kills[match_killer.String()]++

			}

		}

	}

	// set players
	for player := range bc.Kills {
		bc.Players = append(bc.Players, player)
	}

	return err

}
