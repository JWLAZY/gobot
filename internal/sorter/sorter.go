package sorter

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"math/big"

	"github.com/lucas-clemente/quic-go"
)

type Sorter struct {
}

var (
	sorter Sorter
)

const addr = "localhost:4334"

func init() {

}

func Run() error {
	listener, err := quic.ListenAddr(addr, generateTLSConfig(), nil)
	if err != nil {
		return err
	}
	fmt.Println("sorter start network, listen addr:", addr)
	for {
		conn, err := listener.Accept(context.Background())
		if err != nil {
			return err
		}
		go sorter.handleConn(conn)
	}
}

func (sorter *Sorter) handleConn(conn quic.Connection) {
	for {
		stream, err := conn.AcceptStream(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Println("streamId:", stream.StreamID())
		go ReceiveData(stream)
	}
}

func ReceiveData(stream quic.Stream) {
	r := make([]byte, 1000)
	receiveCount := 0
	for {
		c, err := stream.Read(r[receiveCount:])
		if err != nil {
			fmt.Println(err)
		}
		if c == 0 {

		} else {
			receiveCount += c
		}

		if receiveCount >= 1000 {

		}
	}
}

func Stop() {

}

// Setup a bare-bones TLS config for the server
func generateTLSConfig() *tls.Config {
	key, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		panic(err)
	}
	template := x509.Certificate{SerialNumber: big.NewInt(1)}
	certDER, err := x509.CreateCertificate(rand.Reader, &template, &template, &key.PublicKey, key)
	if err != nil {
		panic(err)
	}
	keyPEM := pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(key)})
	certPEM := pem.EncodeToMemory(&pem.Block{Type: "CERTIFICATE", Bytes: certDER})

	tlsCert, err := tls.X509KeyPair(certPEM, keyPEM)
	if err != nil {
		panic(err)
	}
	return &tls.Config{
		Certificates: []tls.Certificate{tlsCert},
		NextProtos:   []string{"quic-echo-example"},
	}
}
