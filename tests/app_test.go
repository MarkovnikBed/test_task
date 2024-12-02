package tests

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"medods/internal/repository"
)

func TestConnectionToDB(t *testing.T) {
	if userName == "" {
		userName = "postgres"
	}
	if password == "" {
		password = "12345"
	}
	if dbname == "" {
		dbname = "avecoder"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}

	str := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", userName, password, dbname, host, port)
	db, err := sql.Open("postgres", str)
	require.NoError(t, err)
	defer db.Close()

	err = db.Ping()
	require.NoError(t, err)
}

func TestInsertRT(t *testing.T) {
	if userName == "" {
		userName = "postgres"
	}
	if password == "" {
		password = "12345"
	}
	if dbname == "" {
		dbname = "avecoder"
	}
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	err := os.Setenv("USERNAME_MEDODS", userName)
	if err != nil {
		require.NoError(t, err)
	}
	err = os.Setenv("PASSWORD_MEDODS", password)
	if err != nil {
		require.NoError(t, err)
	}
	err = os.Setenv("DB_MEDODS", dbname)
	if err != nil {
		require.NoError(t, err)
	}
	err = os.Setenv("HOST_MEDODS", host)
	if err != nil {
		require.NoError(t, err)
	}
	err = os.Setenv("PORT_MEDODS", port)
	if err != nil {
		require.NoError(t, err)
	}

	testRep := repository.CreateRepository()
	defer testRep.DB.Close()
	err = testRep.PrepareTable()
	require.NoError(t, err)
	err = testRep.InsertRT("testToken", "wrongID")
	assert.Error(t, err)
	err = testRep.InsertRT("testToken", "123434134")
	assert.NoError(t, err)
	_, err = testRep.DB.Exec("DELETE FROM refresh_tokens WHERE id=123434134")
	require.NoError(t, err)

}
