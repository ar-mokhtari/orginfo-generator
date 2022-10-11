package env

import "regexp"

const (
	//Minimum eight characters, at least one letter and one number:
	RegLetterNo8 = `^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$`
	//Minimum eight characters, at least one letter, one number and one special character:
	RegLetterNoSPchar8 = `^(?=.*[A-Za-z])(?=.*\d)(?=.*[@$!%*#?&])[A-Za-z\d@$!%*#?&]{8,}$`
	//Minimum eight characters, at least one uppercase letter, one lowercase letter and one number:
	RegULN8 = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)[a-zA-Z\d]{8,}$`
	//Minimum eight characters, at least one uppercase letter, one lowercase letter, one number and one special character:
	RegULNS8 = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
	//Minimum eight and maximum 10 characters, at least one uppercase letter, one lowercase letter, one number and one special character:
	RegULNS10 = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,10}$`

	//Iran national ID
	RegIranNid = `^[\d]{10}$`

	//Iran Mobile number Form
	IranMobileForm = `^09[0-9][0-9]{8}$`
)

var (
	//regexp only lower and upper case + numbers min=2
	RegexazAZ2 = regexp.MustCompile(`^[A-Za-z0-9_ ]{2,}$`)
)
