package oracle

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// SayHello calls the stored function 'booktest.say_hello(nvarchar2) nvarchar2' on db.
func SayHello(ctx context.Context, db DB, name string) (string, error) {
	// call booktest.say_hello
	const sqlstr = `SELECT booktest.say_hello(:1) FROM dual`
	// run
	var r0 string
	logf(sqlstr, name)
	if err := db.QueryRowContext(ctx, sqlstr, name).Scan(&r0); err != nil {
		return "", logerror(err)
	}
	return r0, nil
}