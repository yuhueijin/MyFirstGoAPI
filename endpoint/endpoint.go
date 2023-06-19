package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/yuhueijin/MyFirstGoAPI/service"
)


func MakeAddEndpoint(supportedService service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(AddRequest)
		err := supportedService.Add(input.Name)
		return &AddResponse{
			Err: err,
		}, nil
	}
}

func MakeRemoveEndpoint(supportedService service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		input := request.(RemoveRequest)
		err := supportedService.Remove(input.ID)
		return &RemoveResponse{
			Err: err,
		}, nil
	}
}

func MakeGetAllEndpoint(supportedService service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		res, err := supportedService.GetAll()
		return &GetAllResponse{
			Payload: res,
			Err:     err,
		}, nil
	}
}

type AddRequest struct {
	Name string `json:"name"`
}

type RemoveRequest struct {
	ID int `json:"id"`
}

type AddResponse struct {
	Err error
}

type RemoveResponse struct {
	Err error
}

type GetAllResponse struct {
	Payload []service.Model
	Err     error
}
