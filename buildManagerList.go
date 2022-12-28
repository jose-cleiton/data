package data

type Manager struct {
	ID string
	FirstName string
	LastName string
}

func BuildManagerList(employees []Employee) []Manager {
	managerMap := make(map[string]Manager)
	for _, employee := range employees {
		for _, managerID := range employee.Managers {
			for _, otherEmployee := range employees {
				if otherEmployee.ID == managerID {
					if _, ok := managerMap[otherEmployee.ID]; !ok {
						managerMap[otherEmployee.ID] = Manager{ID: otherEmployee.ID, FirstName: otherEmployee.FirstName, LastName: otherEmployee.LastName}
					}
					break
				}
			}
		}
	}
	managerList := make([]Manager, 0, len(managerMap))
	for _, manager := range managerMap {
		managerList = append(managerList, manager)
	}
	return managerList
}
