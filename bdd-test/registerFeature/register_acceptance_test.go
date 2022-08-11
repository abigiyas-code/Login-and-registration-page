package main

import (
	"github.com/abigiya/internal/Internal/adaptor/framework/dbservice/sqlc"
	"github.com/cucumber/godog"
)

type Connect struct {
	acc *sqlc.Account
}

var con *Connect

// var jsonFile = []byte(`
// 	{
// 		"firstname": binding "required"
// 		"lastname": binding "required"
// 		"email": validate "required, email"
// 		"@username": validate "required, @username"
// 		"password": validate "required, password"
// 	}
// 	`)

func (con *Connect) checkTheEmailAddressExists() error {
	return nil
}

func (con *Connect) checkTheUserNameExists() error {
	return nil
}

func (con *Connect) displayAMessageOf(arg1 string) error {
	return nil
}

func (con *Connect) getAMessageOf(arg1 string) error {
	return nil
}

func (con *Connect) newUser() error {
	return nil
}

func (con *Connect) theResultShouldBe(arg1 string) error {
	return nil
}

func (con *Connect) userDidntFillAskedInformation(FirstName, LastName, EmailAddress, UserName, Password string) error {
	// var userlist UserList
	// err := json.Unmarshal(jsonFile, &userlist)
	// if err != nil {
	// 	err = fmt.Errorf("The error is %s", err)
	// }
	// fmt.Println("FirstName: empty, LastFile: empty, EmailAddress: Empty, Username: Empty, Password: Empty", err)
	return nil
}

func (con *Connect) userFillAllTheNeededInformation(FirstName, LastName, EmailAddress, UserName, Password string) error {
	return nil
}

func (con *Connect) userFillAnInformationOf(arg1 string) error {
	return nil
}

func (con *Connect) userGetASmallBoxesToFillThierInformationOf(FirstName, LastName, EmailAddress, UserName, Password string) error {
	return nil
}

func (con *Connect) userNotBeNeededToFillUserInformationAgainAndAgain() error {
	return nil
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Check the email address exists$`, con.checkTheEmailAddressExists)
	ctx.Step(`^Check the User name exists$`, con.checkTheUserNameExists)
	ctx.Step(`^Display a message of "([^"]*)"$`, con.displayAMessageOf)
	ctx.Step(`^Get a message of "([^"]*)"$`, con.getAMessageOf)
	ctx.Step(`^New user$`, con.newUser)
	ctx.Step(`^The result should be "([^"]*)"$`, con.theResultShouldBe)
	ctx.Step(`^User didn\'t fill asked information  "([^"]*)", "([^"]*)", "([^"]*)", "([^"]*)", "([^"]*)"$`, con.userDidntFillAskedInformation)
	ctx.Step(`^User fill all the needed information "([^"]*)", "([^"]*)", "([^"]*)", "([^"]*)", "([^"]*)"$`, con.userFillAllTheNeededInformation)
	ctx.Step(`^User fill an information of "([^"]*)"$`, con.userFillAnInformationOf)
	ctx.Step(`^User get a small boxes to fill thier information of "([^"]*)", "([^"]*)", "([^"]*)", "([^"]*)", "([^"]*)"$`, con.userGetASmallBoxesToFillThierInformationOf)
	ctx.Step(`^User not be needed to fill user information again and again$`, con.userNotBeNeededToFillUserInformationAgainAndAgain)
}
