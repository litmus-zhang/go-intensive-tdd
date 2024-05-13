package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t testing.TB, got error, want string) {
		t.Helper()
		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("want %q, want %q", got, want)
		}

	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))

		err := wallet.Withdraw(Bitcoin(10))
		fmt.Println("err", err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBal := Bitcoin(20)
		wallet := Wallet{startingBal}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBal)
		assertError(t, err, "insufficient funds")
	})
}
