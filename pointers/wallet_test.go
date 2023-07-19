package pointers

import "testing"

func TestWaller(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, 10)

	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(10)
		assertNoError(t, err)
		assertBalance(t, wallet, 10)
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(100)

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)

	})
}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("want %s got %s", want, got)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted error but didn't get one")
	}
	if want != got {
		t.Errorf("got %q, want %q", got.Error(), want)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Error("unexpected error")
	}
}
