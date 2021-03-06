// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/sut63/team05/ent/bank"
	"github.com/sut63/team05/ent/hospital"
	"github.com/sut63/team05/ent/inquiry"
	"github.com/sut63/team05/ent/insurance"
	"github.com/sut63/team05/ent/member"
	"github.com/sut63/team05/ent/moneytransfer"
	"github.com/sut63/team05/ent/officer"
	"github.com/sut63/team05/ent/payback"
	"github.com/sut63/team05/ent/payment"
	"github.com/sut63/team05/ent/recordinsurance"
	"github.com/sut63/team05/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	bankFields := schema.Bank{}.Fields()
	_ = bankFields
	// bankDescBankType is the schema descriptor for bank_type field.
	bankDescBankType := bankFields[0].Descriptor()
	// bank.BankTypeValidator is a validator for the "bank_type" field. It is called by the builders before save.
	bank.BankTypeValidator = bankDescBankType.Validators[0].(func(string) error)
	hospitalFields := schema.Hospital{}.Fields()
	_ = hospitalFields
	// hospitalDescHospitalName is the schema descriptor for hospital_name field.
	hospitalDescHospitalName := hospitalFields[0].Descriptor()
	// hospital.HospitalNameValidator is a validator for the "hospital_name" field. It is called by the builders before save.
	hospital.HospitalNameValidator = hospitalDescHospitalName.Validators[0].(func(string) error)
	inquiryFields := schema.Inquiry{}.Fields()
	_ = inquiryFields
	// inquiryDescInquiryMessages is the schema descriptor for Inquiry_messages field.
	inquiryDescInquiryMessages := inquiryFields[0].Descriptor()
	// inquiry.InquiryMessagesValidator is a validator for the "Inquiry_messages" field. It is called by the builders before save.
	inquiry.InquiryMessagesValidator = inquiryDescInquiryMessages.Validators[0].(func(string) error)
	// inquiryDescInquiryTimeMessages is the schema descriptor for Inquiry_time_messages field.
	inquiryDescInquiryTimeMessages := inquiryFields[1].Descriptor()
	// inquiry.DefaultInquiryTimeMessages holds the default value on creation for the Inquiry_time_messages field.
	inquiry.DefaultInquiryTimeMessages = inquiryDescInquiryTimeMessages.Default.(func() time.Time)
	insuranceFields := schema.Insurance{}.Fields()
	_ = insuranceFields
	// insuranceDescInsuranceAddress is the schema descriptor for insurance_address field.
	insuranceDescInsuranceAddress := insuranceFields[0].Descriptor()
	// insurance.InsuranceAddressValidator is a validator for the "insurance_address" field. It is called by the builders before save.
	insurance.InsuranceAddressValidator = insuranceDescInsuranceAddress.Validators[0].(func(string) error)
	// insuranceDescInsuranceInsurer is the schema descriptor for insurance_insurer field.
	insuranceDescInsuranceInsurer := insuranceFields[1].Descriptor()
	// insurance.InsuranceInsurerValidator is a validator for the "insurance_insurer" field. It is called by the builders before save.
	insurance.InsuranceInsurerValidator = insuranceDescInsuranceInsurer.Validators[0].(func(string) error)
	// insuranceDescInsuranceTimeBuy is the schema descriptor for insurance_time_buy field.
	insuranceDescInsuranceTimeBuy := insuranceFields[2].Descriptor()
	// insurance.DefaultInsuranceTimeBuy holds the default value on creation for the insurance_time_buy field.
	insurance.DefaultInsuranceTimeBuy = insuranceDescInsuranceTimeBuy.Default.(func() time.Time)
	// insuranceDescInsuranceTimeFirstpay is the schema descriptor for insurance_time_firstpay field.
	insuranceDescInsuranceTimeFirstpay := insuranceFields[3].Descriptor()
	// insurance.DefaultInsuranceTimeFirstpay holds the default value on creation for the insurance_time_firstpay field.
	insurance.DefaultInsuranceTimeFirstpay = insuranceDescInsuranceTimeFirstpay.Default.(func() time.Time)
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescMemberEmail is the schema descriptor for member_email field.
	memberDescMemberEmail := memberFields[0].Descriptor()
	// member.MemberEmailValidator is a validator for the "member_email" field. It is called by the builders before save.
	member.MemberEmailValidator = memberDescMemberEmail.Validators[0].(func(string) error)
	// memberDescMemberName is the schema descriptor for member_name field.
	memberDescMemberName := memberFields[1].Descriptor()
	// member.MemberNameValidator is a validator for the "member_name" field. It is called by the builders before save.
	member.MemberNameValidator = memberDescMemberName.Validators[0].(func(string) error)
	// memberDescMemberPassword is the schema descriptor for member_password field.
	memberDescMemberPassword := memberFields[2].Descriptor()
	// member.MemberPasswordValidator is a validator for the "member_password" field. It is called by the builders before save.
	member.MemberPasswordValidator = memberDescMemberPassword.Validators[0].(func(string) error)
	moneytransferFields := schema.MoneyTransfer{}.Fields()
	_ = moneytransferFields
	// moneytransferDescMoneytransferType is the schema descriptor for moneytransfer_type field.
	moneytransferDescMoneytransferType := moneytransferFields[0].Descriptor()
	// moneytransfer.MoneytransferTypeValidator is a validator for the "moneytransfer_type" field. It is called by the builders before save.
	moneytransfer.MoneytransferTypeValidator = moneytransferDescMoneytransferType.Validators[0].(func(string) error)
	officerFields := schema.Officer{}.Fields()
	_ = officerFields
	// officerDescOfficerEmail is the schema descriptor for officer_email field.
	officerDescOfficerEmail := officerFields[0].Descriptor()
	// officer.OfficerEmailValidator is a validator for the "officer_email" field. It is called by the builders before save.
	officer.OfficerEmailValidator = officerDescOfficerEmail.Validators[0].(func(string) error)
	// officerDescOfficerName is the schema descriptor for officer_name field.
	officerDescOfficerName := officerFields[1].Descriptor()
	// officer.OfficerNameValidator is a validator for the "officer_name" field. It is called by the builders before save.
	officer.OfficerNameValidator = officerDescOfficerName.Validators[0].(func(string) error)
	// officerDescOfficerPassword is the schema descriptor for officer_password field.
	officerDescOfficerPassword := officerFields[2].Descriptor()
	// officer.OfficerPasswordValidator is a validator for the "officer_password" field. It is called by the builders before save.
	officer.OfficerPasswordValidator = officerDescOfficerPassword.Validators[0].(func(string) error)
	paybackFields := schema.Payback{}.Fields()
	_ = paybackFields
	// paybackDescAccountnumber is the schema descriptor for Accountnumber field.
	paybackDescAccountnumber := paybackFields[0].Descriptor()
	// payback.AccountnumberValidator is a validator for the "Accountnumber" field. It is called by the builders before save.
	payback.AccountnumberValidator = paybackDescAccountnumber.Validators[0].(func(string) error)
	// paybackDescTransfertime is the schema descriptor for Transfertime field.
	paybackDescTransfertime := paybackFields[1].Descriptor()
	// payback.DefaultTransfertime holds the default value on creation for the Transfertime field.
	payback.DefaultTransfertime = paybackDescTransfertime.Default.(func() time.Time)
	paymentFields := schema.Payment{}.Fields()
	_ = paymentFields
	// paymentDescAccountName is the schema descriptor for account_name field.
	paymentDescAccountName := paymentFields[0].Descriptor()
	// payment.AccountNameValidator is a validator for the "account_name" field. It is called by the builders before save.
	payment.AccountNameValidator = paymentDescAccountName.Validators[0].(func(string) error)
	// paymentDescAccountNumber is the schema descriptor for account_number field.
	paymentDescAccountNumber := paymentFields[1].Descriptor()
	// payment.AccountNumberValidator is a validator for the "account_number" field. It is called by the builders before save.
	payment.AccountNumberValidator = paymentDescAccountNumber.Validators[0].(func(string) error)
	// paymentDescTransferTime is the schema descriptor for transfer_time field.
	paymentDescTransferTime := paymentFields[2].Descriptor()
	// payment.DefaultTransferTime holds the default value on creation for the transfer_time field.
	payment.DefaultTransferTime = paymentDescTransferTime.Default.(func() time.Time)
	recordinsuranceFields := schema.Recordinsurance{}.Fields()
	_ = recordinsuranceFields
	// recordinsuranceDescRecordinsuranceTime is the schema descriptor for recordinsurance_time field.
	recordinsuranceDescRecordinsuranceTime := recordinsuranceFields[0].Descriptor()
	// recordinsurance.DefaultRecordinsuranceTime holds the default value on creation for the recordinsurance_time field.
	recordinsurance.DefaultRecordinsuranceTime = recordinsuranceDescRecordinsuranceTime.Default.(func() time.Time)
}
