package server

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

type CmdServer struct {
	s service.Service
}

func MustNewServer(c service.ServiceConf, s service.Service) *CmdServer {
	server, err := NewCmdServer(c, s)
	logx.Must(err)
	return server
}

func NewCmdServer(c service.ServiceConf, s service.Service) (*CmdServer, error) {
	if err := c.SetUp(); err != nil {
		return nil, err
	}
	return &CmdServer{
		s: s,
	}, nil
}

func (s *CmdServer) Start() {
	s.s.Start()
}

func (s *CmdServer) Stop() {
	logx.Close()
	s.s.Stop()
}
