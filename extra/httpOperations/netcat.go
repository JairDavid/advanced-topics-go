package httpOperations

import (
	"fmt"
	"io"
	"net"
	"os"
)

func NetcatStart() {
	conn, err := net.Dial("tcp", "localhost:5000")
	if err != nil {
		fmt.Print(err)
	}

	done := make(chan struct{})

	go func() {
		io.Copy(os.Stdout, conn)
		done <- struct{}{}
	}()

	copyContent(conn, os.Stdin)
	conn.Close()
	<-done
}

func copyContent(std io.Writer, src io.Reader) {
	_, err := io.Copy(std, src)
	if err != nil {
		fmt.Print(err)
	}
}
