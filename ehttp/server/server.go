package server

import (
	"context"
	"fmt"
	"net"
)

type Server struct {
	opts           ServerOpts
	listener       *net.Listener
	isShuttingDown bool
}

func NewServer(opts ServerOpts) *Server {
	server := &Server{
		opts:           opts,
		isShuttingDown: false,
	}
	return server
}

func (s *Server) Boot() error {
	fmt.Println("Booting Server")
	fmt.Printf("Listening on %s\n", s.opts.Address)

	conn, err := net.Listen("tcp", s.opts.Address)
	if err != nil {
		return err
	}
	s.listener = &conn

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)

	for {
		if s.isShuttingDown {
			cancel()
			return nil
		}
		conn, err := (*s.listener).Accept()
		if err != nil {
			cancel()
			return err
		}
		go s.handle(&conn, ctx)
	}
}

func (s *Server) handle(conn *net.Conn, ctx context.Context) {
	fmt.Println("Handling Connection")
	h := HttpHandler{
		Conn: conn,
		Ctx:  ctx,
	}

	err := h.Handle()
	if err != nil {
		fmt.Println(err)
	}
}

func (s *Server) Shutdown() error {

	fmt.Println("Shutdown Server")
	s.isShuttingDown = true
	if *s.listener != nil {
		return (*s.listener).Close()
	}
	return nil
}
