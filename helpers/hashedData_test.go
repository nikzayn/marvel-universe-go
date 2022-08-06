package helpers

import (
	"os"
	"strconv"
	"testing"
	"time"
)

type HashTest struct {
	arg1, arg2, expected string
}

func TestAdd(t *testing.T) {

	publicKey := os.Getenv("PUBLIC_KEY")
	currentTime := strconv.Itoa(int(time.Now().Unix()))

	addTests := []HashTest{
		{currentTime, publicKey, HashedString(currentTime, publicKey)},
		{"abc1234", publicKey, HashedString("abc1234", publicKey)},
		{"ABc1234", publicKey, HashedString("ABc1234", publicKey)},
		{"Abc@#$1234", publicKey, HashedString("Abc@#$1234", publicKey)},
	}

	for _, test := range addTests {
		if output := HashedString(test.arg1, test.arg2); output != test.expected {
			t.Errorf("Output %q not equal to expected %q", output, test.expected)
		}
	}
}
