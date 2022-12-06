package service

import (
	"context"
	"fmt"
	"time"

	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/entity"
	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/exception"
	customerpb "github.com/sangianpatrick/dpe-ss-demo-grpc-server/pb/customer"
	"github.com/sangianpatrick/dpe-ss-demo-grpc-server/repository"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CustomerService struct {
	*customerpb.UnimplementedCustomerServer
	location   *time.Location
	logger     *logrus.Logger
	repository repository.AccountRepository
}

func NewCustomerService(location *time.Location, logger *logrus.Logger, repository repository.AccountRepository) *CustomerService {
	return &CustomerService{&customerpb.UnimplementedCustomerServer{}, location, logger, repository}
}

func (s *CustomerService) Register(ctx context.Context, req *customerpb.AccountRegistrationRequest) (resp *customerpb.AccountRegistrationResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*30)
	defer cancel()

	resp = new(customerpb.AccountRegistrationResponse)
	_, err = s.repository.FindByEmail(ctx, req.Email)
	if err == nil {
		return nil, status.Error(codes.AlreadyExists, "account is already registered")
	}

	if err != exception.ErrNotFound {
		if err == exception.ErrTimeout {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}

		if err == exception.ErrCancel {
			return nil, status.Error(codes.Canceled, err.Error())
		}

		return nil, status.Error(codes.Internal, "an error occured while creating new account")
	}

	account := entity.Account{
		Email:     req.Email,
		Password:  req.Password,
		Name:      req.Name,
		CreatedAt: time.Now().In(s.location),
	}

	id, err := s.repository.Save(ctx, account)
	if err != nil {
		if err == exception.ErrTimeout {
			return nil, status.Error(codes.DeadlineExceeded, err.Error())
		}

		if err == exception.ErrCancel {
			return nil, status.Error(codes.Canceled, err.Error())
		}

		return nil, status.Error(codes.Internal, "an error occured while creating new account")
	}

	account.Id = id

	resp = new(customerpb.AccountRegistrationResponse)
	resp.Account = &customerpb.Account{
		Id:        account.Id,
		Email:     account.Email,
		Name:      account.Name,
		CreatedAt: timestamppb.New(account.CreatedAt),
	}

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

func (s *CustomerService) MakePayment(ctx context.Context, req *customerpb.MakePaymentRequest) (resp *customerpb.MakePaymentResponse, err error) {
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	if req.OrderId%5 == 0 {
		grpc.SetTrailer(ctx, metadata.Pairs("status", "INSUFFICIENT_BALANCE"))
		err = status.Error(codes.Aborted, "cannot proceed the payment due to insufficient balance")
		return
	}

	resp = &customerpb.MakePaymentResponse{
		ReferenceNumber: generateReferenceNumber(),
		PaymentDate:     timestamppb.New(time.Now()),
	}

	return
}

func generateReferenceNumber() string {
	now := time.Now()
	referenceNumber := fmt.Sprintf("REF%d", now.Unix())

	return referenceNumber
}
