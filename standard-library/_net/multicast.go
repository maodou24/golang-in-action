package _net

import (
	"fmt"
	"net"
	"time"
)

func JoinMulticastGroup(name, addrstr string) error {
	addr, err := net.ResolveUDPAddr("udp", addrstr)
	if err != nil {
		return fmt.Errorf("resolve udp addr: %w", err)
	}

	conn, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		return fmt.Errorf("listen udp: %w", err)
	}
	defer conn.Close()

	err = conn.SetReadBuffer(1024 * 1024)
	if err != nil {
		return fmt.Errorf("set read buffer: %w", err)
	}
	// read multicast message and send unicast response
	buf := make([]byte, 1024)
	for {
		n, src, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP failed:", err)
			continue
		}

		fmt.Printf("%v Received %d bytes from %s: %s\n", name,
			n, src, string(buf[:n]))
		conn.WriteToUDP([]byte(fmt.Sprintf("%v response meassge %v", name, string(buf[:n]))), src)
	}
	return nil
}

func SendMulticastGroupWaitUnicastResponse(addrstr string) error {
	localAddr := ":8888"
	laddr, err := net.ResolveUDPAddr("udp", localAddr)
	if err != nil {
		return fmt.Errorf("resolve local udp addr: %w", err)
	}

	// list local udp addr to receive unicast response
	conn, err := net.ListenUDP("udp", laddr)
	if err != nil {
		return fmt.Errorf("listen local udp: %w", err)
	}
	defer conn.Close()

	// wait for unicast response
	go func() {
		buf := make([]byte, 1024)
		for {
			n, src, err := conn.ReadFromUDP(buf)
			if err != nil {
				fmt.Println("Read failed:", err)
				return
			}
			fmt.Printf("Received response from %s: %s\n", src, string(buf[:n]))
		}
	}()

	raddr, err := net.ResolveUDPAddr("udp", addrstr)
	if err != nil {
		return fmt.Errorf("resolve udp addr: %w", err)
	}
	// send multicast message
	for i := range 5 {
		_, err = conn.WriteToUDP([]byte(fmt.Sprintf("message %d", i)), raddr)
		if err != nil {
			return fmt.Errorf("write udp: %w", err)
		}

		time.Sleep(time.Second)
	}
	return nil
}
