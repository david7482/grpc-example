package services

import (
	"errors"

	"pb"

	"golang.org/x/net/context"
)

type CalculateService struct {

}

func (this *CalculateService) Calculate(ctx context.Context, req *pb.CalculateRequest) (*pb.CalculateResponse, error) {
	response := pb.CalculateResponse{}

	switch req.Operator {
	case pb.CalculateRequest_ADD:
		response.Value = req.Value0 + req.Value1
	case pb.CalculateRequest_SUB:
		response.Value = req.Value0 - req.Value1
	case pb.CalculateRequest_MUL:
		response.Value = req.Value0 * req.Value1
	case pb.CalculateRequest_DIV:
		response.Value = req.Value0 / req.Value1
	default:
		return nil, errors.New("Invalid operator")
	}

	return &response, nil
}
