package main

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {

		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.assertBalance(t, Bitcoin(10))

	})
	t.Run("withdraw", func(t *testing.T) {
		assertNoError := func(t testing.TB, got error) {
			t.Helper()
			if got != nil {
				t.Fatal("got an error but didn't one")
			}
		}
		wallet := Wallet{10}
		err := wallet.Withdraw(Bitcoin(5))
		wallet.assertBalance(t, Bitcoin(5))
		assertNoError(t, err)

	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		assertError := func(t testing.TB, got error, want error) {
			t.Helper()
			if got == nil {
				t.Fatal("didn't get an error but wanted one")
			}
			if got != want {
				t.Errorf("got %q, want %q", got, want)
			}
		}
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		wallet.assertBalance(t, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func (w *Wallet) assertBalance(t testing.TB, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}
