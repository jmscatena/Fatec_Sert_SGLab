package services

import (
	"fmt"
	"log"

	"github.com/jmscatena/Fatec_Sert_SGLab/database"
	"github.com/jmscatena/Fatec_Sert_SGLab/utils"
)

func New[T utils.Tables](o utils.PersistenceHandler[T]) (int64, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return -1, err
	}
	recid, err := o.Create(db)
	fmt.Println("services:", err)
	if err != nil {
		//log.Fatalln(err)
		return 0, err
	}
	if recid != 0 {
		return recid, nil
	}
	return 0, nil
}

func Update[T utils.Tables](o utils.PersistenceHandler[T], uid uint64) (*T, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	rec, err := o.Update(db, uid)
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}
	return rec, nil
}

func Del[T utils.Tables](o utils.PersistenceHandler[T], uid uint64) (int64, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return -1, err
	}
	rec, err := o.Delete(db, uid)
	if err != nil {
		//log.Fatalln(err)
		return 0, err
	}
	return rec, nil
}

func Get[T utils.Tables](o utils.PersistenceHandler[T], uid uint64) (*T, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	rec, err := o.Find(db, uid)
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}
	return rec, nil
}

func GetAll[T utils.Tables](o utils.PersistenceHandler[T]) (*[]T, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	var rec *[]T
	rec, err = o.List(db)
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}
	return rec, nil
}

func GetBy[T utils.Tables](o utils.PersistenceHandler[T], param string, uid interface{}) (*[]T, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	rec, err := o.FindBy(db, param, uid)
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}
	return rec, nil
}
