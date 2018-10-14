// Package webhooks Implements methods for handling different webhooks
package webhooks

/*
	This file contains different functions to create Rich Embed objects of the class discordgo.MessageEmbed
	using the EmbedBuilder class
*/

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"github.com/dustin/go-humanize"
)

// createScriptEmbed Use EmbedBuilder struct and it's functions to create a MessageEmbed for scripts webhook
func createScriptEmbed(json ScriptJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed()
	if json.Type == "backup" {
		messageEmbed.SetAuthor("Backup").SetThumbnail(backupLogo)
	}
	// Set title and color  based on Returncode
	if *json.Returncode == 0 {
		messageEmbed.SetTitle("Success: Executed successfully").SetColor(successColor)
	} else {
		title := fmt.Sprintf("Error: Exited with returncode %d", *json.Returncode)
		messageEmbed.SetTitle(title).SetColor(errorColor)
	}

	// Add fields for script path and execution time
	messageEmbed.AddField("Path", "Script is located at "+json.Path).
		AddField("Description", json.Description).
		AddField("Execution Time", fmt.Sprintf("Script took %d seconds to execute", *json.Time)).
		SetFooter(fmt.Sprintf("Please see the logs at %s for more details", json.Logfile))
	return messageEmbed.MessageEmbed
}

// createSonarrGrabEmbed Create MessageEmbed when a release is grabbed by Sonarr
func createSonarrGrabEmbed(json SonarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().
		SetThumbnail(sonarrLogo).
		SetColor(infoColor).
		SetTitle("Episode Grabbed :white_check_mark:")

	episode := json.Episodes[0]
	title := fmt.Sprintf(
		"%s Season %d Episode %d",
		json.Series.Title,
		*episode.EpisodeNumber,
		*episode.SeasonNumber,
	)
	information := fmt.Sprintf(
		"* Title: %s\n* Quality: %s\n* Size: %s",
		episode.Title,
		json.Release.Quality,
		humanize.Bytes(uint64(*json.Release.Size)),
	)
	messageEmbed.AddField(title, information)
	return messageEmbed.MessageEmbed
}

// createSonarrDownloadEmbed Create MessageEmbed when a release is downloaded by Sonarr
func createSonarrDownloadEmbed(json SonarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().
		SetThumbnail(sonarrLogo).
		SetColor(successColor).
		SetTitle("Episode Downloaded :white_check_mark:")

	episode := json.Episodes[0]
	title := fmt.Sprintf(
		"%s Season %d Episode %d",
		json.Series.Title,
		*episode.EpisodeNumber,
		*episode.SeasonNumber,
	)
	information := fmt.Sprintf(
		"* Title: %s\n* Quality: %s\n* Path: %s",
		episode.Title,
		json.EpisodeFile.Quality,
		json.EpisodeFile.RelativePath,
	)
	messageEmbed.AddField(title, information)
	return messageEmbed.MessageEmbed
}

// createRadarrGrabEmbed Create MessageEmbed when a movie is grabbed by Radarr
func createRadarrGrabEmbed(json RadarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().
		SetThumbnail(radarrLogo).
		SetColor(infoColor).
		SetTitle("Movie grabbed successfully :white_check_mark:")

	title := fmt.Sprintf("%s (%d)", json.RemoteMovie.Title, *json.RemoteMovie.Year)
	information := fmt.Sprintf(
		"* Release Title: %s\n* Quality: %s\n* Size: %s\n* Indexer: %s",
		json.Release.ReleaseTitle,
		json.Release.Quality,
		humanize.Bytes(uint64(*json.Release.Size)),
		json.Release.Indexer,
	)
	messageEmbed.AddField(title, information)
	return messageEmbed.MessageEmbed
}

// createRadarrDownloadEmbed Create MessageEmbed when a movie is downloaded by Radarr
func createRadarrDownloadEmbed(json RadarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().
		SetThumbnail(radarrLogo).
		SetColor(successColor).
		SetTitle("Movie downloaded :white_check_mark:")

	title := fmt.Sprintf("%s (%d)", json.RemoteMovie.Title, *json.RemoteMovie.Year)
	information := fmt.Sprintf(
		"* Path: %s\n* Quality: %s\n* Release Group: %s",
		json.MovieFile.Path,
		json.MovieFile.Quality,
		json.MovieFile.ReleaseGroup,
	)
	messageEmbed.AddField(title, information)
	return messageEmbed.MessageEmbed
}

// createTestEmbed Create a embed object for testing
func createTestEmbed() *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().SetAuthor("Sonarr", sonarrLogo)
	messageEmbed.SetThumbnail(sonarrLogo).
		AddField("Following episodes of Attack on Titan downloaded", "* Season 1 Episode 02 Episode Title")
	return messageEmbed.MessageEmbed
}
