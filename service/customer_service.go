package service

import (
	"context"
	"fmt"
	"time"

	customerpb "github.com/sangianpatrick/dpe-ss-demo-grpc-server/pb/customer"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerService struct {
	*customerpb.UnimplementedCustomerServer
	logger *logrus.Logger
}

func NewCustomerService(logger *logrus.Logger) *CustomerService {
	return &CustomerService{&customerpb.UnimplementedCustomerServer{}, logger}
}

func (s *CustomerService) Register(ctx context.Context, req *customerpb.AccountRegistrationRequest) (resp *customerpb.AccountRegistrationResponse, err error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	resp = new(customerpb.AccountRegistrationResponse)

	account := new(customerpb.Account)
	account.Id = 1
	account.Email = req.Email
	account.Name = req.Name
	account.CreatedAt = timestamppb.New(time.Now())
	account.UpdatedAt = nil

	resp.Account = account

	return
}
func (s *CustomerService) SubscribeNotification(req *customerpb.SubscribeNotificationRequest, stream customerpb.Customer_SubscribeNotificationServer) (err error) {
	ctx, cancel := context.WithTimeout(stream.Context(), time.Second*5)
	defer cancel()

	var id int64 = 0
	run := true

	for run {
		select {
		case <-ctx.Done():
			run = false

		default:
			stream.Send(&customerpb.Notification{
				Id:      id,
				Topic:   req.Topic,
				Message: fmt.Sprintf("Hi, You got a message with id %d", id),
			})

			id++

			time.Sleep(time.Millisecond * 500)
		}
	}

	return
}
func (s *CustomerService) SumNumbers(stream customerpb.Customer_SumNumbersServer) (err error) {
	ctx, cancel := context.WithTimeout(stream.Context(), time.Second*10)
	defer cancel()

	var total int64

	run := true
	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			req, err := stream.Recv()
			if err != nil {
				s.logger.Error(err)
				return err
			}

			total = total + req.Input
		}
	}

	stream.SendAndClose(&customerpb.SumNumbersResponse{
		Aggregate: total,
	})

	return
}
func (s *CustomerService) Chat(stream customerpb.Customer_ChatServer) (err error) {
	ctx, cancel := context.WithTimeout(stream.Context(), time.Second*15)
	defer cancel()

	run := true
	for run {
		select {
		case <-ctx.Done():
			run = false
		default:
			in, err := stream.Recv()
			if err != nil {
				s.logger.Error(err)
				return err
			}

			s.logger.WithFields(logrus.Fields{
				"message": in.Message,
				"sent_at": in.CreatedAt.AsTime().Format(time.RFC3339),
			}).Info("incomming message")

			stream.Send(&customerpb.ChatResponse{
				SenderId:  -1,
				Message:   "Hello from server",
				CreatedAt: timestamppb.New(time.Now()),
			})
		}
	}
	return
}
