package cli

func Run() {
	loadConfigFromCli()
	loadConfigFromFile()
	setLoggingLevel()
}
