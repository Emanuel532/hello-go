package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit money", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)
		assertAmount(t, got, want)
	})
	t.Run("withdraw money", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(6))
		assertNoError(t, err)

		got := wallet.Balance()
		want := Bitcoin(4)
		assertAmount(t, got, want)
	})
	t.Run("insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(1)
		wallet := Wallet{balance: Bitcoin(initialBalance)}

		err := wallet.Withdraw(Bitcoin(7))

		got := wallet.Balance()
		assertAmount(t, got, initialBalance)
		assertError(t, err, "insufficient funds")
	})
}

func assertAmount(t testing.TB, got, want Bitcoin) {
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}

}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatal("didn't expect any error but got one")
	}
}

func assertError(t testing.TB, err error, want string) {
	t.Helper()
	if err == nil {
		t.Fatal("expected an error but didn't get one")
	}
	if err.Error() != want {
		t.Errorf("got %q expected %s", err.Error(), want)
	}
}
