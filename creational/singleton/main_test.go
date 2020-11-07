package singleton

import (
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	for i := 0; i< 100000;i++  {
		go getInstance()
	}

	time.Sleep(20 * time.Second)
}
