@Scenarios @api
Feature: Login 

    @api-tc-01
    Scenario Outline: Log to the page for not registered user
    When User fill asked information
    Then Check user have been registered 
    And For not registered user display message of "Please register to logged in"

    Example:
      | User information |  Status |
      | User Name        |  fail   |
      | Password         |  fail   |
      | Show Password    |         |
 

    @api-tc-02
    Scenario Outline: Error in User name and format
    When User fill asked information of login "User name", "Password"
    Then Check User name in format of "@Username"
    And For not match of User name, display message "invalid user name, please try again" 
   
    Example:
      | User information | Status |
      | User Name        |  fail  |
      | Password         |   -    |
      | Show Password    |        |


    @api-tc-03
    Scenario Outline: Error in Password
    When User fill asked information of login "User name", "Password"
    Then Check Password
    And For not match of Password via User name, display message "incorrect password" 
   
    Example:
      | User information | Status |
      | User Name        |  OK    |
      | Password         |  fail  |
      | Show Password    |        |


    @api-tc-04
    Scenario Outline: Log to the page for registered user
    When User fill asked information
    Then Check user have been registered 
    And For registered user display message of "Confirmed login"
   
    Example:
      | User information |Status |
      | User Name        |  OK   |
      | Password         |  OK   |
      | Show Password    |       |
