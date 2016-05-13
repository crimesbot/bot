package settings

import (
	"log"
	"os"
)

const (
	tokenEnv   = "CRIMESTOKEN"
	ipEnv      = "OPENSHIFT_GO_IP"
	portEnv    = "OPENSHIFT_GO_PORT"
	logFileEnv = "LOG_FILE"
)

var (
	Token, IP, Port, LogFile string
)

func init() {

	Token = os.Getenv(tokenEnv)
	if Token == "" {
		log.Panicln("TOKEN ENV NOT FOUND!")
	}

	IP = os.Getenv(ipEnv)
	if IP == "" {
		log.Panic("PORT ENV NOT FOUND!")
	}

	Port = os.Getenv(portEnv)
	if Port == "" {
		log.Panic("PORT ENV NOT FOUND!")
	}

	LogFile = os.Getenv(logFileEnv)
	if LogFile == "" {
		LogFile = "~/log.out"
	}

}
