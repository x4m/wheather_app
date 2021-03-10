package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx"
	"time"
)

var (
	successCount = 0.0
	totalCount   = 0.0
)

func main() {
	conn, done := getConnection()
	if done {
		return
	}

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println("Availability ", successCount/totalCount)
		}

	}()

	for {
		totalCount++
		if conn != nil {
			_, err := conn.Exec(context.Background(), "insert into wheather values (now(), 'Yekaterinburg', random()*100 - 40)")
			if err != nil {
				fmt.Println(err)
				if conn.IsClosed() {
					conn = nil
				}
			} else {
				//fmt.Println("Wheather was written")
				successCount++
			}
		}

		if conn == nil {
			conn, done = getConnection()
			if !done {
				fmt.Println("No connection to DB")
			}
		}

		time.Sleep(time.Second / 2)
	}

	/*result := conn.QueryRow(context.Background(), "select pg_is_in_recovery()");
	var inRecovery bool
	err = result.Scan(&inRecovery)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("In recovery: ", inRecovery)*/
}

func getConnection() (*pgx.Conn, bool) {
	conn, err := pgx.Connect(context.Background(),
		"host=rc1c-men71f9ys74qk05u.mdb.yandexcloud.net,rc1b-gss4elrqataonjh7.mdb.yandexcloud.net "+
			"port=6432 "+
			"user=user1 "+
			"password=12345678 "+
			"database=db1 "+
			"target_session_attrs=read-write")
	if err != nil {
		fmt.Println(err)
		return nil, true
	}
	return conn, false
}
