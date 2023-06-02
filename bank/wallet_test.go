package bank

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	assertError := func(t testing.TB, got, want error) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Deposit(Bitcoin(20))
		want := Bitcoin(30)
		assertBalance(t, wallet, want)
	})
	t.Run("withdrawal", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(30))
		err := wallet.Withdrawal(Bitcoin(20))
		want := Bitcoin(10)
		assertError(t, err, nil)
		assertBalance(t, wallet, want)
	})

	t.Run("withdrawal insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(30)
		wallet := Wallet{startingBalance}
		err := wallet.Withdrawal(Bitcoin(100))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, startingBalance)
	})
}

func TestBitcoin(t *testing.T) {
	got := Bitcoin(10).String()
	want := "10 BTC"
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
