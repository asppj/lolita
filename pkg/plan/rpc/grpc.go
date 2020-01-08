package rpc

import (
	"github.com/asppj/t-go-opentrace/ext/grpc-driver/grpc"
)

var addr = "192.168.253.73:6005"

// NewPlanDial dial
func NewPlanDial() (*grpc.ClientConn, error) {
	return grpc.Dial(addr)
}

// NewTaskDial dial
func NewTaskDial() (*grpc.ClientConn, error) {
	return grpc.Dial(addr)
}
