package checkquic

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io"
	"net/http"

	"github.com/lucas-clemente/quic-go"
	"github.com/lucas-clemente/quic-go/http3"
)

func Checkquic() {
	pool, err := x509.SystemCertPool()
	if err != nil {
	}
	var quicconf quic.Config
	roundTripper := &http3.RoundTripper{
		TLSClientConfig: &tls.Config{
			RootCAs: pool,
		},
		QuicConfig: &quicconf,
	}
    defer roundTripper.Close()
    hclient := &http.Client{
        Transport: roundTripper,
    }
    rsp,err := hclient.Get("https://cloudflare-quic.com")
    if err != nil {
        fmt.Println("quic blocked")
        return
    }
    body := &bytes.Buffer{}
    _, err = io.Copy(body, rsp.Body)
    if err != nil {
    }
    fmt.Println("UDP 443 can access internet")
}






