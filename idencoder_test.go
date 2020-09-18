package idencoder

import (
    "fmt"
    "math/rand"
    "testing"
    "time"
)

var (
    codePasswordCasesArgFirst  = []string{"netmt", "market", "yaxin"}
    codePasswordCasesArgSecond = []int{41, 64, 64}
    codePasswordCasesArgRet    = []int{474148531572, 120265298896244, 521326324078}
)

func TestEncodeDecode(t *testing.T) {
    rand.Seed(time.Now().Unix())

    for i := 0; i < 100; i++ {
        uniqueId := rand.Int63n(2199023255521 - 1)
        encoded, err := Encode(uniqueId, "nick", 8)
        if err != nil {
            t.Errorf("encode failed: %s", err.Error())
        }
        actual, err := Decode(encoded, "nick")
        if err != nil {
            t.Errorf("decode failed: %s", err.Error())
        }
        if uniqueId != actual {
            t.Errorf("expected: %d, actual: %d", uniqueId, actual)
        }
        fmt.Println(uniqueId, encoded)
    }
}

func TestEncodePassword(t *testing.T) {
    for idx, ret := range codePasswordCasesArgRet {
        actual := encodePassword(codePasswordCasesArgFirst[idx], codePasswordCasesArgSecond[idx])
        if actual != ret {
            t.Errorf(
                "encodePassword(%s, %d) != %d, actual is %d",
                codePasswordCasesArgFirst[idx], codePasswordCasesArgSecond[idx], ret, actual,
            )
        }
    }
}

func TestDecodePassword(t *testing.T) {
    for idx, v := range codePasswordCasesArgRet {
        actual := decodePassword(v)
        expected := codePasswordCasesArgFirst[idx]
        if actual != expected {
            t.Errorf("decodePassword(%d) expected %s, actual %s", v, expected, actual)
        }
    }
}

func TestExtendEuclid(t *testing.T) {

}
