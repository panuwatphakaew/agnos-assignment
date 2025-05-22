package patient

type Patient struct {
	ID                  int    `json:"id" bun:"id,pk"`
	FirstName           string `json:"firstName" bun:"first_name"`
	MiddleName          string `json:"middleName" bun:"middle_name"`
	LastName            string `json:"lastName" bun:"last_name"`
	DateOfBirth         string `json:"dateOfBirth" bun:"date_of_birth"`
	PatientHospitalName string `json:"patientHospitalName" bun:"patient_hn"`
	NationalID          string `json:"nationalId" bun:"national_id"`
	PassportID          string `json:"passportId" bun:"passport_id"`
	PhoneNumber         string `json:"phoneNumber" bun:"phone_number"`
	Email               string `json:"email" bun:"email"`
	Gender              string `json:"gender" bun:"gender"`
	HospitalID          int    `json:"hospitalId" bun:"hospital_id"`
}
