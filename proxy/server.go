package proxy

import (
	"github.com/HYmian/jhs/config"
	"github.com/HYmian/jhs/router"
	"net"
	"runtime"
	"strings"

	"github.com/siddontang/go-log/log"
)

type MysqlConfig map[string]string

type Server struct {
	cfg *config.Config

	addr     string
	user     string
	password string
	database string
	default_ *router.Rule

	running bool

	listener net.Listener

	nodes map[string]*Node

	rules map[string]*router.Rule

	mysql_config MysqlConfig
}

func (s *Server) SetMysqlConfig(mysql_config MysqlConfig) {
	s.mysql_config = mysql_config
}

func (s *Server) GetMySqlVariable(name string) (v string, ok bool) {
	v, ok = s.mysql_config[name]
	return
}

func (s *Server) GetRule(table string) *router.Rule {
	if r, ok := s.rules[table]; ok {
		return r
	} else {
		return s.default_
	}
}

func NewServer(cfg *config.Config) (*Server, error) {
	s := new(Server)

	s.cfg = cfg

	s.addr = cfg.Addr
	s.user = cfg.User
	s.password = cfg.Password
	s.database = cfg.Database
	s.default_ = router.NewDefaultRule(cfg.Default)

	if err := s.parseNodes(); err != nil {
		return nil, err
	}

	if err := s.parseRules(); err != nil {
		return nil, err
	}

	var err error
	netProto := "tcp"
	if strings.Contains(netProto, "/") {
		netProto = "unix"
	}
	s.listener, err = net.Listen(netProto, s.addr)

	if err != nil {
		return nil, err
	}

	log.Info("Server run MySql Protocol Listen(%s) at [%s]", netProto, s.addr)
	return s, nil
}

func (s *Server) Run() error {
	s.running = true

	for s.running {
		conn, err := s.listener.Accept()
		if err != nil {
			log.Error("accept error %s", err.Error())
			continue
		}

		go s.onConn(conn)
	}

	return nil
}

func (s *Server) Close() {
	s.running = false
	if s.listener != nil {
		s.listener.Close()
	}
}

func (s *Server) onConn(c net.Conn) {
	conn := s.newConn(c)

	defer func() {
		if err := recover(); err != nil {
			const size = 4096
			buf := make([]byte, size)
			buf = buf[:runtime.Stack(buf, false)]
			log.Error("onConn panic %v: %v\n%s", c.RemoteAddr().String(), err, buf)
		}

		conn.Close()
	}()

	if err := conn.Handshake(); err != nil {
		log.Error("handshake error %s", err.Error())
		c.Close()
		return
	}

	conn.Run()
}
