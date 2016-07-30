package proxy

import (
	"testing"
)

func TestStmt_DropTable(t *testing.T) {
	server := newTestServer(t)
	n := server.nodes["node1"]
	c, err := n.getMasterConn()
	if err != nil {
		t.Fatal(err)
	}
	c.UseDB("mixer")
	if _, err := c.Execute(`drop table if exists mixer_test_proxy_stmt`); err != nil {
		t.Fatal(err)
	}
	c.Close()
}

func TestStmt_CreateTable(t *testing.T) {
	str := `CREATE TABLE IF NOT EXISTS mixer_test_proxy_stmt (
          id BIGINT(64) UNSIGNED  NOT NULL,
          str VARCHAR(256),
          f DOUBLE,
          e enum("test1", "test2"),
          u tinyint unsigned,
          i tinyint,
          PRIMARY KEY (id)
        ) ENGINE=InnoDB DEFAULT CHARSET=utf8`

	server := newTestServer(t)
	n := server.nodes["node1"]
	c, err := n.getMasterConn()
	if err != nil {
		t.Fatal(err)
	}

	c.UseDB("mixer")
	defer c.Close()
	if _, err := c.Execute(str); err != nil {
		t.Fatal(err)
	}
}

func TestStmt_Insert(t *testing.T) {
	str := `insert into mixer_test_proxy_stmt (id, str, f, e, u, i) values (?, ?, ?, ?, ?, ?)`

	c := newTestDBConn(t)
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Execute(1, "a", 3.14, "test1", 255, -127); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()
}

func TestStmt_Select(t *testing.T) {
	str := `select str, f, e from mixer_test_proxy_stmt where id = ?`

	c := newTestDBConn(t)
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if result, err := s.Execute(1); err != nil {
		t.Fatal(err)
	} else {
		if len(result.Values) != 1 {
			t.Fatal(len(result.Values))
		}

		if len(result.Fields) != 3 {
			t.Fatal(len(result.Fields))
		}

		if str, _ := result.GetString(0, 0); str != "a" {
			t.Fatal("invalid str", str)
		}

		if f, _ := result.GetFloat(0, 1); f != float64(3.14) {
			t.Fatal("invalid f", f)
		}

		if e, _ := result.GetString(0, 2); e != "test1" {
			t.Fatal("invalid e", e)
		}

		if str, _ := result.GetStringByName(0, "str"); str != "a" {
			t.Fatal("invalid str", str)
		}

		if f, _ := result.GetFloatByName(0, "f"); f != float64(3.14) {
			t.Fatal("invalid f", f)
		}

		if e, _ := result.GetStringByName(0, "e"); e != "test1" {
			t.Fatal("invalid e", e)
		}

	}

	s.Close()
}

func TestStmt_NULL(t *testing.T) {
	str := `insert into mixer_test_proxy_stmt (id, str, f, e) values (?, ?, ?, ?)`

	c := newTestDBConn(t)
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Execute(2, nil, 3.14, nil); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()

	str = `select * from mixer_test_proxy_stmt where id = ?`
	s, err = c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if r, err := s.Execute(2); err != nil {
		t.Fatal(err)
	} else {
		if b, err := r.IsNullByName(0, "id"); err != nil {
			t.Fatal(err)
		} else if b == true {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "str"); err != nil {
			t.Fatal(err)
		} else if b == false {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "f"); err != nil {
			t.Fatal(err)
		} else if b == true {
			t.Fatal(b)
		}

		if b, err := r.IsNullByName(0, "e"); err != nil {
			t.Fatal(err)
		} else if b == false {
			t.Fatal(b)
		}
	}

	s.Close()
}

func TestStmt_Unsigned(t *testing.T) {
	str := `insert into mixer_test_proxy_stmt (id, u) values (?, ?)`

	c := newTestDBConn(t)
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if pkg, err := s.Execute(3, uint8(255)); err != nil {
		t.Fatal(err)
	} else {
		if pkg.AffectedRows != 1 {
			t.Fatal(pkg.AffectedRows)
		}
	}

	s.Close()

	str = `select u from mixer_test_proxy_stmt where id = ?`

	s, err = c.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if r, err := s.Execute(3); err != nil {
		t.Fatal(err)
	} else {
		if u, err := r.GetUint(0, 0); err != nil {
			t.Fatal(err)
		} else if u != uint64(255) {
			t.Fatal(u)
		}
	}

	s.Close()
}

func TestStmt_Signed(t *testing.T) {
	str := `insert into mixer_test_proxy_stmt (id, i) values (?, ?)`

	c := newTestDBConn(t)
	defer c.Close()

	s, err := c.Prepare(str)

	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(4, 127); err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(uint64(18446744073709551516), int8(-128)); err != nil {
		t.Fatal(err)
	}

	s.Close()

}

func TestStmt_Trans(t *testing.T) {
	c1 := newTestDBConn(t)
	defer c1.Close()

	if _, err := c1.Execute(`insert into mixer_test_proxy_stmt (id, str) values (1002, "abc")`); err != nil {
		t.Fatal(err)
	}

	var err error
	if err = c1.Begin(); err != nil {
		t.Fatal(err)
	}

	str := `select str from mixer_test_proxy_stmt where id = ?`

	s, err := c1.Prepare(str)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := s.Execute(1002); err != nil {
		t.Fatal(err)
	}

	if err := c1.Commit(); err != nil {
		t.Fatal(err)
	}

	if r, err := s.Execute(1002); err != nil {
		t.Fatal(err)
	} else {
		if str, _ := r.GetString(0, 0); str != `abc` {
			t.Fatal(str)
		}
	}

	if err := s.Close(); err != nil {
		t.Fatal(err)
	}
}
