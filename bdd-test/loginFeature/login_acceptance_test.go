package main

import "github.com/cucumber/godog"

func checkPassword() error {
	return godog.ErrPending
}

func checkUserHaveBeenRegistered() error {
	return godog.ErrPending
}

func checkUserNameInFormatOf(arg1 string) error {
	return godog.ErrPending
}

func forNotMatchOfPasswordViaUserNameDisplayMessage(arg1 string) error {
	return godog.ErrPending
}

func forNotMatchOfUserNameDisplayMessage(arg1 string) error {
	return godog.ErrPending
}

func forNotRegisteredUserDisplayMessageOf(arg1 string) error {
	return godog.ErrPending
}

func forRegisteredUserDisplayMessageOf(arg1 string) error {
	return godog.ErrPending
}

func userFillAskedInformation() error {
	return godog.ErrPending
}

func userFillAskedInformationOfLogin(arg1, arg2 string) error {
	return godog.ErrPending
}

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^Check Password$`, checkPassword)
	ctx.Step(`^Check user have been registered$`, checkUserHaveBeenRegistered)
	ctx.Step(`^Check User name in format of "([^"]*)"$`, checkUserNameInFormatOf)
	ctx.Step(`^For not match of Password via User name, display message "([^"]*)"$`, forNotMatchOfPasswordViaUserNameDisplayMessage)
	ctx.Step(`^For not match of User name, display message "([^"]*)"$`, forNotMatchOfUserNameDisplayMessage)
	ctx.Step(`^For not registered user display message of "([^"]*)"$`, forNotRegisteredUserDisplayMessageOf)
	ctx.Step(`^For registered user display message of "([^"]*)"$`, forRegisteredUserDisplayMessageOf)
	ctx.Step(`^User fill asked information$`, userFillAskedInformation)
	ctx.Step(`^User fill asked information of login "([^"]*)", "([^"]*)"$`, userFillAskedInformationOfLogin)
}
