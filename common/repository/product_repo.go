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

type ProductRepository interface {
	Add(product dto.ProductDB) error
	List(paramSearch ParamSearchProductList) (output dto.ProductDBList, err error)
	Get(productID int) (output dto.ProductDB, err error)
	Delete(productID int) (err error)
	Update(productID int, product dto.ProductDB) (err error)
}

type productRepository struct {
	db *sqlx.DB
}

func NewProductRepo(db *sqlx.DB) ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) Add(product dto.ProductDB) error {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Insert(goqu.T(constant.TableProducts)).Rows(product)

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err := r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryInsert(dataset, "Add Product")
		return err
	}

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *productRepository) Get(productID int) (output dto.ProductDB, err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableProducts).As("p")).
		Select(
			goqu.I("p.id"),
			goqu.I("p.product_name"),
			goqu.I("p.stock"),
			goqu.I("p.price"),
			goqu.COALESCE(goqu.I("p.image"), "").As("image"),
			goqu.COALESCE(goqu.I("p.category_id"), 0).As("category_id"),
		).Where(goqu.I("p.id").Eq(productID))

	_, err = dataset.ScanStruct(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "Get Product")
		return output, err
	}

	return output, nil
}

type ParamSearchProductList struct {
	ProductIDList []int
	ProductName   string
	CategoryID    int
}

func (r *productRepository) List(paramSearch ParamSearchProductList) (output dto.ProductDBList, err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableProducts).As("p")).
		Select(
			goqu.I("p.id"),
			goqu.I("p.product_name"),
			goqu.I("p.stock"),
			goqu.I("p.price"),
			goqu.COALESCE(goqu.I("p.image"), "").As("image"),
			goqu.COALESCE(goqu.I("p.category_id"), 0).As("category_id"),
		)

	if len(paramSearch.ProductIDList) > 0 {
		dataset = dataset.Where(
			goqu.I("p.id").In(paramSearch.ProductIDList),
		)
	}

	if paramSearch.ProductName != "" {
		dataset = dataset.Where(
			goqu.I("p.product_name").ILike("%" + paramSearch.ProductName + "%"),
		)
	}

	if paramSearch.CategoryID != 0 {
		dataset = dataset.Where(
			goqu.I("p.category_id").Eq(paramSearch.CategoryID),
		)
	}

	err = dataset.ScanStructs(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "List Product")
		return output, err
	}

	return output, nil
}

func (r *productRepository) Delete(productID int) (err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Delete(goqu.T(constant.TableProducts)).Where(goqu.I("id").Eq(productID))

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err = r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryDelete(dataset, "Delete Product")
		return err
	}

	return nil
}

func (r *productRepository) Update(productID int, product dto.ProductDB) (err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Update(goqu.T(constant.TableProducts)).
		Set(product).
		Where(goqu.I("id").Eq(productID))

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err = r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryUpdate(dataset, "Update Product")
		return err
	}

	return nil
}
