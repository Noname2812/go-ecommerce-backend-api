package userserviceimpl

import (
	"context"
	"time"

	commonenum "github.com/Noname2812/go-ecommerce-backend-api/internal/common/enum"
	userpb "github.com/Noname2812/go-ecommerce-backend-api/internal/common/protogen/user"
	commonvo "github.com/Noname2812/go-ecommerce-backend-api/internal/common/vo"
	usermodel "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/model"
	userrepository "github.com/Noname2812/go-ecommerce-backend-api/internal/services/user/domain/repository"
)

type userServiceServer struct {
	userpb.UnimplementedUserServiceServer
	userInfoRepo userrepository.UserInfoRepository
}

func NewUserServiceServer(repo userrepository.UserInfoRepository) userpb.UserServiceServer {
	return &userServiceServer{
		userInfoRepo: repo,
	}
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {

	phone, err := commonvo.NewPhone(req.UserPhone)
	if err != nil {
		return nil, err
	}

	account, err := commonvo.NewEmail(req.UserAccount)
	if err != nil {
		return nil, err
	}

	var birthday *time.Time
	if req.UserBirthday != "" {
		t, err := time.Parse("2006-01-02", req.UserBirthday)
		if err != nil {
			return nil, err
		}
		birthday = &t
	}
	userInfo := &usermodel.UserInfo{
		UserAccount:  account.String(),
		UserNickname: req.UserNickName,
		UserState:    commonenum.Activated,
		UserGender:   commonenum.Gender(req.UserGender),
		UserPhone:    phone,
		UserBirthday: birthday,
	}

	userId, err := s.userInfoRepo.CreateUserInfo(ctx, userInfo)

	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{
		UserId: userId,
	}, nil
}
