package crud_proto

import (
	"context"
	"fmt"

	"github.com/mohammad-siraj/crud_grpc/database"
	"github.com/mohammad-siraj/crud_grpc/entities"
)

type Server struct {
	UnimplementedCarInfoServer
}

var db database.Database

func Init() error {
	var err error
	db, err = database.NewDatabase()
	return err
}

func (s *Server) GetCarInfo(ctx context.Context, carid *Id) (*Car, error) {
	car, err := db.Getdata_db(int(carid.Carid))
	if err != nil {
		fmt.Printf("enteted get car info")
	}
	return &Car{Model: car.Model, Make: car.Make, Year: car.Year}, err
}
func (s *Server) UpdateInfo(ctx context.Context, carcomplete *CarComplete) (*Car, error) {
	car, err := db.Updatedata_db(entities.Car{Model: carcomplete.Querycar.Model, Make: carcomplete.Querycar.Make, Year: carcomplete.Querycar.Year}, int(carcomplete.Queryid.Carid))
	if err != nil {
		fmt.Printf("enteted get car info")
	}
	return &Car{Model: car.Model, Make: car.Make, Year: car.Year}, err
}
func (s *Server) DeleteInfo(ctx context.Context, carid *Id) (*Car, error) {
	car, err := db.Deletedata_db(int(carid.Carid))
	if err != nil {
		fmt.Printf("enteted get car info")
	}
	return &Car{Model: car.Model, Make: car.Make, Year: car.Year}, err
}
func (s *Server) CreateInfo(ctx context.Context, carcomplete *Car) (*Car, error) {
	car := entities.Car{
		Model: carcomplete.Model,
		Make:  carcomplete.Make,
		Year:  carcomplete.Year}
	err := db.Postdata_db(car)
	if err != nil {
		fmt.Printf("enteted get car info")
	}
	return &Car{Model: car.Model, Make: car.Make, Year: car.Year}, err
}
