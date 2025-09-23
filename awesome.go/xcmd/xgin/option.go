package xgin

import (
	"github.com/gin-gonic/gin"
)

type Option func(*XGinServer)

// WithHost 设置host的Option函数
func WithHost(host string) Option {
	return func(s *XGinServer) {
		s.Host = host
	}
}

// WithPort 设置port的Option函数
func WithPort(port int) Option {
	return func(s *XGinServer) {
		s.Port = port
	}
}

func Use(middleware ...gin.HandlerFunc) Option {
	return func(s *XGinServer) {
		s.Gin.Use(middleware...)
	}
}

func Gin(op func(r *gin.Engine)) Option {
	return func(s *XGinServer) {
		op(s.Gin)
	}
}
