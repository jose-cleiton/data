package zoo



type Resident struct {
  Name string
  Sex  string
  Age  int
}

type Species struct {
  ID          string
  Name        string
  Popularity  int
  Location    string
  Availability []string
  Residents   []Resident
}

type Employee struct {
  ID          string
  FirstName   string
  LastName    string
  Managers    []string
  ResponsibleFor []string
}

type Hours struct {
  Open  int
  Close int
}

type Prices struct {
  Adult  float64
  Senior float64
  Child  float64
}

type Zoo struct {
  Species    []Species
  Employees  []Employee
  Hours      map[string]Hours
  Prices     Prices
}




var  Data = Zoo{
		Species: []Species{
			{
				ID:          "0938aa23-f153-4937-9f88-4858b24d6bce",
				Name:        "lions",
				Popularity:  4,
				Location:    "NE",
				Availability: []string{"Tuesday", "Thursday", "Saturday", "Sunday"},
				Residents: []Resident{
					{Name: "Zena", Sex: "female", Age: 12},
					{Name: "Maxwell", Sex: "male", Age: 15},
					{Name: "Faustino", Sex: "male", Age: 7},
					{Name: "Dee", Sex: "female", Age: 14},
				},
			},
			{
				ID:          "e8481c1d-42ea-4610-8e11-1752cfc05a46",
				Name:        "tigers",
				Popularity:  5,
				Location:    "NW",
				Availability: []string{"Wednesday", "Friday", "Saturday", "Tuesday"},
				Residents: []Resident{
					{Name: "Shu", Sex: "female", Age: 19},
					{Name: "Esther", Sex: "female", Age: 17},
				},
			},

			{
				ID: "ef3778eb-2844-4c7c-b66c-f432073e1c6b",
				Name: "penguins",
				Popularity:  4,
				Location:    "SE",
				Availability: []string{"Tuesday", "Wednesday", "Sunday"},
				Residents: []Resident{
					{Name: "Joe",  Sex: "male", Age: 10},
					{Name: "Tad",  Sex: "male", Age: 12},
					{Name: "Keri", Sex: "female", Age: 2},
					{Name: "Nicholas", Sex: "male", Age: 2},

				},

			},
			{
				ID: "ottersId",
				Name: "otters",
				Popularity:  4,
				Location:    "SE",
				Availability: []string{"Friday", "Thursday", "Wednesday", "Saturday"},
				Residents: []Resident{
					{Name: "Neville",  Sex: "male", Age: 9},
					{Name: "Lloyd",  Sex: "male", Age: 8},
					{Name: "Mercedes", Sex: "female", Age: 9},
					{Name: "Margherita", Sex: "female", Age: 10},

				},

			},
			{
				ID: "frogsId",
				Name: "frogs",
				Popularity:  2,
				Location:    "SW",
				Availability: []string{"Saturday", "Friday", "Thursday", "Wednesday"},
				Residents: []Resident{
					{Name: "Cathey",  Sex: "female", Age: 3},
					{Name: "Annice",  Sex: "female", Age: 2},
					
				},

			},
			{
				ID: "snakesId",
				Name: "snakes",
				Popularity:  3,
				Location:    "SW",
				Availability: []string{"Sunday", "Saturday", "Friday", "Thursday"},
				Residents: []Resident{
					{Name: "Paulette",  Sex: "female", Age: 5},
					{Name: "Bill",  Sex: "male", Age: 6},
					
				},

			},
			{
				ID: "elephantsId",
				Name: "elephants",
				Popularity:  5,
				Location:    "NW",
				Availability: []string{"Friday", "Saturday", "Sunday", "Tuesday"},
				Residents: []Resident{
					{Name: "Ilana",  Sex: "female", Age: 11},
					{Name: "Orval",  Sex: "male", Age: 15},
					{Name: "Bea",  Sex: "female", Age: 12},
					{Name: "Jefferson",  Sex: "male", Age: 4},
					
				},

			},
			{
				ID: "01422318-ca2d-46b8-b66c-3e9e188244ed",
				Name: "giraffes",
				Popularity:  4,
				Location:    "NE",
				Availability: []string{"Wednesday", "Thursday", "Tuesday", "Friday"},
				Residents: []Resident{
					{Name: "Gracia",  Sex: "female", Age: 11},
					{Name: "Antone",  Sex: "male", 	 Age: 9},
					{Name: "Vicky",  	Sex: "female", Age: 12},
					{Name: "Clay",  	Sex: "male",   Age: 4},
					{Name: "Arron",  	Sex: "male", 	 Age: 7},
					{Name: "Bernard", Sex: "male",   Age: 6},
					
				},

			},
		
		},

		Employees: []Employee{
			{
				ID:          "c5b83cb3-a451-49e2-ac45-ff3f54fbe7e1",
				FirstName:   "Nigel",
				LastName:    "Nelson",
				Managers:    []string{"burlId", "olaId"},
				ResponsibleFor: []string{"lionId", "tigersId"},
			},
			{
				ID:          "burlId",
				FirstName:   "Burl",
				LastName:    "Bethea",
				Managers:    []string{"stephanieId"},
				ResponsibleFor: []string{"lionId", "tigersId", "bearsId", "ef3778eb-2844-4c7c-b66c-f432073e1c6b"},
			},
			{
				ID:          "olaId",
				FirstName:   "Ola",
				LastName:    "Orloff",
				Managers:    []string{"stephanieId"},
				ResponsibleFor: []string{"ottersId", "frogsId", "snakesId", "elephantsId"},
			},
			{
				ID:          "56d43ba3-a5a7-40f6-8dd7-cbb05082383f",
				FirstName:   "Wilburn",
				LastName:    "Wishart",
				Managers:    []string{"stephanieId"},
				ResponsibleFor: []string{"burlId", "olaId"},
			},
			{
				ID:          "stephanieId",
				FirstName:   "Stephanie",
				LastName:    "Strauss",
				Managers:    []string{},
				ResponsibleFor: []string{"ottersId", "01422318-ca2d-46b8-b66c-3e9e188244ed"},
			},
			{
				ID:          "4b40a139-d4dc-4f09-822d-ec25e819a5ad",
				FirstName:   "Sharonda",
				LastName:    "Spry",
				Managers:    []string{"burlId", "olaId"},
				ResponsibleFor: []string{"ottersId", "frogsId"},
			},
			{
				ID:          "c1f50212-35a6-4ecd-8223-f835538526c2",
				FirstName:   "Ardith",
				LastName:    "Azevado",
				Managers:    []string{"b0dc644a-5335-489b-8a2c-4e086c7819a2"},
				ResponsibleFor: []string{"tigersId", "bearsId"},
			},
			{
				ID:          "b0dc644a-5335-489b-8a2c-4e086c7819a2",
				FirstName:   "Emery",
				LastName:    "Elser",
				Managers:    []string{"stephanieId"},
				ResponsibleFor: []string{"lionId", "bearsId", "elephantsId"},
			},

		},
		Hours: map[string]Hours{
			"Tuesday":   {Open: 8, Close: 6},
			"Wednesday": {Open: 8, Close: 6},
			"Thursday":  {Open: 10, Close: 8},
			"Friday":    {Open: 10, Close: 8},
			"Saturday":  {Open: 8, Close: 10},
			"Sunday":    {Open: 8, Close: 8},
			"Monday":    {Open: 0, Close: 0},
		},
		Prices: Prices{
			Adult:  49.99,
			Senior: 24.99,
			Child:  20.99,
		},

	}

	



