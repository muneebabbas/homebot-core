// Package webhooks Implements methods for handling different webhooks
package webhooks

/*
	This file contains different functions to create Rich Embed objects of the class discordgo.MessageEmbed
	using the EmbedBuilder class
*/

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

// createScriptEmbed Use EmbedBuilder struct and it's functions to create a MessageEmbed for scripts webhook
func createScriptEmbed(json ScriptJSON) *discordgo.MessageEmbed {
	embed := NewEmbed().SetTimestamp()
	// Set title and color  based on Returncode
	if *json.Returncode == 0 {
		embed.SetTitle("Success: Executed successfully").SetColor(successColor)
	} else {
		title := fmt.Sprintf("Error: Exited with returncode %d", *json.Returncode)
		embed.SetTitle(title).SetColor(errorColor)
	}

	// Add fields for script path and execution time
	embed.AddField("Path", "Script is located at "+json.Path).
		AddField("Description", json.Description).
		AddField("Execution Time", fmt.Sprintf("Script took %d seconds to execute", *json.Time)).
		SetFooter(fmt.Sprintf("Please see the logs at %s for more details", json.Logfile))
	return embed.MessageEmbed
}

// createSonarrGrabEmbed Create MessageEmbed when a release is grabbed by Sonarr
func createSonarrGrabEmbed(json SonarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().SetTimestamp()
	title := fmt.Sprintf(
		"%s Season %d Episode %d grabbed by Sonarr",
		json.Series.Title,
		*json.Episodes[0].EpisodeNumber,
		*json.Episodes[0].SeasonNumber,
	)
	messageEmbed.SetTitle(title).SetColor(infoColor)
	return messageEmbed.MessageEmbed
}

// createSonarrDownloadEmbed Create MessageEmbed when a release is downloaded by Sonarr
func createSonarrDownloadEmbed(json SonarrJSON) *discordgo.MessageEmbed {
	messageEmbed := NewEmbed().SetTimestamp()
	title := fmt.Sprintf(
		"%s Season %d Episode %d downloaded by Sonarr",
		json.Series.Title,
		*json.Episodes[0].EpisodeNumber,
		*json.Episodes[0].SeasonNumber,
	)
	messageEmbed.SetTitle(title).SetColor(successColor)
	return messageEmbed.MessageEmbed
}
