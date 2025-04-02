package main

import (
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
			{"Name": "Pepper", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := []Player{
			{"Pepper", 10},
			{"Chris", 33},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)
	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		got := store.GetPlayerScore("Chris")
		want := 33

		assertScoreEqual(t, got, want)
	})

	t.Run("store wins for existing player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Chris")

		got := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEqual(t, got, want)
	})

	t.Run("store wins for new player", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
		{"Name": "Cleo", "Wins": 10},
		{"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)

		store.RecordWin("Pepper")

		got := store.GetPlayerScore("Pepper")
		want := 1

		assertScoreEqual(t, got, want)
	})
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")

	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile
}

func assertScoreEqual(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}
