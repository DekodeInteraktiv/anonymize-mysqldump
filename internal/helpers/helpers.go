package helpers

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/xwb1989/sqlparser"
	"syreclabs.com/go/faker"
)

// GetFakerFuncs creates a map of faker helper functions.
func GetFakerFuncs() map[string]func(*sqlparser.SQLVal) *sqlparser.SQLVal {
	fakerHelpers := map[string]func(*sqlparser.SQLVal) *sqlparser.SQLVal{
		"username":             generateUsername,
		"password":             generatePassword,
		"email":                generateEmail,
		"url":                  generateURL,
		"name":                 generateName,
		"firstName":            generateFirstName,
		"lastName":             generateLastName,
		"phoneNumber":          generatePhoneNumber,
		"addressFull":          generateAddress,
		"addressStreet":        generateStreetAddress,
		"addressPostCode":      generatePostcode,
		"addressCountry":       generateCountry,
		"paragraph":            generateParagraph,
		"shortString":          generateShortString,
		"ipv4":                 generateIPv4,
		"companyName":          generateCompanyName,
		"companyNumber":        generateCompanyNumber,
		"creditCardNumber":     generateCreditCardNumber,
		"creditCardExpiryDate": generateCreditCardExpiryDate,
		"creditCardType":       generateCreditCardType,
	}

	return fakerHelpers
}

func generateUsername(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().UserName()))
}

func generatePassword(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	password := md5.Sum([]byte(faker.Internet().Password(8, 14)))
	return sqlparser.NewStrVal([]byte(hex.EncodeToString(password[:])))
}

func generateEmail(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().SafeEmail()))
}

func generatePhoneNumber(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.PhoneNumber().CellPhone()))
}

func generateURL(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().Url()))
}

func generateName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().Name()))
}

func generateFirstName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().FirstName()))
}

func generateLastName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Name().LastName()))
}

func generateParagraph(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Lorem().Sentence(3)))
}

func generateIPv4(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Internet().IpV4Address()))
}

func generateAddress(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().String()))
}

func generateStreetAddress(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().StreetAddress()))
}

func generatePostcode(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().Postcode()))
}

func generateCountry(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Address().Country()))
}

func generateCreditCardNumber(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Business().CreditCardNumber()))
}

func generateCreditCardExpiryDate(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Business().CreditCardExpiryDate()))
}

func generateCreditCardType(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Business().CreditCardType()))
}

func generateCompanyName(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Company().Name()))
}

func generateCompanyNumber(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Number().Number(9)))
}

func generateShortString(value *sqlparser.SQLVal) *sqlparser.SQLVal {
	return sqlparser.NewStrVal([]byte(faker.Lorem().Characters(30)))
}
