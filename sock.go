package main

import (
	"log"
	"net"
	"os"
	"os/exec"
)

var wkd string

const SockAddr = "C:\\tmp\\socktest.sock"

func echoServer(c net.Conn) {
	log.Printf("Client connected [%s]", c.RemoteAddr().Network())

	buf := make([]byte, 512)
	nr, err := c.Read(buf)

	if err != nil {
		return
	}

	data := buf[0:nr]
	println("Server got:", string(data))

	ps1("test.ps1")

	_, err = c.Write(data)
	if err != nil {
		log.Fatal("Write: ", err)
	}

	c.Close()
}

func ps1(scriptName string) {
	path := wkd + "\\" + scriptName
	log.Print(path)
	cmd := exec.Command("C:\\Windows\\System32\\WindowsPowerShell\\v1.0\\powershell.exe", "-file", "C:\\git\\sockecho\\test.ps1", "-state", "dev")
	//cmd := exec.Command("C:\\Users\\George\\Downloads\\npcap-1.60.exe")

	//cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	//cmd_output, err := cmd.Output()

	err := cmd.Run()
	if err != nil {
		log.Print(err)

	}

}

func main() {

	wkd, _ = os.Getwd()

	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		// Accept new connections, dispatching them to echoServer
		// in a goroutine.
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go echoServer(conn)
	}
}
