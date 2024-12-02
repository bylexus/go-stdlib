package server

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
)

type HttpHandler struct {
	Conn *net.Conn
	Ctx  context.Context
}

func (h *HttpHandler) Handle() error {
	defer (*h.Conn).Close()

	bufReader := bufio.NewReaderSize(*h.Conn, 4096)

	// stop if context is cancelled:
	if h.Ctx.Err() != nil {
		return h.Ctx.Err()
	}
	// read request start line
	reqStartLine, err := h.readRequestStartLine(bufReader)
	if err != nil {
		return err
	}
	if reqStartLine == "" {
		return nil
	}

	fmt.Printf("Read request start line: %s\n", reqStartLine)

	headers, err := h.readHeaders(h.Ctx, bufReader)
	if err != nil {
		// fmt.Printf("Read and echoed %d bytes\n", written)
		// Answer:
		fmt.Fprint(*h.Conn, "HTTP/1.1 500 Internal Server Error\r\n")
		fmt.Fprint(*h.Conn, "Content-Type: text/plain\r\n")
		fmt.Fprint(*h.Conn, "Connection: Close\n\r")
		fmt.Fprint(*h.Conn, "\r\n")
		fmt.Fprintf(*h.Conn, "Cannot read headers: %s\r\n", err)

		return nil

	} else {
		// Answer:
		fmt.Fprint(*h.Conn, "HTTP/1.1 200 OK\r\n")
		fmt.Fprint(*h.Conn, "Content-Type: text/plain\r\n")
		fmt.Fprint(*h.Conn, "Connection: Close\r\n")
		fmt.Fprint(*h.Conn, "\r\n")
		fmt.Fprintf(*h.Conn, "Read headers: %s\r\n", headers)

		return nil
	}
}

func (h *HttpHandler) readRequest(ctx context.Context, bufReader *bufio.Reader) (*http.Request, error) {

	return nil, nil
}

func (h *HttpHandler) readRequestStartLine(bufReader *bufio.Reader) (string, error) {
	line, isPrefix, err := bufReader.ReadLine()
	if err != nil {
		if err == io.EOF {
			return string(line), nil
		}
		return "", err
	}
	if isPrefix {
		return "", fmt.Errorf("request start line too long")
	}
	return string(line), nil
}

func (h *HttpHandler) readHeaders(ctx context.Context, bufReader *bufio.Reader) ([]string, error) {
	maxHeaders := 1024
	maxHeaderLength := 8192
	headers := make([]string, 0)
	for {
		// stop if context is cancelled:
		if h.Ctx.Err() != nil {
			return nil, h.Ctx.Err()
		}

		line, isPrefix, err := bufReader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		if isPrefix {
			return nil, fmt.Errorf("header line too long")
		}
		if len(line) == 0 {
			break
		}
		if len(line) > maxHeaderLength {
			return nil, fmt.Errorf("header line too long")
		}
		header := string(line)
		fmt.Printf("Read header line: %s\n", header)
		headers = append(headers, header)
		if len(headers) > maxHeaders {
			return nil, fmt.Errorf("too many headers")
		}
	}

	return headers, nil
}
