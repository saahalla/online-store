package repository

import (
	"log"
	"online-store/common/dto"
	"online-store/common/logger"
	"online-store/shared/constant"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/jmoiron/sqlx"
)

type CartRepository interface {
	Add(cart dto.CartDB) error
	List() (dto.CartDBList, error)
	Get(params ParamSearchGetCart) (output dto.CartDB, err error)
	Delete(cartID int) (err error)
	Update(cartID int, cart dto.CartDB) (err error)
}

type cartRepository struct {
	db *sqlx.DB
}

func NewCartRepo(db *sqlx.DB) CartRepository {
	return &cartRepository{
		db: db,
	}
}

func (r *cartRepository) Add(cart dto.CartDB) error {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Insert(goqu.T(constant.TableCarts)).Rows(cart)

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err := r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryInsert(dataset, "Add Cart")
		return err
	}

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type ParamSearchGetCart struct {
	CartID int
	UserID int
}

func (r *cartRepository) Get(params ParamSearchGetCart) (output dto.CartDB, err error) {
	if params.CartID == 0 && params.UserID == 0 {
		return output, nil
	}

	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableCarts).As("c")).
		Select(
			goqu.I("c.id"),
			goqu.I("c.user_id"),
		)

	if params.CartID != 0 {
		dataset = dataset.Where(goqu.I("c.id").Eq(params.CartID))
	}

	if params.UserID != 0 {
		dataset = dataset.Where(goqu.I("c.user_id").Eq(params.UserID))
	}

	_, err = dataset.ScanStruct(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "Get Cart")
		return output, err
	}

	return output, nil
}

func (r *cartRepository) List() (output dto.CartDBList, err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableCarts).As("c")).
		Select(
			goqu.I("c.id"),
			goqu.I("c.user_id"),
		)

	err = dataset.ScanStructs(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "List Cart")
		return output, err
	}

	return output, nil
}

func (r *cartRepository) Delete(cartID int) (err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Delete(goqu.T(constant.TableCarts)).Where(goqu.I("id").Eq(cartID))

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err = r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryDelete(dataset, "Delete Cart")
		return err
	}

	return nil
}

func (r *cartRepository) Update(cartID int, cart dto.CartDB) (err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Update(goqu.T(constant.TableCarts)).
		Set(cart).
		Where(goqu.I("id").Eq(cartID))

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err = r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryUpdate(dataset, "Update Cart")
		return err
	}

	return nil
}
