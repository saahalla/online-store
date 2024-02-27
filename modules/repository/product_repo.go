package repository

import (
	"context"
	"log"
	"online-store/modules/dto"
	"online-store/shared/constant"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	"github.com/jmoiron/sqlx"
)

type ProductRepository interface {
	Add(product dto.ProductDB) error
	List() (dto.ProductDBList, error)
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

	_, err := r.db.ExecContext(context.Background(),
		`INSERT INTO products(
			product_name, 
			price, 
			stock, 
			image, 
			created_at, 
			created_by, 
			modified_at, 
			modified_by) 
		VALUES(?,?,?,?,?,?,?,?)`,
		product.ProductName,
		product.Price,
		product.Stock,
		product.Image,
		product.CreatedAt,
		product.CreatedBy,
		product.ModifiedAt,
		product.ModifiedBy,
	)

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
		).Where(goqu.I("p.id").Eq(productID))

	_, err = dataset.ScanStruct(&output)

	if err != nil {
		LogQuerySelect(dataset, "Get Product")
		return output, err
	}

	return output, nil
}

func (r *productRepository) List() (output dto.ProductDBList, err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableProducts).As("p")).
		Select(
			goqu.I("p.id"),
			goqu.I("p.product_name"),
			goqu.I("p.stock"),
			goqu.I("p.price"),
			goqu.COALESCE(goqu.I("p.image"), "").As("image"),
		)

	err = dataset.ScanStructs(&output)

	if err != nil {
		LogQuerySelect(dataset, "List Product")
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
		LogQueryDelete(dataset, "Delete Product")
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
		LogQueryUpdate(dataset, "Update Product")
		return err
	}

	return nil
}

func LogQuerySelect(dataset *goqu.SelectDataset, name string) {
	query, _, _ := dataset.Prepared(false).ToSQL()
	log.Printf("%v: %v", name, query)
}

func LogQueryDelete(dataset *goqu.DeleteDataset, name string) {
	query, _, _ := dataset.Prepared(false).ToSQL()
	log.Printf("%v: %v", name, query)
}

func LogQueryUpdate(dataset *goqu.UpdateDataset, name string) {
	query, _, _ := dataset.Prepared(false).ToSQL()
	log.Printf("%v: %v", name, query)
}
