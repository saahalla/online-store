package repository

import (
	"log"
	"online-store/common/dto"
	"online-store/common/logger"
	"online-store/shared/constant"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
)

type UserRepository interface {
	Add(dto.UserDB) error
	Get(paramSearch ParamGetUser) (output dto.UserDB, err error)
}

type userRepository struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) Add(user dto.UserDB) error {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Insert(goqu.T(constant.TableUsers)).Rows(user)

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err := r.db.Queryx(query, values...)
	if err != nil {
		log.Println(err)
		logger.LogQueryInsert(dataset, "Add User")
		return err
	}

	return nil
}

type ParamGetUser struct {
	UserID   int
	Username string
	Email    string
}

func (r *userRepository) Get(paramSearch ParamGetUser) (output dto.UserDB, err error) {

	if paramSearch.UserID == 0 && paramSearch.Username == "" && paramSearch.Email == "" {
		return output, nil
	}

	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableUsers).As("u")).
		Select(
			goqu.I("u.id"),
			goqu.I("u.username"),
			goqu.I("u.password"),
			goqu.I("u.email"),
			goqu.COALESCE(goqu.I("u.phone"), "").As("phone"),
			goqu.I("u.user_role_id"),
		)

	if paramSearch.UserID != 0 {
		dataset = dataset.Where(goqu.I("u.id").Eq(paramSearch.UserID))
	}

	if paramSearch.Email != "" {
		dataset = dataset.Where(goqu.I("u.email").Eq(paramSearch.Email))
	}

	if paramSearch.Username != "" {
		dataset = dataset.Where(goqu.I("u.username").Eq(paramSearch.Username))
	}

	_, err = dataset.ScanStruct(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "Get User")
		return output, err
	}

	return output, nil
}
