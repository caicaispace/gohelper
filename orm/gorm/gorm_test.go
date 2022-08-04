package gorm_test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/caicaispace/gohelper/orm/gorm"

	"github.com/caicaispace/gohelper/setting"

	orm "gorm.io/gorm"
)

const connName = "default"

var config = &setting.DBSetting{
	Username: "root",
	Password: "123456",
	Host:     "127.0.0.1",
	DbName:   "gohelper",
}

type Menu struct {
	gorm.BaseModel
	Name  string `json:"name" gorm:"column:name;type:varchar(255);not null;default:''"`
	IsDel uint8  `json:"is_del" gorm:"default:0"`
}

func TestAutoMigrate(t *testing.T) {
	db := gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
	err := db.AutoMigrate(&Menu{}) // auto generate table ddl
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

func TestInsterData(t *testing.T) {
	db := gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
	var menu Menu
	menu.Name = "test"
	ret := db.Table("menus").Create(&menu)
	t.Logf("%+v", ret)
	t.Logf("%+v", menu)
}

func TestFindData(t *testing.T) {
	db := gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
	var results []*Menu
	ret := db.Table("menus").Where("id > ?", 0).Find(&results)
	t.Log(ret)
	for _, result := range results {
		t.Logf("%+v", result)
	}
}

func TestFindScan(t *testing.T) {
	db := gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
	var results []*Menu
	db.Raw("select * from menus where id > ?", 0).Scan(&results)
	for _, result := range results {
		// fmt.Println(result)
		t.Logf("%+v", result)
	}
}

func TestFindInBatches(t *testing.T) {
	db := gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
	var results []*Menu
	ret := db.Table("menus").Where("id > ?", 0).FindInBatches(&results, 5, func(tx *orm.DB, batch int) error {
		for _, result := range results {
			t.Logf("%+v", result)
			// batch processing found records
		}
		// tx.Save(&results)
		// tx.RowsAffected // number of records in this batch
		// batch // Batch 1, 2, 3
		// returns error will stop future batches
		return nil
	})
	t.Logf("%+v", ret)
}

func TestUpdate(t *testing.T) {
	db := gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
	var menu Menu
	menu.Name = "test"
	ret := db.Table("menus").Create(&menu)
	t.Logf("%+v", ret)
	t.Logf("%+v", menu)
	menu.Name = "test2"
	ret = db.Table("menus").Where("id = ?", 1).Updates(&menu)
	t.Logf("%+v", ret)
	t.Logf("%+v", menu)
}

func TestSoftDelete(t *testing.T) {
	db := gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
	timeNow := time.Now()
	var menu Menu
	menu.Name = "test3"
	menu.IsDel = 1
	menu.DeletedAt = &timeNow
	ret := db.Table("menus").Where("id > ?", 1).Updates(&menu)
	t.Logf("%+v", ret)
	t.Logf("%+v", menu)
}

func TestDelete(t *testing.T) {
	db := gorm.New().AddConnWithConfig(config, connName).GetDB(connName)
	var menu Menu
	menu.Name = "test"
	ret := db.Table("menus").Create(&menu)
	t.Logf("%+v", ret)
	t.Logf("%+v", menu)
	ret = db.Table("menus").Where("id = ?", menu.ID).Delete(&menu)
	t.Logf("%+v", ret)
}
