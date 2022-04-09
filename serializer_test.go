package hivego

import (
    "bytes"
    "testing"
)

func TestOpIdB(t *testing.T) {
    got := opIdB("custom_json")
    expected := byte(18)

    if got != expected {
        t.Error("Expected", expected, "got")
    }
}

func TestRefBlockNumB(t *testing.T) {
    got := refBlockNumB(36029)
    expected := []byte{189, 140}

    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}

func TestRefBlockPrefixB(t *testing.T) {
    got := refBlockPrefixB(1164960351)
    expected := []byte{95, 226, 111, 69}

    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}

func TestExpTimeB(t *testing.T) {
    got, _ := expTimeB("2016-08-08T12:24:17")
    expected := []byte{241, 121, 168, 87}

    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}

func TestCountOpsB(t *testing.T) {
    got := countOpsB(getTwoTestOps())
    expected := []byte{2}

    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}

//func TestExtensionsB

func TestAppendVString(t *testing.T) {
    var buf bytes.Buffer
    got := appendVString("xeroc", &buf)
    expected := []byte{5, 120, 101, 114, 111, 99}
    if !bytes.Equal(got.Bytes(), expected) {
        t.Error("Expected", expected, "got", got)
    }
}

func TestAppendVStringArray(t *testing.T) {
    var buf bytes.Buffer
    got := appendVStringArray([]string{"xeroc", "piston"}, &buf).Bytes()
    expected := []byte{2, 5, 120, 101, 114, 111, 99, 6, 112, 105, 115, 116, 111, 110}
    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}

func TestSerializeTx(t *testing.T) {
    got, _ := serializeTx(getTestVoteTx())
    expected := []byte{189, 140, 95, 226, 111, 69, 241, 121, 168, 87, 1, 0, 5, 120, 101, 114, 111, 99, 5, 120, 101, 114, 111, 99, 6, 112, 105, 115, 116, 111, 110, 16, 39, 0}
    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}

func TestSerializeOps(t *testing.T) {
    got, _ := serializeOps(getTwoTestOps())
    expected := []byte{2, 0, 5, 120, 101, 114, 111, 99, 5, 120, 101, 114, 111, 99, 6, 112, 105, 115, 116, 111, 110, 16, 39, 18, 0, 1, 5, 120, 101, 114, 111, 99, 7, 116, 101, 115, 116, 45, 105, 100, 17, 123, 34, 116, 101, 115, 116, 107, 34, 58, 34, 116, 101, 115, 116, 118, 34, 125}
    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}

func TestSerializeOpVoteOperation(t *testing.T) {
    got, _ := getTestVoteOp().serializeOp()
    expected := []byte{0, 5, 120, 101, 114, 111, 99, 5, 120, 101, 114, 111, 99, 6, 112, 105, 115, 116, 111, 110, 16, 39}
    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}

func TestSerializeOpCustomJsonOperation(t *testing.T) {
    got, _ := getTestCustomJsonOp().serializeOp()
    expected := []byte{18, 0, 1, 5, 120, 101, 114, 111, 99, 7, 116, 101, 115, 116, 45, 105, 100, 17, 123, 34, 116, 101, 115, 116, 107, 34, 58, 34, 116, 101, 115, 116, 118, 34, 125}
    if !bytes.Equal(got, expected) {
        t.Error("Expected", expected, "got", got)
    }
}