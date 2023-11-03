package error_handling

import (
	"testing"
)

func TestWallet_Withdraw(t *testing.T) {
	t.Run("Should return InsufficientFundsError when there is no balance", func(t *testing.T) {
		w := &Wallet{balance: 10}
		err := w.Withdraw(15)
		AssertError(t, err, InsufficientFundsError)
		AssertBalance(t, w, 10)
	})
	t.Run("Should successfuly withdraw when there are sufficient funds", func(t *testing.T) {
		w := &Wallet{balance: 10}
		err := w.Withdraw(10)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		AssertBalance(t, w, 0)
	})
	t.Run("Should successfuly deposit and withdraw", func(t *testing.T) {
		w := &Wallet{balance: 10}
		err := w.Withdraw(10)
		AssertNoError(t, err)
	})
}

func AssertNoError(t testing.TB, actual error) {
	t.Helper()
	if actual != nil {
		t.Errorf("Expected nil error, got an actual error")
	}
}

func AssertError(t testing.TB, expected error, actual error) {
	t.Helper()
	if actual == nil {
		t.Fatalf("Expected error, got none")
	}
	if expected != actual {
		t.Errorf("Wrong error returned: expected %s, got %s", expected, actual)
	}
}

func AssertBalance(t testing.TB, w *Wallet, expected BitCoin) {
	actual := w.Balance()
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
