package hivego

import "testing"

func TestGenerateTrxIdHiveTransaction(t *testing.T) {
	tx := getTestVoteTx()
	got, _ := tx.generateTrxId()
	expected := "12164dcee518674c586e6a61d08623c44980e326"
	if got != expected {
		t.Error("Expected", expected, "got", got)
	}
}

func TestPrepareJsonHiveTransaction(t *testing.T) {
	tx := getTestVoteTx()
	tx.prepareJson()
	got := len(tx.OperationsJs)
	expected := 1

	if got != expected {
		t.Error("Expected", expected, "got", got)
	}
}
