package main

import (
	"net/http"
)

type Server interface {
	Route(pattern string, handleFunc http.HandlerFunc)
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

// Route 路由实现
func (s *sdkHttpServer) Route(pattern string, handleFunc http.HandlerFunc) {
	panic("implement Server route")
}

// Start 实现
func (s *sdkHttpServer) Start(address string) error {
	panic("implement Server Start")
}

func NewServer() Server {
	return &sdkHttpServer{}
}