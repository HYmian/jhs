package main

import (
	"flag"
	"fmt"
	"github.com/HYmian/jhs/config"
	"github.com/HYmian/jhs/proxy"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"
	"strings"
	"syscall"

	"github.com/siddontang/go-log/log"
)

var (
	jhs_config   *string = flag.String("config", "./jhs.yaml", "jhs proxy config file")
	mysql_config *string = flag.String("mysql_config", "./mysql.yaml", "mysql config file")
	logLevel     *string = flag.String("log-level", "", "log level [debug|info|warn|error], default error")
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	flag.Parse()

	if len(*jhs_config) == 0 {
		fmt.Println("must use a jhs config file")
		return
	}

	jhs_cfg, err := config.ParseJHSConfigFile(*jhs_config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	if *logLevel != "" {
		setLogLevel(*logLevel)
	} else {
		setLogLevel(jhs_cfg.LogLevel)
	}

	if len(*mysql_config) == 0 {
		fmt.Println("must use a mysql config file")
	}

	mysql_cfg, err := config.ParseMySqlConfigFile(*mysql_config)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	var svr *proxy.Server
	svr, err = proxy.NewServer(jhs_cfg)
	if err != nil {
		log.Error(err.Error())
		return
	}
	svr.SetMysqlConfig(mysql_cfg)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc,
		//syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT)

	go func() {
		sig := <-sc
		log.Info("Got signal [%d] to exit.", sig)
		svr.Close()
	}()

	go func() {
		log.Error(http.ListenAndServe(jhs_cfg.Pprof_addr, nil).Error())
	}()

	svr.Run()
}

func setLogLevel(level string) {
	switch strings.ToLower(level) {
	case "debug":
		log.SetLevel(log.LevelDebug)
	case "info":
		log.SetLevel(log.LevelInfo)
	case "warn":
		log.SetLevel(log.LevelWarn)
	case "error":
		log.SetLevel(log.LevelError)
	default:
		log.SetLevel(log.LevelError)
	}
}
