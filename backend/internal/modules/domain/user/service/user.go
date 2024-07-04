package service

import (
	"context"
	"fmt"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/pkg/transaction"
)

type IUserRepo interface {
	GetByUuid(ctx context.Context, uuid string) (*dto.User, error)
	GetRolesByUuid(ctx context.Context, uuid string) (dto.Roles, error)
	List(ctx context.Context) (dto.Users, error)

	ListByUuids(ctx context.Context, uuids []string) (dto.Users, error)
	Update(ctx context.Context, uuid string, dtm *dto.UserUpdate) (*dto.User, error)
	Delete(ctx context.Context, uuid string) error
	Restore(ctx context.Context, uuid string) (*dto.User, error)

	AddRole(ctx context.Context, uuid string, role roles.Type) error
	RemoveRole(ctx context.Context, uuid string, role roles.Type) error

	transaction.TxRepo
}

type UserService struct {
	repo IUserRepo
}

func NewUserService(repo IUserRepo) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (r *UserService) GetByUuid(ctx context.Context, uuid string) (*dto.User, error) {
	user, err := r.repo.GetByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить пользователя: %w", err)
	}

	return user, nil
}

func (r *UserService) GetRolesByUuid(ctx context.Context, uuid string) (dto.Roles, error) {
	userRoles, err := r.repo.GetRolesByUuid(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить роли пользователя: %w", err)
	}

	return userRoles, nil
}

func (r *UserService) ListByUuids(ctx context.Context, uuids []string) (dto.Users, error) {
	users, err := r.repo.ListByUuids(ctx, uuids)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить пользователей: %w", err)
	}

	return users, nil
}

func (r *UserService) List(ctx context.Context) (dto.Users, error) {
	users, err := r.repo.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("не удалось получить список пользователей: %w", err)
	}

	return users, nil
}

func (r *UserService) Update(ctx context.Context, uuid string, dtm *dto.UserUpdate) (*dto.User, error) {
	newUser, err := r.repo.Update(ctx, uuid, dtm)
	if err != nil {
		return nil, fmt.Errorf("не удалось обновить пользователя: %w", err)
	}

	return newUser, nil
}

func (r *UserService) AddRole(ctx context.Context, uuid string, role roles.Type) (dto.Roles, error) {
	err := r.repo.AddRole(ctx, uuid, role)
	if err != nil {
		return nil, fmt.Errorf("не удалось добавить роль: %w", err)
	}

	return r.repo.GetRolesByUuid(ctx, uuid)
}

func (r *UserService) RemoveRole(ctx context.Context, uuid string, role roles.Type) (dto.Roles, error) {
	err := r.repo.RemoveRole(ctx, uuid, role)
	if err != nil {
		return nil, fmt.Errorf("не удалось удалить роль: %w", err)
	}

	return r.repo.GetRolesByUuid(ctx, uuid)
}

func (r *UserService) Delete(ctx context.Context, uuid string) error {
	err := r.repo.Delete(ctx, uuid)
	if err != nil {
		return fmt.Errorf("не удалось удалить пользователя: %w", err)
	}

	return nil
}

func (r *UserService) Restore(ctx context.Context, uuid string) (*dto.User, error) {
	user, err := r.repo.Restore(ctx, uuid)
	if err != nil {
		return nil, fmt.Errorf("не удалось восстановить пользователя: %w", err)
	}

	return user, nil
}
