package staff

type Staff struct {
	ID           int    `json:"id" bun:",pk,autoincrement"`
	Username     string `json:"username" bun:",unique,notnull"`
	Password     string `json:"password" bun:",notnull"`
	HospitalName string `json:"hospital_name" bun:",notnull"`
	HospitalID   int    `json:"hospital_id" bun:",notnull"`
}
