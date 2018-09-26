package tests

import (
	"net/http"
	"testing"

	"github.com/lib/pq"
)

var client = &http.Client{}

// TODO...
func TestDb(t *testing.T) {
	db, err := pq.Open("localhost")

	if err != nil || db.Begin != nil {
		t.Error("Postgres is not available")
	}
}

// TODO .....
func TestGetAll(t *testing.T) {
	req, err := http.NewRequest("GET", "http://localhost:9010/lumos/asset/asset/v1/get-list", nil)

	if err != nil {
		t.Fail()
	}

	resp, err := client.Do(req)

	if err != nil || resp.StatusCode != http.StatusOK {
		t.Fail()
	}
}
