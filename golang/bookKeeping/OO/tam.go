package OO

import (
	"bookKeeping/db"
	"encoding/json"
	"errors"
	"fmt"
)

type Tam struct {
	Rtx          string
	Tam_group    string
	Assosciate_1 string
	Assosciate_2 string
}

//var DB = db.DB

func GetAllTam() {
	var tam = make([]Tam, 0, 0)

	rows, e := db.DB.Query("select * from tam;")
	if e == nil {
		errors.New("query incur error")
	}
	for rows.Next() {
		var tempTam Tam
		e := rows.Scan(&tempTam.Rtx, &tempTam.Tam_group, &tempTam.Assosciate_1, &tempTam.Assosciate_2)
		if e != nil {
			fmt.Println(json.Marshal(tam))
		} else {
			tam = append(tam, tempTam)
		}
	}

	for i, v := range tam {
		fmt.Printf("%d. [%s]: %s -> %s + %s\n", i, v.Tam_group, v.Rtx, v.Assosciate_1, v.Assosciate_2)
	}

}
