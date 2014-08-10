package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"time"
)

type attachmentMessage struct {
	Fallback string            `json:"fallback"`
	Text     string            `json:"text"`
	Pretext  string            `json:"pretext"`
	Color    string            `json:"color"`
	Fields   []attachmentField `json:"fields"`
}

type attachmentField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func main() {
	for {
		for _, account := range config.Accounts {
			latestID, err := getLatestMatch(account.DotaID)
			if err != nil {
				log.Println("getLatestMatch() =>", err)
				continue
			} else if latestID <= account.LastMatchID {
				continue
			}

			result, err := formatMatchResult(account.DotaID, latestID, account.SlackName)
			if err != nil {
				log.Println("formatMatchResult() =>", err)
				continue
			}

			err = announceMatchResult(result)
			if err != nil {
				log.Println("announceMatchResult() =>", err)
				continue
			}

			account.LastMatchID = latestID
			saveConfig()
			break
		}

		time.Sleep(20 * time.Second)
	}
}

func getLatestMatch(accountID uint64) (uint64, error) {
	var history dotaMatchHistory

	url := fmt.Sprintf(
		"http://api.steampowered.com/IDOTA2Match_570/GetMatchHistory/v1?key=%s&account_id=%d",
		config.SteamAPIKey,
		accountID,
	)

	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&history)
	if err != nil {
		return 0, err
	}

	if history.Result.Status != 1 {
		return 0, errors.New(history.Result.StatusMsg)
	}

	return history.Result.Matches[0].ID, nil
}

func formatMatchResult(accountID uint64, matchID uint64, slackName string) (*attachmentMessage, error) {
	var player dotaMatchDetailsPlayer
	var details dotaMatchDetails

	url := fmt.Sprintf(
		"http://api.steampowered.com/IDOTA2Match_570/GetMatchDetails/v1?key=%s&match_id=%d",
		config.SteamAPIKey,
		matchID,
	)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&details)
	if err != nil {
		return nil, err
	}

	for _, p := range details.Result.Players {
		if p.AccountID == accountID {
			player = p
			break
		}
	}

	onDire := player.Slot&0x80 != 0
	wonGame := (details.Result.RadiantWin && !onDire) || (!details.Result.RadiantWin && onDire)
	kdaRatio := float64(player.Kills+player.Assists) / math.Max(1.0, float64(player.Deaths))
	resultString := ""
	sideNameLower := ""
	sideNameUpper := ""

	if wonGame {
		resultString = "won"
	} else {
		resultString = "lost"
	}

	if onDire {
		sideNameLower = "dire"
		sideNameUpper = "Dire"
	} else {
		sideNameLower = "radiant"
		sideNameUpper = "Radiant"
	}

	text := fmt.Sprintf(
		"%s just %s a game playing as %s on the %s side with a KDA ratio of %.2f\n",
		slackName,
		resultString,
		dotaHeroNames[player.HeroID],
		sideNameLower,
		kdaRatio,
	)

	fields := []attachmentField{
		{
			Title: "ID",
			Value: fmt.Sprintf("<http://dotabuff.com/matches/%d|%d>", matchID, matchID),
			Short: true,
		},
		{
			Title: "Length",
			Value: (time.Duration(details.Result.Duration) * time.Second).String(),
			Short: true,
		},
		{
			Title: "Game Mode",
			Value: dotaModeNames[details.Result.GameMode],
			Short: true,
		},
		{
			Title: "Lobby Type",
			Value: dotaLobbyTypes[details.Result.LobbyType],
			Short: true,
		},
		{
			Title: "Hero",
			Value: dotaHeroNames[player.HeroID],
			Short: true,
		},
		{
			Title: "Side",
			Value: sideNameUpper,
			Short: true,
		},
		{
			Title: "Kills/Deaths/Assists",
			Value: fmt.Sprintf("%d/%d/%d", player.Kills, player.Deaths, player.Assists),
			Short: true,
		},
		{
			Title: "GPM/XPM",
			Value: fmt.Sprintf("%d/%d", player.GPM, player.XPM),
			Short: true,
		},
	}

	message := &attachmentMessage{
		Fallback: text,
		Pretext:  text,
		Fields:   fields,
	}

	if wonGame {
		message.Color = "good"
	} else {
		message.Color = "danger"
	}

	return message, nil
}

func announceMatchResult(message *attachmentMessage) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(
		config.SlackAPIURL,
		"application/json",
		bytes.NewBuffer(data),
	)
	if err != nil {
		return err
	}

	resp.Body.Close()
	return nil
}
