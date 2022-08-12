package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	m := strings.Fields(ln)[0]
	u := strings.Fields(ln)[1]
	b := new(strings.Builder)
	fmt.Println(m, u)
	// multiplexer
	if m == "GET" && u == "/" {
		if err := tpl.ExecuteTemplate(b, "index.gohtml", nil); err != nil {
			log.Fatalln(err)
		}
		response(b.String(), conn)
	} else if m == "GET" && u == "/about" {
		if err := tpl.ExecuteTemplate(b, "about.gohtml", nil); err != nil {
			log.Fatalln(err)
		}
		response(b.String(), conn)
	} else if m == "GET" && u == "/contact" {
		if err := tpl.ExecuteTemplate(b, "contact.gohtml", nil); err != nil {
			log.Fatalln(err)
		}
		response(b.String(), conn)
	} else if m == "GET" && u == "/apply" {
		if err := tpl.ExecuteTemplate(b, "apply.gohtml", nil); err != nil {
			log.Fatalln(err)
		}
		response(b.String(), conn)
	} else if m == "POST" && u == "/apply" {
		if err := tpl.ExecuteTemplate(b, "apply_process.gohtml", nil); err != nil {
			log.Fatalln(err)
		}
		response(b.String(), conn)
	} else {
		if err := tpl.ExecuteTemplate(b, "404.gohtml", nil); err != nil {
			log.Fatalln(err)
		}
		response(b.String(), conn)
	}
}

func response(body string, conn net.Conn) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
