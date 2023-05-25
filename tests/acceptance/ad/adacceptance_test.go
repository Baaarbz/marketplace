package ad

import (
	"barbz.dev/marketplace/tests/acceptance"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	acceptance.InitAcceptanceTest()
	exitCode := m.Run()
	acceptance.StopAcceptanceTest()
	os.Exit(exitCode)
}
