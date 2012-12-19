package main

import (
	"code.google.com/p/goconf/conf"
	"flag"
	"fmt"
	"log"
	"os"
)

const TEST_CONF_FILE string = "kanban.conf"

var CONF_FILE = flag.String("c", "kanban.conf", "Configuration file")
var LOG_FILE = flag.String("l", "kanban.log", "Log file")
var SPORT = flag.Int("p", 9658, "Server port")
var TLS = flag.Bool("tls", false, "Activate tls")
var VERBOSE = flag.Bool("v", false, "Verbose mode")

var LOGGER *log.Logger
var LOG_FLAGS = log.LstdFlags

var (
	info_connect_bdd string // postgres://kanban:mdp@127.0.0.1:5432/kanban
	server_port      int    // 9658
	log_file         string // kanban.log
	tls_mode         bool   // false
	verbose_mode     bool   // false
)

func readConf(filename string) error {
	c, err := conf.ReadConfigFile(filename)
	if err != nil {
		return err
	}
	var err_psql, err_serv_port, err_log_file, err_tls_mode, err_verbose_mode error
	info_connect_bdd, err_psql = c.GetString("default", "db-uri")
	server_port, err_serv_port = c.GetInt("default", "server-port")
	log_file, err_log_file = c.GetString("default", "log-file")
	tls_mode, err_tls_mode = c.GetBool("default", "tls")
	verbose_mode, err_verbose_mode = c.GetBool("default", "verbose")

	// default value if not in the config file
	if err_psql != nil {
		info_connect_bdd = "postgres://kanban:mdp@127.0.0.1:5432/kanban"
	}
	if err_serv_port != nil {
		server_port = 9658
	}
	if err_log_file != nil {
		log_file = "kanban.log"
	}
	if err_tls_mode != nil {
		tls_mode = false
	}
	if err_verbose_mode != nil {
		verbose_mode = false
	}
	return nil
}

func main() {
	flag.Parse()
	if err := readConf(*CONF_FILE); err != nil {
		fmt.Println("Error with the configuration file:", err)
		return
	}
	f, err := os.OpenFile(*LOG_FILE, os.O_CREATE | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Impossible to open log file")
		return
	}
	defer f.Close()
	LOGGER = log.New(f, "", LOG_FLAGS)
	LOGGER.Print("toto")
	// startServer()
}
