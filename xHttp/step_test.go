package xHttp_test

import (
	"os"
	"testing"

	"github.com/Min-Feng/goutils/xLog"
)

func TestMain(m *testing.M) {
	xLog.TestingMode()

	code := m.Run()
	os.Exit(code)
}