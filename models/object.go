package models

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

var (
	Objects map[string]*Object
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	database = "user"
)

type Object struct {
	FirstName   string `json:"FirstName"`
	Lastname    string `json:"LastName"`
	Phonenumber string `json:"Phonenumber"`
	Email       string `json:"Email"`
	Password    string `json:"Password"`
	DateOfBirth string `json:"DateOfBirth"`
}

func init() {
	Objects = make(map[string]*Object)

}

func AddOne(object Object) (ObjectId string) {
	var Errormsg = ""
	//initialize connection database
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, database)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	// database hanling
	datatable := "SELECT * FROM information_schema.tables WHERE table_name = 'user_information';"
	fmt.Println(datatable)
	table, _ := db.Exec(datatable)
	tresult, _ := (table.RowsAffected())
	fmt.Println(tresult)

	if tresult < 1 {
		Createtable := "CREATE TABLE user_information(Firstname VARCHAR(255),Lastname VARCHAR (255),Phonenumber VARCHAR (255),Email VARCHAR (255) PRIMARY KEY,password VARCHAR (255) NOT NULL,Birthdate VARCHAR(255));"
		db.Exec(Createtable)
		fmt.Println("Table created")

		// validation checking
		rep := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
		phone := rep.MatchString(object.Phonenumber)
		red := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		bdate := red.MatchString(object.DateOfBirth)
		rem := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		email := rem.MatchString(object.Email)
		pass, _ := HashPassword(object.Password)
		fmt.Println(phone, bdate, email, pass)
		if phone && bdate && email {

		} else {
			if !phone {
				Errormsg = Errormsg + "--Phone Number is Invalid--  "
			}
			if !bdate {
				Errormsg = Errormsg + "--Bithdate is Invalid--  "
			}
			if !email {
				Errormsg = Errormsg + "--Email is Invalid--  "
			}
			return Errormsg
		}

	} else {
		fmt.Println("Has table")

		// validation checking
		rep := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
		phone := rep.MatchString(object.Phonenumber)
		red := regexp.MustCompile("(0?[1-9]|[12][0-9]|3[01])/(0?[1-9]|1[012])/((19|20)\\d\\d)")
		bdate := red.MatchString(object.DateOfBirth)
		rem := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
		email := rem.MatchString(object.Email)
		pass, _ := HashPassword(object.Password)
		fmt.Println(phone, bdate, email, pass)
		if phone && bdate && email {
			Errormsg = "successfullt Inserted"
		} else {
			if !phone {
				Errormsg = Errormsg + "--Phone Number is Invalid--  "
			}
			if !bdate {
				Errormsg = Errormsg + "--Bithdate is Invalid--  "
			}
			if !email {
				Errormsg = Errormsg + "--Email is Invalid--  "
			}
			return Errormsg
		}
	}
	// database handling

	Objects[object.FirstName] = &object
	Errormsg = "successfullt Inserted"
	return Errormsg
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}
func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func GetOne(ObjectId string) (object *Object, err error) {
	if v, ok := Objects[ObjectId]; ok {
		return v, nil
	}
	return nil, errors.New("ObjectId Not Exist")
}

func GetAll() map[string]*Object {
	return Objects
}

// func Update(ObjectId string, Score int64) (err error) {
// 	if v, ok := Objects[ObjectId]; ok {
// 		v.Score = Score
// 		return nil
// 	}
// 	return errors.New("ObjectId Not Exist")
// }

func Delete(ObjectId string) {
	delete(Objects, ObjectId)
}
