// Package webhooks Implements methods for handling different webhooks
package webhooks

/*
	This file contains different functions to create Rich Embed objects of the class discordgo.MessageEmbed
	using the EmbedBuilder class
*/

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
)

// createScriptEmbed Use EmbedBuilder struct and it's functions to create a MessageEmbed for scripts webhook
func createScriptEmbed(json ScriptJSON) *discordgo.MessageEmbed {
	if json.Type == "backup" {
		return createBackupScriptEmbed(json)
	}
	return nil
}

func createBackupScriptEmbed(data ScriptJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed()
	details := fmt.Sprintf(
		"script: %s\nTime: %d\nDescription: %s",
		fmt.Sprintf("`%s`", data.Path),
		*data.Time,
		data.Description,
	)
	messageEmbed.AddField("Details", details).
		SetFooter(fmt.Sprintf("See full logs at %s", data.Logfile)).
		SetThumbnail(backupLogo)

	// Set title and color based on Returncode
	if *data.Returncode == 0 {
		logs := getLogs(data.Stdout)
		messageEmbed.SetTitle("Backup Complete "+checkMark).SetColor(successColor).
			AddField("Logs", logs)
	} else {
		logs := getLogs(data.Stderr)
		title := fmt.Sprintf("Backup Error %s (exit code %d)", crossMark, *data.Returncode)
		messageEmbed.SetTitle(title).SetColor(errorColor).
			AddField("Logs", logs)
	}
	return messageEmbed.MessageEmbed
}

// getLastLogLines Get the last n lines of the raw log output
func getLogs(stdout string) string {
	logs := strings.Split(stdout, "\n")
	if len(logs) > logLines {
		logs = logs[len(logs)-logLines:]
	}
	return fmt.Sprintf("`%s`", strings.Join(logs, "\n"))
}

// createSonarrGrabEmbed Create MessageEmbed when a release is grabbed by Sonarr
func createSonarrGrabEmbed(data SonarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().
		SetThumbnail(sonarrLogo).
		SetColor(infoColor)

	episode := data.Episodes[0]
	title := fmt.Sprintf(
		"%s Season %d Episode %d - Grabbed",
		data.Series.Title,
		*episode.SeasonNumber,
		*episode.EpisodeNumber,
	)
	information := fmt.Sprintf(
		"Title: %s\nQuality: %s\nSize: %s",
		episode.Title,
		data.Release.Quality,
		humanize.Bytes(*data.Release.Size),
	)
	messageEmbed.SetTitle(title).SetDescription(information)
	return messageEmbed.MessageEmbed
}

// createSonarrDownloadEmbed Create MessageEmbed when a release is downloaded by Sonarr
func createSonarrDownloadEmbed(data SonarrJSON) *discordgo.MessageEmbed {

	episode := data.Episodes[0]
	title := fmt.Sprintf(
		"%s Season %d Episode %d %s",
		data.Series.Title,
		*episode.SeasonNumber,
		*episode.EpisodeNumber,
		checkMark,
	)
	messageEmbed := NewEmbed().
		SetThumbnail(sonarrLogo).
		SetColor(successColor)

	information := fmt.Sprintf(
		"Title: %s\nQuality: %s",
		episode.Title,
		data.EpisodeFile.Quality,
	)
	messageEmbed.SetTitle(title).SetDescription(information)
	return messageEmbed.MessageEmbed
}

// createRadarrGrabEmbed Create MessageEmbed when a movie is grabbed by Radarr
func createRadarrGrabEmbed(data RadarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().
		SetThumbnail(radarrLogo).
		SetColor(infoColor)

	title := fmt.Sprintf("%s (%d) - Grabbed", data.RemoteMovie.Title, *data.RemoteMovie.Year)
	information := fmt.Sprintf(
		"Release Title: %s\nQuality: %s\nSize: %s\nIndexer: %s",
		data.Release.ReleaseTitle,
		data.Release.Quality,
		humanize.Bytes(*data.Release.Size),
		data.Release.Indexer,
	)
	messageEmbed.SetTitle(title).SetDescription(information)
	return messageEmbed.MessageEmbed
}

// createRadarrDownloadEmbed Create MessageEmbed when a movie is downloaded by Radarr
func createRadarrDownloadEmbed(data RadarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().
		SetThumbnail(radarrLogo).
		SetColor(successColor)

	title := fmt.Sprintf("%s (%d) %s", data.RemoteMovie.Title, *data.RemoteMovie.Year, checkMark)
	information := fmt.Sprintf(
		"Quality: %s\nRelease Group: %s",
		data.MovieFile.Quality,
		data.MovieFile.ReleaseGroup,
	)
	messageEmbed.SetTitle(title).SetDescription(information)
	return messageEmbed.MessageEmbed
}

// createTestEmbed Create a embed object for testing
func createTestEmbed() *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().SetAuthor("Sonarr", sonarrLogo)
	messageEmbed.SetThumbnail(sonarrLogo).
		AddField("Following episodes of Attack on Titan downloaded", "* Season 1 Episode 02 Episode Title")
	return messageEmbed.MessageEmbed
}
