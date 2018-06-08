package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pin/tftp"
)

var (
	addr      string
	path      string
	filename  string
	operation string
	mode      string
)

func init() {
	flag.StringVar(&addr, "addr", "localhost:69", "Server address")
	flag.StringVar(&path, "path", ".", "Local file path")
	flag.StringVar(&filename, "file", "<filename>", "Name of the file on server")
	flag.StringVar(&operation, "op", "<get|put>", "What to do: download or upload file")
	flag.StringVar(&mode, "mode", "octet", "Transfer mode: 'octet' or 'netascii'")
	flag.Parse()
}

func main() {
	if filename == "<filename>" {
		fmt.Fprintf(os.Stderr, "missing filename!\n\n")
		flag.Usage()
		os.Exit(1)
	}
	if mode != "netascii" && mode != "octet" {
		fmt.Fprintf(os.Stderr, "invalid mode: %s\n\n", mode)
		flag.Usage()
		os.Exit(1)
	}
	if operation == "put" {
		putFile(addr, path, filename, mode)
	} else if operation == "get" {
		getFile(addr, path, filename, mode)
	} else {
		fmt.Fprintf(os.Stderr, "missing or invalid operation!\n\n")
		flag.Usage()
		os.Exit(1)
	}
}

func putFile(addr string, path string, filename string, mode string) {
	c, err := tftp.NewClient(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	c.SetTimeout(5 * time.Second) //optional timeout

	file, err := os.Open(filepath.Join(path, filename))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer file.Close()

	rf, err := c.Send(filename, mode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	n, err := rf.ReadFrom(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	fmt.Printf("%d bytes sent\n", n)
}

func getFile(addr string, path string, filename string, mode string) {
	c, err := tftp.NewClient(addr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	wt, err := c.Receive(filename, mode)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	file, err := os.Create(filepath.Join(path, filename))
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	defer file.Close()

	//Optionally obtain transfer size before actual data.
	if n, ok := wt.(tftp.IncomingTransfer).Size(); ok {
		fmt.Printf("Transfer size: %d\n", n)
	}

	n, err := wt.WriteTo(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}
	fmt.Printf("%d bytes received\n", n)
}
