package poker_test

import (
	"fmt"
	"io"
	"os"
	"poker"
	"testing"
	"time"
)

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(5, os.Stdout)

		cases := []poker.ScheduledAlert{
			generateAlert(0*time.Second, 100),
			generateAlert(10*time.Minute, 200),
			generateAlert(20*time.Minute, 300),
			generateAlert(30*time.Minute, 400),
			generateAlert(40*time.Minute, 500),
			generateAlert(50*time.Minute, 600),
			generateAlert(60*time.Minute, 800),
			generateAlert(70*time.Minute, 1000),
			generateAlert(80*time.Minute, 2000),
			generateAlert(90*time.Minute, 4000),
			generateAlert(100*time.Minute, 8000),
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

	t.Run("schedules alerts on game start for 7 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7)

		cases := []poker.ScheduledAlert{
			generateAlert(0*time.Second, 100),
			generateAlert(12*time.Minute, 200),
			generateAlert(24*time.Minute, 300),
			generateAlert(36*time.Minute, 400),
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})

}

func TestGame_Finish(t *testing.T) {
	store := &poker.StubPlayerStore{}
	game := poker.NewTexasHoldem(dummyBlindAlerter, store)
	winner := "Ruth"

	game.Finish(winner)
	poker.AssertPlayerWin(t, store, winner)
}

func generateAlert(at time.Duration, amount int) poker.ScheduledAlert {
	return struct {
		At     time.Duration
		Amount int
		To     io.Writer
	}{
		At:     at,
		Amount: amount,
		To:     os.Stdout,
	}
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, blindAlerter *poker.SpyBlindAlerter) {
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {

			if len(blindAlerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, blindAlerter.Alerts)
			}

			got := blindAlerter.Alerts[i]
			assertScheduledAlert(t, got, want)
		})
	}
}
