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

type CategoryRepository interface {
	Add(category dto.CategoryDB) error
	List() (dto.CategoryDBList, error)
	Get(categoryID int) (output dto.CategoryDB, err error)
	Delete(categoryID int) (err error)
	Update(categoryID int, category dto.CategoryDB) (err error)
}

type categoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepo(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (r *categoryRepository) Add(category dto.CategoryDB) error {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Insert(goqu.T(constant.TableCategories)).Rows(category)

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err := r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryInsert(dataset, "Add Category")
		return err
	}

	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (r *categoryRepository) Get(categoryID int) (output dto.CategoryDB, err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableCategories).As("p")).
		Select(
			goqu.I("p.id"),
			goqu.I("p.category_name"),
		).Where(goqu.I("p.id").Eq(categoryID))

	_, err = dataset.ScanStruct(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "Get Category")
		return output, err
	}

	return output, nil
}

func (r *categoryRepository) List() (output dto.CategoryDBList, err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.From(goqu.T(constant.TableCategories).As("p")).
		Select(
			goqu.I("p.id"),
			goqu.I("p.category_name"),
		)

	err = dataset.ScanStructs(&output)

	if err != nil {
		logger.LogQuerySelect(dataset, "List Category")
		return output, err
	}

	return output, nil
}

func (r *categoryRepository) Delete(categoryID int) (err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Delete(goqu.T(constant.TableCategories)).Where(goqu.I("id").Eq(categoryID))

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err = r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryDelete(dataset, "Delete Category")
		return err
	}

	return nil
}

func (r *categoryRepository) Update(categoryID int, category dto.CategoryDB) (err error) {
	dialect := goqu.New(r.db.DriverName(), r.db)

	dataset := dialect.Update(goqu.T(constant.TableCategories)).
		Set(category).
		Where(goqu.I("id").Eq(categoryID))

	query, values, _ := dataset.Prepared(true).ToSQL()

	_, err = r.db.Queryx(query, values...)
	if err != nil {
		logger.LogQueryUpdate(dataset, "Update Category")
		return err
	}

	return nil
}
