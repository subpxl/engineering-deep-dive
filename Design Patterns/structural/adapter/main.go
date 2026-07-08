package main

import "fmt"

func main() {

	slack := NewSlackAdapter(&SlackClient{})
	teams := NewTeamsAdapter(&TeamsWebhook{}, "https://webhook.teams.com/abc")
	discord := NewDiscordAdapter(&DiscordBot{}, 1234545)
slack.Send("general","build successful")
teams.Send("deployment","service deployed to production")
discord.Send("alerts","cpu usage above 90%")

}

type NotificationSender interface {
	Send(recipient string, message string)
}

type SlackClient struct{}

func (s *SlackClient) PostMessage(channel string, text string, asbot bool) {
	fmt.Printf("slack _. %s: %s (bot=%t)\n", channel, text, asbot)
}

type TeamsWebhook struct{}

func (t *TeamsWebhook) SendCard(title string, body string, wehookUrl string) {
	fmt.Printf("Teams -> %s: [%s] %s\n", wehookUrl, title, body)
}

type DiscordBot struct{}

func (d *DiscordBot) SendMessage(channelId int64, content string, tts bool) {
	fmt.Printf("Discord -> channel %d: %s (tts=%t)\n", channelId, content, tts)
}

// adapters

type DiscordAdapter struct {
	discordBot *DiscordBot
	channelId  int64
}

func NewDiscordAdapter(discordBot *DiscordBot, channelId int64) *DiscordAdapter {
	return &DiscordAdapter{discordBot: discordBot, channelId: channelId}
}

func (d *DiscordAdapter) Send(recipient string, message string) {
	d.discordBot.SendMessage(d.channelId, message, false)
}

type TeamsAdapter struct {
	teamsWebhook *TeamsWebhook
	webhookUrl   string
}

func NewTeamsAdapter(tramsWebhook *TeamsWebhook, webhookurl string) *TeamsAdapter {
	return &TeamsAdapter{teamsWebhook: tramsWebhook, webhookUrl: webhookurl}
}

func (t *TeamsAdapter) Send(recipient string, message string) {
	t.teamsWebhook.SendCard(recipient, message, t.webhookUrl)
}

type SlackAdapter struct {
	slackClient *SlackClient
}

func NewSlackAdapter(slackClient *SlackClient) *SlackAdapter {
	return &SlackAdapter{slackClient: slackClient}
}

func (s *SlackAdapter) Send(recipient string, mesage string) {
	s.slackClient.PostMessage(recipient, mesage, true)
}
