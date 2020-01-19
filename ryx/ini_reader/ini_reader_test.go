package ini_reader_test

import (
	"github.com/tlarsen7572/Golang-Public/ryx/ini_reader"
	"path/filepath"
	"testing"
)

func TestIniReader(t *testing.T) {
	ini := ini_reader.LoadIni(filepath.Join(`.`, `R-3.5.3_Macros.ini`))
	if ini[`Path`] != `C:\Program Files\Alteryx\R-3.5.3\Plugin\Macros` {
		t.Fatalf(`expected Path of 'C:\Program Files\Alteryx\R-3.5.3\Plugin\Macros' but got '%v'`, ini[`Path`])
	}

	ini = ini_reader.LoadIni(filepath.Join(`.`, `User_1d706833-09e8-4887-8f78-db8d9f071248.ini`))
	if ini[`Path`] != `C:\Users\tlarsen\Documents\SAP Sets Tutorial` {
		t.Fatalf(`expected Path of 'C:\Users\tlarsen\Documents\SAP Sets Tutorial' but got '%v'`, ini[`Path`])
	}
}
