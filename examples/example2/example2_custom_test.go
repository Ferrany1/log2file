package example2

import (
	"github.com/Ferrany1/log2file/src/directory"
	"os"
	"path"
	"strings"
	"testing"
)

func TestExampleCustomOptions(t *testing.T) {
	var (
		mFileName = map[string]bool{"log_main.log": false, "log_backup.log": false}
		logText   = "Message:test"
		extraPath = "./testDir"
	)

	for i := 0; i < 2; i++ {
		ExampleCustomOptions()
	}

	_, _, err := directory.ReadDirectory(extraPath)
	if err != nil {
		t.Errorf("failed to read dir: %s", err.Error())
	}

	for fileName, _ := range mFileName {
		b, err := os.ReadFile(path.Join(extraPath, fileName))
		if err != nil {
			t.Errorf("failed to read file %s: %s", fileName, err.Error())
		}

		if !strings.Contains(string(b), logText) {
			t.Errorf("wrong log text inside: %s", fileName)
		}

	}
}
