package service

import (
	"time"

	"github.com/captrep/go-crud-rest-api/model/domain"
	"github.com/captrep/go-crud-rest-api/model/web"
	"github.com/captrep/go-crud-rest-api/repository"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserService struct {
	UserRepository *repository.UserRepository
	Validate       *validator.Validate
}

func NewUserService(UserRepository *repository.UserRepository, validate *validator.Validate) *UserService {
	return &UserService{
		UserRepository: UserRepository,
		Validate:       validate,
	}
}

func (svc *UserService) CreateUser(req *web.CreateUserRequest) (*web.UserResponse, error) {
	var err error
	id := uuid.New()
	req.Id = id.String()
	err = svc.Validate.Struct(req)
	if err != nil {
		return nil, err
	}
	user := domain.User{
		Id:        req.Id,
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	user, err = svc.UserRepository.Save(user)
	if err != nil {
		return nil, err
	}

	return &web.UserResponse{
		Id:        user.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (svc *UserService) FindAll() ([]web.UserResponse, error) {
	var userResponse []web.UserResponse
	result, err := svc.UserRepository.GetAll()
	if err != nil {
		return nil, err
	}
	for _, r := range result {
		userResponse = append(userResponse, web.UserResponse{
			Id:        r.Id,
			Firstname: r.Firstname,
			Lastname:  r.Lastname,
			Email:     r.Email,
			CreatedAt: r.CreatedAt,
			UpdatedAt: r.UpdatedAt,
		})
	}
	return userResponse, nil
}

func (service *UserService) FindById(userID string) (*web.UserResponse, error) {
	result, err := service.UserRepository.FindById(userID)
	if err != nil {
		return nil, err
	}
	return &web.UserResponse{
		Id:        result.Id,
		Firstname: result.Firstname,
		Lastname:  result.Lastname,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}, nil
}

func (svc *UserService) Update(req *web.UpdateUserRequest) (*web.UserResponse, error) {
	curr, err := svc.UserRepository.FindById(req.Id)
	if err != nil {
		return nil, err
	}
	curr.Firstname = req.Firstname
	curr.Lastname = req.Lastname
	curr.Email = req.Email
	curr.UpdatedAt = time.Now()

	res, errU := svc.UserRepository.Update(curr)
	if errU != nil {
		return nil, errU
	}

	return &web.UserResponse{
		Id:        res.Id,
		Firstname: res.Firstname,
		Lastname:  res.Lastname,
		Email:     res.Email,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}

func (svc *UserService) Delete(userID string) (*web.WebResponse, error) {
	curr, err := svc.UserRepository.FindById(userID)
	if err != nil {
		return nil, err
	}
	errD := svc.UserRepository.Delete(curr)
	if errD != nil {
		return nil, errD
	}

	return &web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   curr.Id,
	}, nil
}
