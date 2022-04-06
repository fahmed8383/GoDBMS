func ParseModifyTable(query string) (*ParserStructs.ModifyTableStatement, error) {
	
	query = strings.Trim(query, " ;")

	querySplit := strings.Split(query, " ")

	DataType := ""

	//check if statement has correct number of arguments
	if len(querySplit) != 6 {
		return nil, errors.New("Invalid alter table statement")
	}

	//check if valid syntax/ which type of alter table statement it is
	if querySplit[3] == "add" {
		TableName := querySplit[2]
		ColumnName := querySplit[4]
		DataType = querySplit[5]
	} else if querySplit[3] == "drop" && querySplit[4] == "column"{
		TableName := querySplit[2]
		ColumnName := querySplit[5]
	} else {
		return nil, errors.New("Invalid alter table statement")
	}

	return &ParserStructs.ModifyTableStatement{TableName, ColumnName, DataType}, nil
}