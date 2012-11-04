package main

import (
	"crypto/tls"
	"time"
	"bufio"
	"fmt"
	"log"
	"net"
)

var SERVER_CERT = []byte(`-----BEGIN CERTIFICATE-----
MIICATCCAWoCCQC19LSQvNIqmTANBgkqhkiG9w0BAQUFADBFMQswCQYDVQQGEwJG
UjETMBEGA1UECAwKU29tZS1TdGF0ZTEhMB8GA1UECgwYSW50ZXJuZXQgV2lkZ2l0
cyBQdHkgTHRkMB4XDTEyMTEwNDIxNDIzOVoXDTEzMTEwNDIxNDIzOVowRTELMAkG
A1UEBhMCRlIxEzARBgNVBAgMClNvbWUtU3RhdGUxITAfBgNVBAoMGEludGVybmV0
IFdpZGdpdHMgUHR5IEx0ZDCBnzANBgkqhkiG9w0BAQEFAAOBjQAwgYkCgYEAuPrm
CMjKxBC8M6QP/kpi3ccJzKGicHKVNwdQAFVuFvBEM4j6Ms4xKuFgGlorvizClEjx
PkeYtO/QttYIMTKng8gAEAuXvYoVelUvQTgbUQuaccBchbSfi3IjKAdfEeUTPXB0
24AIB2ePSJoxR5V3+PX91HY3B5nHFlvh3n+Q9a8CAwEAATANBgkqhkiG9w0BAQUF
AAOBgQB7rKZhPkhVYhoga7m8NKPfQkSkbDZThMITmanvjDB2h3gnphUVWVGnwF9g
seLi1bBtyb/5DxVzfzIWmmz36ck3OpPglTUMXgpwZg49QCdp4KPd4Six4crf46WM
GYkFaAWx59sbrYz658pFdZxm8uR+TV/+YJ8NhBgCnC1fioGv2g==
-----END CERTIFICATE-----`)
var SERVER_KEY = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQC4+uYIyMrEELwzpA/+SmLdxwnMoaJwcpU3B1AAVW4W8EQziPoy
zjEq4WAaWiu+LMKUSPE+R5i079C21ggxMqeDyAAQC5e9ihV6VS9BOBtRC5pxwFyF
tJ+LciMoB18R5RM9cHTbgAgHZ49ImjFHlXf49f3UdjcHmccWW+Hef5D1rwIDAQAB
AoGAe1VN6q12BCPkV8obn8CZCqWaswVR+Qds1bPac16CeGjaEEJUD2vK/HwoR2m3
oKTXV2cK8itqdbkvv7gU3jiX1vIcT8nBj4nVgDEoq5nCjXYc+nc9XIy+Gyd24X5E
usiDfb80zAcaAWSJVQ5iDuQ/alvhlKpPSCizYkSYDiAnowECQQDvAxVaV1AkW/3E
qN/rH+LMtai6l59JG2J/8njh0LVc3R/S5d+HMGeDYj0TKNJtCj0m9e9tYz9tFE29
nOJ3V/oHAkEAxiCuuoYdRyLiO1u/fmSoIBYWnQRuwBmIbtm8bWrIEPxuBv/fl+Ga
1oViYL77isUzKlshnRk6TSh0fO6Kp3NdGQJAbn4Va/s7UGO6kCSlx5OpDIvaYdBg
UbK4OYAFhBcxEKok4SFl0aB96g4LQAU6KjB2jsFZG0+rbajaO5MAWr2wFwJBAJsh
7g5IgtY3XzKhJTCfOfFdujkZxmoN1AEP5fU6ngqGzNQYN4fh824zJJOiFq1SuTxZ
/NjYbwkJOySzVsfEkckCQC5aFhUCtPOd1FCoPP9Z4okrCSaN+St8xgzNamq6RoWK
w6BX0Uv0jUqrOL+XglYKCtmkNEreu/5/4cuahNZB/lA=
-----END RSA PRIVATE KEY-----`)

func handleConnection(conn net.Conn) {
	data, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal("get client data error: ", err)
	}

	fmt.Printf("%#v\n", data)

	fmt.Fprintf(conn, "hello client\n")
	conn.Close()
}

func main() {
	config := &tls.Config{
		Time: time.Now,
	}
	config.Certificates = make([]tls.Certificate, 1)
	config.Certificates[0], _ = tls.X509KeyPair(SERVER_CERT, SERVER_KEY)
	ln, err := tls.Listen("tcp", ":6010", config)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal("get client connection error: ", err)
		}

		go handleConnection(conn)
	}
}
