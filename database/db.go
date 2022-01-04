package database

import (
	"errors"
	"fmt"
	"main/domain"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const USERNAME = "root"
const PASSWORD = "root1234"
const DATABASE = "my_db"
const HOST = "127.0.0.1"
const PORT = "3306"

type Database struct {
	DbConn       *gorm.DB
	MemberEntity domain.Member
}

var instance *Database
var once sync.Once

func GetDbInstance() *Database {
	once.Do(func() {
		instance = &Database{}
		db_infp := USERNAME + ":" + PASSWORD + "@tcp" + "(" + HOST + ":" + PORT + ")/" + DATABASE + "?" + "parseTime=true&loc=Local"
		fmt.Println("db info : ", db_infp)
		db, err := gorm.Open(mysql.Open(db_infp), &gorm.Config{})
		if err != nil {
			instance.DbConn = nil
		}

		db.AutoMigrate(&domain.Member{})

		instance.DbConn = db
	})

	return instance
}

func (database *Database) Save(args interface{}) error {
	switch args.(type) {
	case *domain.Member:
		member := args.(*domain.Member)
		return database.SaveMember(member)
	case nil:
		return database.SaveNothing()
	default:
		return errors.New("new error")
	}
}

func (database *Database) SaveMember(member *domain.Member) error {
	err := database.DbConn.Create(&member).Error
	if err != nil {
		return err
	}

	return nil
}

func (database *Database) SaveNothing() error {
	return errors.New("save nothing")
}

func (database *Database) FindAll(args interface{}) error {
	switch args.(type) {
	case *[]domain.Member:
		member := args.(*[]domain.Member)
		return database.FindAllMember(member)
	default:
		return errors.New("new error")
	}
}

func (database *Database) FindAllMember(memberList *[]domain.Member) error {
	err := database.DbConn.Find(memberList).Error
	if err != nil {
		return err
	}

	return nil
}
