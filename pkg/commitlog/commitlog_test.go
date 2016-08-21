package commitlog_test

import (
	"testing"

	"github.com/lucindo/krarup/pkg/commitlog"
)

func TestSingleIO(t *testing.T) {
	testDataIn := []byte("test")
	//	testDataOut := make([]byte, len(testDataIn))
	cmlog, err := commitlog.GetInstance("teste")

	tdinlen, err := cmlog.Write(testDataIn)
	if err != nil {
		t.Errorf("error writing to commitlog: %v", err)
	}
	if len(testDataIn) != tdinlen {
		t.Errorf("Excpected write len of %d and got %d", len(testDataIn), tdinlen)
	}
}
