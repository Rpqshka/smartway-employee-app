package smartway_employee_app

type Employee struct {
	Id         int        `json:"id" db:"id"`
	Name       string     `json:"name" db:"name" binding:"required"`
	Surname    string     `json:"surname" db:"surname" binding:"required"`
	Phone      string     `json:"phone" db:"employee_phone" binding:"required"`
	CompanyId  int        `json:"companyId" db:"company_id" binding:"required"`
	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type Passport struct {
	Type   string `json:"type" db:"passport_type" binding:"required"`
	Number string `json:"number" db:"passport_number" binding:"required"`
}

type Department struct {
	Name  string `json:"name" db:"department_name" binding:"required"`
	Phone string `json:"phone" db:"department_phone" binding:"required"`
}

type UpdateEmployee struct {
	Id        int    `json:"id" db:"id" binding:"required"`
	Name      string `json:"name" db:"name"`
	Surname   string `json:"surname" db:"surname"`
	Phone     string `json:"phone" db:"employee_phone"`
	CompanyId int    `json:"companyId" db:"company_id"`
	Passport  struct {
		Type   string `json:"type" db:"passport_type"`
		Number string `json:"number" db:"passport_number"`
	} `json:"passport"`
	Department struct {
		Name  string `json:"name" db:"department_name"`
		Phone string `json:"phone" db:"department_phone"`
	} `json:"department"`
}
