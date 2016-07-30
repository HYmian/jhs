package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMySqlConfig(t *testing.T) {
	cfg, err := ParseMySqlConfigFile("E:\\WorkSpace\\bin\\mysql.yaml")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(cfg)
}

func TestConfig(t *testing.T) {
	var testConfigData = []byte(
		`
addr : 127.0.0.1:4000
user : root
password :
database: jhs
#默认节点，所有未在下面列出的表都将派发到这个节点，如果分表的话应该去除
default: node2

#监控地址
pprof_addr: 192.168.0.152:6060
log_level : error

nodes :
-
  name : node1
  down_after_noalive : 300
  idle_conns : 16
  rw_split: true
  user: root
  password:
  master : 127.0.0.1:3306
  slave : 127.0.0.1:4306
-
  name : node2
  user: root
  master : 127.0.0.1:3307
-
  name : node3
  down_after_noalive : 300
  idle_conns : 16
  rw_split: false
  user: root
  password:
  master : 127.0.0.1:3308

rules:
  -
    table: mixer_test_shard_hash
    key: id
    nodes: [node1, node2, node3]
    type: hash
    shard_num: 2
  -
    table: mixer_test_shard_range
    key: id
    type: range
    nodes: [node2, node3]
    range: -10000-
`)

	cfg, err := ParseConfigData(testConfigData)
	if err != nil {
		t.Fatal(err)
	}

	if len(cfg.Nodes) != 3 {
		t.Fatal(len(cfg.Nodes))
	}

	if len(cfg.ShardRule) != 2 {
		t.Fatal(len(cfg.ShardRule))
	}

	testNode := NodeConfig{
		Name:             "node1",
		DownAfterNoAlive: 300,
		IdleConns:        16,
		RWSplit:          true,

		User:     "root",
		Password: "",

		Master: "127.0.0.1:3306",
		Slave:  "127.0.0.1:4306",
	}

	if !reflect.DeepEqual(cfg.Nodes[0], testNode) {
		fmt.Printf("%v\n", cfg.Nodes[0])
		t.Fatal("node1 must equal")
	}

	testNode_2 := NodeConfig{
		Name:   "node2",
		User:   "root",
		Master: "127.0.0.1:3307",
	}

	if !reflect.DeepEqual(cfg.Nodes[1], testNode_2) {
		t.Fatal("node2 must equal")
	}

	testShard_1 := ShardConfig{
		Table:    "mixer_test_shard_hash",
		Key:      "id",
		Nodes:    []string{"node1", "node2", "node3"},
		ShardNum: 2,
		Type:     "hash",
	}
	if !reflect.DeepEqual(cfg.ShardRule[0], testShard_1) {
		t.Fatal("ShardConfig0 must equal")
	}

	testShard_2 := ShardConfig{
		Table: "mixer_test_shard_range",
		Key:   "id",
		Nodes: []string{"node2", "node3"},
		Type:  "range",
		Range: "-10000-",
	}
	if !reflect.DeepEqual(cfg.ShardRule[1], testShard_2) {
		t.Fatal("ShardConfig1 must equal")
	}

	if cfg.LogLevel != "error" || cfg.User != "root" || cfg.Password != "" || cfg.Addr != "127.0.0.1:4000" {
		t.Fatal("Top Config not equal.")
	}
}
