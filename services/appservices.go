package services

import (
	"fmt"
	"github.com/google/uuid"
	"log"

	"github.com/jmscatena/Fatec_Sert_SGLab/database"
	"github.com/jmscatena/Fatec_Sert_SGLab/handlers"
)

func New[T handlers.Tables](o handlers.PersistenceHandler[T]) (uuid.UUID, error) {
	/* metodo com devolucao do UUID */
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return uuid.Nil, err
	}
	recid, err := o.Create(db)
	fmt.Println("services:", err)
	if err != nil {
		log.Fatalln(err)
		return uuid.Nil, err
	}
	return recid, nil
}

func Update[T handlers.Tables](o handlers.PersistenceHandler[T], uid uuid.UUID) (*T, error) {
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

func Del[T handlers.Tables](o handlers.PersistenceHandler[T], uid uuid.UUID) (int64, error) {
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

/*func Get[T handlers.Tables](o handlers.PersistenceHandler[T], uid uuid.UUID) (*T, error) {
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
*/

func GetAll[T handlers.Tables](o handlers.PersistenceHandler[T]) (*[]T, error) {
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
func Get[T handlers.Tables](o handlers.PersistenceHandler[T], param string, values string) (*T, error) {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	rec, err := o.Find(db, param, values)
	if err != nil {
		//log.Fatalln(err)
		return nil, err
	}
	return rec, nil
}
