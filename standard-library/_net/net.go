package _net

import "net"

// echo the data back to the client
func EchoSocketServer(host string) error {
	listener, err := net.Listen("tcp4", host)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func() {
			defer conn.Close()

			buf := make([]byte, 1024)
			n, err := conn.Read(buf)
			if err != nil {
				return
			}
			conn.Write(buf[:n])
		}()
	}

	return nil
}
