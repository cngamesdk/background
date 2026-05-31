package chat_monitor

type ServiceGroup struct {
	GameService
	ChatService
	SensitiveService
	WhitelistService
	BanService
	StatsService
}
