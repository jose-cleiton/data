package data

import (
	"errors"
	
	"strings"
)

func GetRelatedEmployees(id string, zoo Zoo) ([]string, error) {
	var relatedEmployees []string

	exist := IsManager(id, zoo)
	if !exist {
		return nil, errors.New("O id inserido não é de uma pessoa colaboradora gerente!")
	}

	for _, employee := range zoo.Employees {
		idManagers := employee.Managers
		if ContainsString(idManagers, id) {
			completeName := []string{employee.FirstName, employee.LastName}
			relatedEmployees = append(relatedEmployees, strings.Join(completeName, " "))
		}
	}
	
	return relatedEmployees, nil
}
