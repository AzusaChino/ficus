package my_service

import (
	"github.com/azusachino/ficus/pkg/mydb"
	"log"
)

type Hello struct {
	Name string `json:"name"`
	Msg  string `json:"msg"`
}

func GetMsg(name string) string {
	var hello Hello
	row := mydb.DB.QueryRow(`select name, msg from hello where name = ?`, name)
	// scan use select order, must match
	err := row.Scan(&hello.Name, &hello.Msg)
	if err != nil {
		log.Println(err)
	}
	return hello.Msg
}
