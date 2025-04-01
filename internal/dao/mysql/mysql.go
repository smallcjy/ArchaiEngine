package mysql

import (
	"fmt"
	"game/internal/conf"
	"game/internal/logs"
	"game/internal/utils/errors"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	dB     *gorm.DB
	newDao *MysqlDao
)

func New(db *gorm.DB) *MysqlDao {
	return &MysqlDao{
		Db: db,
	}
}

type MysqlDao struct {
	Db *gorm.DB
}

func MysqlStart() *errors.Error {
	dia := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.GetServerConf().Mysql.UserName, conf.GetServerConf().Mysql.PassWord, conf.GetServerConf().Mysql.Addr, conf.GetServerConf().Mysql.DataBase)

	var err error
	dB, err = gorm.Open(mysql.Open(dia), &gorm.Config{
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
		PrepareStmt:                              true,
	})

	if err != nil {
		return errors.ErrDb
	}
	sqlDb, _ := dB.DB()
	sqlDb.SetMaxIdleConns(10)
	sqlDb.SetMaxOpenConns(100)
	sqlDb.SetConnMaxLifetime(time.Hour)

	newDao = &MysqlDao{
		Db: dB,
	}
	logs.Info("MysqlStart OK... ", zap.String("conn", dia))
	return nil
}

func MysqlClose() *errors.Error {
	if dB == nil {
		return nil
	}
	sqlDb, err := dB.DB()
	if err != nil {
		logs.Error("MysqlClose err", zap.Error(err))
		return errors.ErrDb
	}
	err = sqlDb.Close()
	if err != nil {
		logs.Error("MysqlClose err", zap.Error(err))
		return errors.ErrDb
	}
	logs.Info("MysqlClose OK...")
	return nil
}

func AutoCreateTable() {
	// TODO:
}

func GetDao() *MysqlDao {
	return newDao
}