package config

import (
	"fmt"
	"io/ioutil"

	"github.com/siddontang/go-yaml/yaml"
)

type NodeConfig struct {
	Name             string `yaml:"name"`
	DownAfterNoAlive int    `yaml:"down_after_noalive"`
	IdleConns        int    `yaml:"idle_conns"`
	RWSplit          bool   `yaml:"rw_split"`

	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DB       string `yaml:"db"`

	Master string `yaml:"master"`
	Slave  string `yaml:"slave"`
}

type ShardConfig struct {
	Table string   `yaml:"table"`
	Key   string   `yaml:"key"`
	Nodes []string `yaml:"nodes"`
	Type  string   `yaml:"type"`
	Range string   `yaml:"range"`
	//by mian
	ShardNum int `yaml:"shard_num"`
}

type Config struct {
	Addr     string `yaml:"addr"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Default  string `yaml:"default"`

	Pprof_addr string `yaml:"pprof_addr"`
	LogLevel   string `yaml:"log_level"`

	Nodes []NodeConfig `yaml:"nodes"`

	ShardRule []ShardConfig `yaml:"rules"`
}

func ParseConfigData(data []byte) (*Config, error) {
	var cfg Config
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	}

	if !includeNode(cfg.Nodes, cfg.Default) {
		return nil, fmt.Errorf("default node[%s] not in the nodes list.", cfg.Default)
	}

	for _, rule := range cfg.ShardRule {
		if len(rule.Nodes) < 1 {
			return nil, fmt.Errorf("rule table[%s] node number less than 1", rule.Table)
		}
		for _, node := range rule.Nodes {
			if !includeNode(cfg.Nodes, node) {
				return nil, fmt.Errorf("rule table[%s] node[%s] not in the nodes list.", rule.Table, node)
			}
		}
	}

	return &cfg, nil
}

func ParseJHSConfigFile(fileName string) (*Config, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	return ParseConfigData(data)
}

func ParseMySqlConfigFile(file_name string) (map[string]string, error) {
	data, err := ioutil.ReadFile(file_name)
	if err != nil {
		return nil, err
	}

	cfg := make(map[string]string)
	if err := yaml.Unmarshal([]byte(data), &cfg); err != nil {
		return nil, err
	} else {
		return cfg, err
	}
}

func includeNode(nodes []NodeConfig, node string) bool {
	for _, n := range nodes {
		if n.Name == node {
			return true
		}
	}
	return false
}
