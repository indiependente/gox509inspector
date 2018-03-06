package test

import (
	"fmt"
	"gox509inspector/inspector"
	"math/big"
	"os"
	"testing"
	"time"
)

func TestParsex509Cert(t *testing.T) {

	f, err := os.Open("../data/cert.pem")
	if err != nil {
		t.Errorf("error during test certificate open: %v", err)
	}

	c, err := inspector.Parsex509Cert(f)
	if c == nil {
		t.Errorf("error during test certificate parsing: %v", err)
	}
	if err != nil {
		t.Errorf("warning during test certificate parsing: %v", err)
	}

	layout := "2006-01-02 15:04:05 -0700 MST"
	strBefore := "2016-08-03 16:03:23 +0100 BST"
	strAfter := "2026-03-04 15:03:23 +0000 GMT"
	tNotBefore, err := time.Parse(layout, strBefore)
	if err != nil {
		fmt.Println(err)
	}
	tNotAfter, err := time.Parse(layout, strAfter)
	if err != nil {
		fmt.Println(err)
	}
	if c.Version != 3 {
		t.Errorf("error: parsed version data differs from expected data")
	}

	if c.SerialNumber.Cmp(big.NewInt(2098790)) != 0 {
		t.Errorf("error: parsed serial number data differs from expected data: %d - %d", c.SerialNumber, big.NewInt(2098790))
	}
	if c.Issuer.String() != "CN=CA,O=SSOCircle,C=DE" {
		t.Errorf("error: parsed issuer data differs from expected data")
	}
	if c.Subject.String() != "CN=idp.ssocircle.com,O=SSOCircle,C=DE" {
		t.Errorf("error: parsed subject data differs from expected data")
	}
	if !c.NotBefore.Equal(tNotBefore) {
		t.Errorf("error: parsed validfrom data differs from expected data")
	}
	if !c.NotAfter.Equal(tNotAfter) {
		t.Errorf("error: parsed validto data differs from expected data")
	}
}
