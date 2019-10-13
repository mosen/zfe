package hosts

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"testing"
)

func TestNewRepository(t *testing.T) {
	db, err := sqlx.Open("postgres", "user=zabbix dbname=zabbix password=zabbix")
	if err != nil {
		t.Errorf("Cannot open database connection: %s", err)
	}
	defer func() {
		err := db.Close()
		fmt.Println(err)
	}()

	repo := NewRepository(db)
	if repo == nil {
		t.Failed()
	}
}

func TestRepository_Find(t *testing.T) {
	db, err := sqlx.Open("postgres", "user=zabbix dbname=zabbix password=zabbix sslmode=disable")
	if err != nil {
		t.Errorf("Cannot open database connection: %s", err)
	}
	defer func() {
		err := db.Close()
		fmt.Println(err)
	}()

	repo := NewRepository(db)

	hosts, err := repo.Find()
	if err != nil {
		t.Errorf("While executing query: %s", err)
	}

	if len(hosts) == 0 {
		t.Fail()
	} else {
		t.Logf("Got %d Host(s) from Host Repository", len(hosts))
	}
}
