package data

func GetEmployeeByName(employees []Employee, firstName string, lastName string) Employee {
	if len(employees) == 0 {
			return Employee{}
	}

	for _, employee := range employees {
			if employee.FirstName == firstName || employee.LastName == lastName {
					return employee
			}
	}

	return Employee{}
}
