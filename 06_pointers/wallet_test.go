package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
    assertBalance := func(t testing.TB, w Wallet, expectedBalance Bitcoin) {
        t.Helper()

		got := w.Balance()
		if got != expectedBalance {
			t.Errorf("Got %q expected %q", got, expectedBalance)
		}
    }

    assertError := func(t testing.TB, got, want error) {
        t.Helper()
        if got == nil {
            t.Fatal("Expected Error, but didn't get one")
        }

        if got != want {
            t.Errorf("got error %q, expected %q", got, want)
        }
    }

	t.Run("Should deposit correctly", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Should withdraw correctly", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
        assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Should not allow to withdraw insufficient funds", func(t *testing.T) {
        startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
        err := wallet.Withdraw(Bitcoin(20))
        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, startingBalance)
	})
}
