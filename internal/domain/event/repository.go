package event

import (
	"fmt"
	"github.com/test_server/pkg/dbsettings"
	"github.com/upper/db/v4/adapter/postgresql"
	"log"
)

type Repository interface {
	FindAll() ([]Event, error)
	FindOne(id int64) (*Event, error)
}

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() ([]Event, error) {
	//Підключаємося до БД
	sess, err := postgresql.Open(dbsettings.Settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	//Відкладенне відключення до БД
	defer sess.Close()
	//Вибірка даних
	res, err := sess.SQL().Query("SELECT * FROM events")
	if err != nil {
		fmt.Println(err.Error())
	}
	var events []Event
	//Обнуляємо список
	events = []Event{}
	//В циклі додаємо всі об'єкти в слайс
	for res.Next() {
		var event Event
		err = res.Scan(&event.Id, &event.Name)
		if err != nil {
			fmt.Println(err.Error())
		}
		events = append(events, event)
	}
	return events, err
}

func (r *repository) FindOne(id int64) (*Event, error) {
	//Підключаємося до БД
	sess, err := postgresql.Open(dbsettings.Settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	//Відкладенне відключення до БД
	defer sess.Close()
	//Вибірка даних

	res, err := sess.SQL().Query(fmt.Sprintf("SELECT * FROM events WHERE id = %d", id))
	if err != nil {
		fmt.Println(err.Error())
	}
	var event Event
	for res.Next() {
		err = res.Scan(&event.Id, &event.Name)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
	return &event, err
}
