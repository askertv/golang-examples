package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

func main() {
    db, err := sql.Open("mysql", "DATABASE:PASSWORD/DB")
    if err != nil {
        panic(err)
    }

    res, err := db.Query("SELECT * FROM `user`")
    if err != nil {
        panic(err)
    }

/*
select * from user \G
         id: 1
 created_at: 2020-11-16 01:15:47
 updated_at: 2020-11-16 01:15:47
       name: Owner
description:
      login: owner
   password: *****
   auth_key:
      email:
       role: admin
     active: 1
1 row in set (0.00 sec)
*/
// Users ...
type Users struct {
    id int64 `field:"id"`
	created_at string `field:"created_at"`
	updated_at string `field:"updated_at"`
	name string `field:"name"`
	description string `field:"description"`
    login string `field:"login"`
	password string `field:"password"`
	auth_key string `field:"auth_key"`
	email string `field:"email"`
    role string `field:"role"`
	active int64 `field:"active"`
}

    for res.Next() {
	    user := Users{}
        //fmt.Print(res.Scan())
		//fmt.Print(res.StructScan())
		//err := res.Scan(&bk.isbn, &bk.title, &bk.author, &bk.price)
		err := res.Scan(
		    &user.id,
			&user.created_at,
			&user.updated_at,
			&user.name,
			&user.description,
			&user.login,
			&user.password,
	        &user.auth_key,
	        &user.email,
			&user.role,
            &user.active)
		
		if err != nil {
            panic(err)
        }
		
		fmt.Print(user)
    }

    db.Close()
}
