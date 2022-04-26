package cli

import (
	"flag"
	"strings"

	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

type index struct {
	Root string
	Type string
}

type config struct {
	Indexes  []index
	Loglevel string
}

var (
	configPathPtr = flag.String("config", "/etc/gopile/config.toml", "Path to configuration file")
	logLevelPtr   = flag.String("loglevel", "INFO", "Logging level")
	helpPtr       = flag.Bool("help", false, "Display help message and quit")
)

func loadConfigFromCli() {
	flag.Parse()

	if *helpPtr {
		printHelpMsg()
		return
	}
}

func loadConfigFromFile() {
	var conf config
	_, err := toml.DecodeFile(*configPathPtr, &conf)
	if err != nil {

	}
}

func setLoggingLevel() {
	switch loglevel := strings.ToLower(*logLevelPtr); loglevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
		log.SetReportCaller(true)
	default:
		log.SetLevel(log.InfoLevel)
	}
}
