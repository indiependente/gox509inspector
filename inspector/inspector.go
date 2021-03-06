package inspector

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"strings"
)

/*Parsex509Cert takes an io.Reader in input and returns an x509 certificate and nil
if no errors occurred while reading and parsing the certificate.
If the certificate has reminder data, both the certificate and the error will be returned.
*/
func Parsex509Cert(r io.Reader) (*x509.Certificate, error) {
	var err error
	pemData, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("error while reading the certificate: %w", err)
	}

	var block *pem.Block
	block, rest := pem.Decode(pemData)
	if block == nil {
		return nil, fmt.Errorf("error: PEM not parsed")
	}
	if len(rest) != 0 {
		return nil, fmt.Errorf("warning: reminder data\n%v", rest)
	}

	c, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("error while parsing the certificate: %w", err)
	}
	return c, err
}

/*
GetQuickInfo returns the certificate's most useful info as a byte array

Version             int
SerialNumber        *big.Int
Issuer              pkix.Name
Subject             pkix.Name
NotBefore, NotAfter time.Time // Validity bounds.
KeyUsage            KeyUsage
*/
func GetQuickInfo(c *x509.Certificate) []byte {
	info := []byte(fmt.Sprintf("Version:\t%d\n", c.Version))
	info = append(info, []byte(fmt.Sprintf("SerialNumber:\t%d\n", c.SerialNumber))...)
	info = append(info, []byte(fmt.Sprintf("Issuer:\t\t%s\n", c.Issuer))...)
	info = append(info, []byte(fmt.Sprintf("Subject:\t%s\n", c.Subject))...)
	info = append(info, []byte(fmt.Sprintf("DNS Names:\t%s\n", strings.Join(c.DNSNames, ",")))...)
	info = append(info, []byte(fmt.Sprintf("IP Addresses:\t%s\n", joinIPs(c.IPAddresses, ",")))...)
	info = append(info, []byte(fmt.Sprintf("Valid from:\t%s\n", c.NotBefore.Local()))...)
	info = append(info, []byte(fmt.Sprintf("Valid to:\t%s\n", c.NotAfter.Local()))...)
	return info
}

func joinIPs(ips []net.IP, sep string) string {
	var ipsNo int
	if ipsNo = len(ips); ipsNo == 0 {
		return ""
	}
	var sb strings.Builder
	for _, ip := range ips[:len(ips)-1] {
		sb.WriteString(ip.To4().String())
		sb.WriteString(sep)
	}
	sb.WriteString(ips[len(ips)-1].To4().String())
	return sb.String()
}
