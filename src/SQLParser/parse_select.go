package SQLParser

import (
	"GoDBMS/Database"
	"strings"
	"errors"
)

func parseSelectQuery(query string) (*Database.SelectStatement, error) {
	querySplit := strings.Split(query, " ")
	if len(querySplit) != 4 {
		return nil, errors.New("Invalid number of parameters for select statement")
	}
	if querySplit[1] != "*" || querySplit[2] != "from" {
		return nil, errors.New("Invalid syntax for select statement")
	}

	name := querySplit[3]
	
	return &Database.SelectStatement{name}, nil
}

