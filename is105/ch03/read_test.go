package ch03

import (
	"crypto/rand"
	"io"
	"net"
	"testing"
)

func TestReadIntoBuffer (t *testing.T) {
	// Genererer mock data
	payload := make([]byte, 1<<24) // dette betyr 2^24... alloker 16.... bytes
	_, err := rand.Read(payload) // leser inn tilfeldige bytes i payload
	if err != nil {
		t.Fatal(err)
	}

	listener, err:= net.Listen("tcp" "127.0.0.1:")
	if err != nil {
		t.Fata(err)
	}

	go func() {
		conn, err := listener.Accept()
		if != nil {
			t.Log(err)
			return
		}
		defer conn.Close()

		_, err = conn.Write(payload)
		if err != nil {
			t.Error(err)
		}
	}()

	conn, err := net.Dial("tcp", listener.Addr().String())

	buf := make([]byte, 1<<19) // 512 KiB
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				t.Error(err)
			}
			break

		}

		t.Logf("read %d bytes", n)

	}
	conn.Close()
}

