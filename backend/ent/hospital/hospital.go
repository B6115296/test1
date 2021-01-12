// Code generated by entc, DO NOT EDIT.

package hospital

const (
	// Label holds the string label denoting the hospital type in the database.
	Label = "hospital"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldHospitalName holds the string denoting the hospital_name field in the database.
	FieldHospitalName = "hospital_name"

	// EdgeHospitalInsurance holds the string denoting the hospital_insurance edge name in mutations.
	EdgeHospitalInsurance = "hospital_insurance"
	// EdgeHospitalRecordinsurance holds the string denoting the hospital_recordinsurance edge name in mutations.
	EdgeHospitalRecordinsurance = "hospital_recordinsurance"

	// Table holds the table name of the hospital in the database.
	Table = "hospitals"
	// HospitalInsuranceTable is the table the holds the hospital_insurance relation/edge.
	HospitalInsuranceTable = "insurances"
	// HospitalInsuranceInverseTable is the table name for the Insurance entity.
	// It exists in this package in order to avoid circular dependency with the "insurance" package.
	HospitalInsuranceInverseTable = "insurances"
	// HospitalInsuranceColumn is the table column denoting the hospital_insurance relation/edge.
	HospitalInsuranceColumn = "hospital_id"
	// HospitalRecordinsuranceTable is the table the holds the hospital_recordinsurance relation/edge.
	HospitalRecordinsuranceTable = "recordinsurances"
	// HospitalRecordinsuranceInverseTable is the table name for the Recordinsurance entity.
	// It exists in this package in order to avoid circular dependency with the "recordinsurance" package.
	HospitalRecordinsuranceInverseTable = "recordinsurances"
	// HospitalRecordinsuranceColumn is the table column denoting the hospital_recordinsurance relation/edge.
	HospitalRecordinsuranceColumn = "hospital_id"
)

// Columns holds all SQL columns for hospital fields.
var Columns = []string{
	FieldID,
	FieldHospitalName,
}

var (
	// HospitalNameValidator is a validator for the "hospital_name" field. It is called by the builders before save.
	HospitalNameValidator func(string) error
)