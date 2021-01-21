package example1

import (
	"github.com/Ferrany1/log2file/src/directory"
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

func TestExampleEmptyOptions(t *testing.T) {
	var (
		mFileName = map[string]bool{"log_1.log": false, "log_2.log": false}
		logText   = "test"
	)

	for i := 0; i < 2; i++ {
		ExampleEmptyOptions()
	}

	fi, dir, err := directory.ReadCurrentDirectory()
	if err != nil {
		t.Errorf("failed to read dir: %s", err.Error())
	}

	for _, f := range fi {
		if _, ok := mFileName[f.Name()]; ok {
			mFileName[f.Name()] = true

			b, err := ioutil.ReadFile(path.Join(dir, f.Name()))
			if err != nil {
				t.Errorf("failed to read file %s: %s", f.Name(), err.Error())
			}

			if !strings.Contains(string(b), logText) {
				t.Errorf("wrong log text inside: %s", f.Name())
			}
		}
	}

	for k, v := range mFileName {
		if v != true {
			t.Errorf("no such file: %s", k)
		}
	}
}
