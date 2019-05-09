// Copyright (c) 2019 Minoru Osuka
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package federator

import (
	"log"
	"net"

	"github.com/mosuka/blast/protobuf/federation"
	"google.golang.org/grpc"
)

type GRPCServer struct {
	service  *GRPCService
	server   *grpc.Server
	listener net.Listener

	logger *log.Logger
}

func NewGRPCServer(managerAddr string, grpcAddr string, logger *log.Logger) (*GRPCServer, error) {
	server := grpc.NewServer()

	service, err := NewGRPCService(managerAddr, logger)
	if err != nil {
		return nil, err
	}

	federation.RegisterFederationServer(server, service)

	listener, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		return nil, err
	}

	return &GRPCServer{
		service:  service,
		server:   server,
		listener: listener,
		logger:   logger,
	}, nil
}

func (s *GRPCServer) Start() error {
	err := s.service.StartWatchCluster()
	if err != nil {
		return err
	}

	err = s.server.Serve(s.listener)
	if err != nil {
		return err
	}

	return nil
}

func (s *GRPCServer) Stop() error {
	err := s.service.StopWatchCluster()
	if err != nil {
		return err
	}

	s.server.GracefulStop()

	return nil
}