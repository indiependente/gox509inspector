[![Go Report Card](https://goreportcard.com/badge/github.com/indiependente/gox509inspector)](https://goreportcard.com/report/github.com/indiependente/gox509inspector)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-0%25-brightgreen.svg?longCache=true&style=flat)</a>
[![Workflow Status](https://github.com/indiependente/gox509inspector/workflows/lint-test/badge.svg)](https://github.com/indiependente/gox509inspector/actions)
# gox509inspector
Go tool to inspect x.509 certificates in PEM format

### API

```go
/*
Parsex509Cert takes an io.Reader in input and returns an x509 certificate and nil
if no errors occurred while reading and parsing the certificate.
If the certificate has reminder data, both the certificate and the error will be returned.
*/
func Parsex509Cert(r io.Reader) (*x509.Certificate, error)
```

```go
/*
GetQuickInfo returns the certificate's most useful info as a byte array

Version             int
SerialNumber        *big.Int
Issuer              pkix.Name
Subject             pkix.Name
NotBefore, NotAfter time.Time // Validity bounds.
KeyUsage            KeyUsage
*/
func GetQuickInfo(c *x509.Certificate) []byte
```

### Example usage
```go
package main

import (
	"fmt"
	"gox509inspector/inspector"
	"os"
)

func main() {
	// pipe certificate in
	c, err := inspector.Parsex509Cert(os.Stdin)
	if c == nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
	}

	bytes := inspector.GetQuickInfo(c)
	fmt.Print(string(bytes))
}

// Output
// Version:        3
// SerialNumber:   2098790
// Issuer:         CN=CA,O=SSOCircle,C=DE
// Subject:        CN=idp.ssocircle.com,O=SSOCircle,C=DE
// Valid from:     2016-08-03 16:03:23 +0100 BST
// Valid to:       2026-03-04 15:03:23 +0000 GMT
```
