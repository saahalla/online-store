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

type CartItemRepository interface {
	Add(cartItem dto.CartItemDB) error
	List(params ParamSearchGetCartItemList) (output dto.CartItemDBList, err error)
	Get(params ParamSearchGetCartItem) (output dto.CartItemDB, err error)
	DetailList(params ParamSearchDetailCartItem) (output dto.CartItemDataList, err error)
	Delete(cartItemID int) (err error)
	Update(cartItemID int, cartItem dto.CartItemDB) (err error)
}

type cartItemRepository struct {
	db *sqlx.DB
}

func NewCartItemRepo(db *sqlx.DB) CartItemRepository {
	return &cartItemRepository{
		db: db,
	}
}

func (r *cartItemRepository) Add(cartItem dto.CartItemDB) error {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Insert(goqu.T(constant.TableCartItems)).Rows(cartItem)

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err := r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryInsert(dataset, "Add CartItem")
		return err
	}

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

type ParamSearchGetCartItem struct {
	CartItemID int
	CartID     int
	ProductID  int
}

func (r *cartItemRepository) Get(params ParamSearchGetCartItem) (output dto.CartItemDB, err error) {
	if params.CartItemID == 0 && params.CartID == 0 && params.ProductID == 0 {
		return output, nil
	}

	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableCartItems).As("c")).
		Select(
			goqu.I("c.id"),
			goqu.I("c.cart_id"),
			goqu.I("c.product_id"),
			goqu.I("c.qty"),
		)

	if params.CartItemID != 0 {
		dataset = dataset.Where(goqu.I("c.id").Eq(params.CartItemID))
	}

	if params.CartID != 0 {
		dataset = dataset.Where(goqu.I("c.cart_id").Eq(params.CartID))
	}

	if params.ProductID != 0 {
		dataset = dataset.Where(goqu.I("c.product_id").Eq(params.ProductID))
	}

	_, err = dataset.ScanStruct(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "Get CartItem")
		return output, err
	}

	return output, nil
}

type ParamSearchGetCartItemList struct {
	CartItemID int
	CartID     int
	ProductID  int
}

func (r *cartItemRepository) List(params ParamSearchGetCartItemList) (output dto.CartItemDBList, err error) {
	if params.CartItemID == 0 && params.CartID == 0 && params.ProductID == 0 {
		return output, nil
	}

	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableCartItems).As("c")).
		Select(
			goqu.I("c.id"),
			goqu.I("c.cart_id"),
			goqu.I("c.product_id"),
			goqu.I("c.qty"),
		)

	if params.CartItemID != 0 {
		dataset = dataset.Where(goqu.I("c.id").Eq(params.CartItemID))
	}

	if params.CartID != 0 {
		dataset = dataset.Where(goqu.I("c.cart_id").Eq(params.CartID))
	}

	if params.ProductID != 0 {
		dataset = dataset.Where(goqu.I("c.product_id").Eq(params.ProductID))
	}

	err = dataset.ScanStructs(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "List CartItem")
		return output, err
	}

	return output, nil
}

type ParamSearchDetailCartItem struct {
	CartItemID int
	CartID     int
	ProductID  int
}

func (r *cartItemRepository) DetailList(params ParamSearchDetailCartItem) (output dto.CartItemDataList, err error) {
	if params.CartItemID == 0 && params.CartID == 0 && params.ProductID == 0 {
		return output, nil
	}

	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.
		Select(
			goqu.I("c.id"),
			goqu.I("c.product_id"),
			goqu.I("p.product_name"),
			goqu.I("c.qty"),
		).From(goqu.T(constant.TableCartItems).As("c")).
		Join(goqu.T(constant.TableProducts).As("p"), goqu.On(
			goqu.I("p.id").Eq(goqu.I("c.product_id")),
		))

	if params.CartItemID != 0 {
		dataset = dataset.Where(goqu.I("c.id").Eq(params.CartItemID))
	}

	if params.CartID != 0 {
		dataset = dataset.Where(goqu.I("c.cart_id").Eq(params.CartID))
	}

	if params.ProductID != 0 {
		dataset = dataset.Where(goqu.I("c.product_id").Eq(params.ProductID))
	}

	err = dataset.ScanStructs(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "List CartItem")
		return output, err
	}

	return output, nil
}

func (r *cartItemRepository) Delete(cartItemID int) (err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Delete(goqu.T(constant.TableCartItems)).Where(goqu.I("id").Eq(cartItemID))

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err = r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryDelete(dataset, "Delete CartItem")
		return err
	}

	return nil
}

func (r *cartItemRepository) Update(cartItemID int, cartItem dto.CartItemDB) (err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Update(goqu.T(constant.TableCartItems)).
		Set(cartItem).
		Where(goqu.I("id").Eq(cartItemID))

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err = r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryUpdate(dataset, "Update CartItem")
		return err
	}

	return nil
}
