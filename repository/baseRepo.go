package repository

import (
	"gorm.io/gorm"
	"management-platform/database"
)

type BaseRepo struct {
	Conn database.IMysql `inject:""`
}

func (b *BaseRepo) Create(value interface{}) error {
	return b.Conn.DB().Create(value).Error
}

func (b *BaseRepo) Save(value interface{}) error {
	return b.Conn.DB().Save(value).Error
}

func (b *BaseRepo) Updates(model interface{}, value interface{}) error {
	return b.Conn.DB().Model(model).Updates(value).Error
}

func (b *BaseRepo) DeleteByWhere(model, where interface{}) (count int64, err error) {
	db := b.Conn.DB().Where(where).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

func (b *BaseRepo) DeleteByID(model interface{}, id int) error {
	return b.Conn.DB().Where("id=?", id).Delete(model).Error
}

func (b *BaseRepo) DeleteByIDS(model interface{}, ids []int) (count int64, err error) {
	db := b.Conn.DB().Where("id in (?)", ids).Delete(model)
	err = db.Error
	if err != nil {
		return
	}
	count = db.RowsAffected
	return
}

func (b *BaseRepo) First(where interface{}, out interface{}, selects ...string) error {
	db := b.Conn.DB().Where(where)
	if len(selects) > 0 {
		for _, sel := range selects {
			db = db.Select(sel)
		}
	}
	return db.First(out).Error
}

func (b *BaseRepo) FirstByID(out interface{}, id int) error {
	return b.Conn.DB().First(out, id).Error
}

func (b *BaseRepo) Find(where interface{}, out interface{}, sel string, orders ...string) error {
	db := b.Conn.DB().Where(where)
	if sel != "" {
		db = db.Select(sel)
	}
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	return db.Find(out).Error
}

func (b *BaseRepo) GetPages(model interface{}, out interface{}, selects string, pageIndex int, pageSize int, totalCount *int64, where interface{}, orders ...string) error {
	db := b.Conn.DB().Model(model).Where(model)
	//where_ := make(map[string]interface{})
	//for k, v := range where.(map[string]interface{}) {
	//	where_[utils.Camel2Case(k)] = v
	//}
	db = db.Where(where)
	if len(orders) > 0 {
		for _, order := range orders {
			db = db.Order(order)
		}
	}
	err := db.Count(totalCount).Error
	if err != nil {
		return err
	}
	if *totalCount == 0 {
		return nil
	}
	return db.Select(selects).Offset((pageIndex - 1) * pageSize).Limit(pageSize).Find(out).Error
}

func (b *BaseRepo) PluckList(model, where interface{}, out interface{}, fieldName string) error {
	return b.Conn.DB().Model(model).Where(where).Pluck(fieldName, out).Error
}

func (b *BaseRepo) GetTransaction() *gorm.DB {
	return b.Conn.DB().Begin()
}
