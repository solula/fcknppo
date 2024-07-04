package repo

import (
	"context"
	"errors"
	"fmt"
	"github.com/lib/pq"
	"strings"
	"waterfall-backend/internal/constants/err_const"
	"waterfall-backend/internal/models/roles"
	"waterfall-backend/internal/modules/domain/user/dto"
	"waterfall-backend/internal/modules/stores/db/converters"
	"waterfall-backend/internal/modules/stores/db/ent"
	"waterfall-backend/internal/modules/stores/db/ent/user"
	"waterfall-backend/internal/modules/stores/db/schema"
	"waterfall-backend/internal/modules/stores/db/utils"
)

type UserRepo struct {
	client *ent.Client
}

func NewUserRepo(client *ent.Client) *UserRepo {
	return &UserRepo{
		client: client,
	}
}

func (r *UserRepo) GetByUuid(ctx context.Context, uuid string) (*dto.User, error) {
	usr, err := r.client.User.Get(schema.SkipSoftDelete(ctx), uuid)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToUserDTO(usr), nil
}

func (r *UserRepo) GetExtendedByUuid(ctx context.Context, uuid string) (*dto.ExtendedUser, error) {
	userQuery := r.client.User.Query().Where(user.ID(uuid))

	return r.queryExtendedUser(schema.SkipSoftDelete(ctx), userQuery)
}

func (r *UserRepo) GetExtendedByEmail(ctx context.Context, email string) (*dto.ExtendedUser, error) {
	userQuery := r.client.User.Query().Where(user.Email(email))

	return r.queryExtendedUser(ctx, userQuery)
}

func (r *UserRepo) GetExtendedByVKId(ctx context.Context, vkId int64) (*dto.ExtendedUser, error) {
	userQuery := r.client.User.Query().Where(user.VkID(vkId))

	return r.queryExtendedUser(ctx, userQuery)
}

func (r *UserRepo) GetRolesByUuid(ctx context.Context, uuid string) (dto.Roles, error) {
	userRoles, err := r.client.User.Query().
		Where(user.ID(uuid)).
		QueryRoles().
		All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToRoleDTOs(userRoles), nil
}

func (r *UserRepo) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	exists, err := r.client.User.Query().Where(user.Email(email)).Exist(schema.SkipSoftDelete(ctx))
	if err != nil {
		return false, wrap(err)
	}

	return exists, nil
}

func (r *UserRepo) List(ctx context.Context) (dto.Users, error) {
	users, err := r.client.User.Query().All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToUserDTOs(users), nil
}

func (r *UserRepo) ListByUuids(ctx context.Context, uuids []string) (dto.Users, error) {
	users, err := r.client.User.Query().
		Where(user.IDIn(uuids...)).
		All(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToUserDTOs(users), nil
}

func (r *UserRepo) CreateWithPassword(ctx context.Context, dtm *dto.UserPasswordCreate) (*dto.User, error) {
	nextSerialNumber, err := r.getNextSerialNumber(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	usr, err := r.client.User.Create().
		AddRoleIDs(dtm.Roles...).
		SetEmail(dtm.Email).
		SetNillablePasswordHash(dtm.PasswordHash).
		SetSerialNumber(nextSerialNumber).
		SetFullname(fmt.Sprintf("Пользователь %d", nextSerialNumber)).
		SetUsername(fmt.Sprintf("%d", nextSerialNumber)).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToUserDTO(usr), nil
}

func (r *UserRepo) CreateWithService(ctx context.Context, dtm *dto.UserServiceCreate) (*dto.ExtendedUser, error) {
	nextSerialNumber, err := r.getNextSerialNumber(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	createdUsr, err := r.client.User.Create().
		AddRoleIDs(dtm.Roles...).
		SetNillableEmail(dtm.Email).
		SetNillableVkID(dtm.VkId).
		SetSerialNumber(nextSerialNumber).
		SetFullname(dtm.Fullname).
		SetUsername(fmt.Sprintf("%d", nextSerialNumber)).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	userQuery := r.client.User.Query().Where(user.ID(createdUsr.ID))

	return r.queryExtendedUser(ctx, userQuery)
}

func (r *UserRepo) Update(ctx context.Context, uuid string, dtm *dto.UserUpdate) (*dto.User, error) {
	usr, err := r.client.User.UpdateOneID(uuid).
		SetFullnameIfPresent(dtm.Fullname).
		SetUsernameIfPresent(dtm.Username).
		Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToUserDTO(usr), nil
}

func (r *UserRepo) VerifyEmail(ctx context.Context, uuid string) error {
	_, err := r.client.User.UpdateOneID(uuid).
		SetEmailVerified(true).
		Save(ctx)
	if err != nil {
		return wrap(err)
	}

	return nil
}

func (r *UserRepo) AddRole(ctx context.Context, uuid string, role roles.Type) error {
	err := r.client.User.UpdateOneID(uuid).
		AddRoleIDs(role).
		Exec(ctx)
	if err != nil {
		return wrap(err)
	}

	return nil
}

func (r *UserRepo) RemoveRole(ctx context.Context, uuid string, role roles.Type) error {
	err := r.client.User.UpdateOneID(uuid).
		RemoveRoleIDs(role).
		Exec(ctx)
	if err != nil {
		return wrap(err)
	}

	return nil
}

func (r *UserRepo) Delete(ctx context.Context, uuid string) error {
	err := r.client.User.DeleteOneID(uuid).Exec(ctx)
	if err != nil {
		return wrap(err)
	}

	return nil
}

func (r *UserRepo) Restore(ctx context.Context, uuid string) (*dto.User, error) {
	usr, err := r.client.User.UpdateOneID(uuid).ClearDeletedAt().Save(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToUserDTO(usr), nil
}

func (r *UserRepo) getNextSerialNumber(ctx context.Context) (uint, error) {
	query := `SELECT nextval('auth_users_serial_number_seq')`
	rows, err := r.client.QueryContext(ctx, query)
	if err != nil {
		return 0, wrap(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return 0, fmt.Errorf("вызов Next() вернул false")
	}

	var nextNumber uint
	err = rows.Scan(&nextNumber)
	if err != nil {
		return 0, wrap(err)
	}

	return nextNumber, nil
}

func (r *UserRepo) queryExtendedUser(ctx context.Context, userQuery *ent.UserQuery) (*dto.ExtendedUser, error) {
	usr, err := userQuery.
		WithRoles(func(q *ent.RoleQuery) {
			q.WithPermissions()
		}).
		Only(ctx)
	if err != nil {
		return nil, wrap(err)
	}

	return converters.ToExtendedUserDTO(usr), nil
}

func wrap(err error) error {
	if ent.IsConstraintError(err) {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			// Код 23505 - нарушение уникальности
			if pqErr.Code == "23505" {
				if strings.Contains(pqErr.Constraint, user.FieldEmail) {
					return fmt.Errorf("%w: пользователь с таким email уже существует", err_const.ErrUniqueConstraint)
				}
			}
		}
	}

	return utils.DefaultErrorWrapper(err)
}
