package preparename

import (
	"github.com/ar-mokhtari/orginfo-generator/entity"
	"testing"
)

func TestPrepareDomainName(t *testing.T) {
	domainNameTest := func(t testing.TB, send string, got, want entity.Domain) {
		if want.SnakeName != got.SnakeName {
			t.Errorf("\t ------------>\t send: %s \t<------------", send)
			t.Errorf("\t\t (file ) ---> got: %s \t\t ---> want: %s", got.SnakeName, want.SnakeName)
		}
		if want.LowerName != got.LowerName {
			t.Errorf("\t\t (Lower) ---> got: %s \t\t ---> want: %s", got.LowerName, want.LowerName)
		}
		if want.UpperName != got.UpperName {
			t.Errorf("\t\t (Upper) ---> got: %s \t\t ---> want: %s", got.UpperName, want.UpperName)
		}
	}
	var dataTest = []string{
		"UserRole", "user role", "User Role", "User _Role", "User_Role", "user_role", "User_ role", "user Role", "User role", "userRole", "_User_ role", "user Role_", "User_ role", "userRole_", "user _role",
	}
	expect := entity.Domain{
		LowerName:  "userRole",
		UpperName:  "UserRole",
		SnakeName:  "user_role",
		AllLowName: "userrole",
		Fields:     nil,
	}
	for _, data := range dataTest {
		t.Run("Test#1. right domain's name test", func(t *testing.T) {
			result, _ := PrepareDomainName(data)
			domainNameTest(t, data, result, expect)
		})
	}
	{
		inputs := []string{"1@nd", " _df#12 jfS", " ", "", "_", "1", "!3D"}
		for _, input := range inputs {
			t.Run("Test#2. wrong regex domain test", func(t *testing.T) {
				_, err := PrepareDomainName(input)
				if err == nil {
					t.Errorf("except error for %s but get nil.", input)
				}
			})
		}
	}
	{
		inputs := []string{"sW _sd12", " _df12 jfS", "12_sd ", "___12_", "_2"}
		for _, input := range inputs {
			t.Run("Test#3. right regex domain test", func(t *testing.T) {
				_, err := PrepareDomainName(input)
				if err != nil {
					t.Errorf("not except error for %s but get error.", input)
				}
			})
		}
	}
}
