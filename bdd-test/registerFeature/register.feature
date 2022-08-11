Feature: Register

  # Background: "/register"

    Scenario Outline: Empty information 
    Given New user 
    When User didn't fill asked information  "<First name>", "<Last name>", "<email address>", "<User name>", "<Password>"
    Then The result should be "Please fill all the information" 

     Example:
      | User information | Status |         format             |
      | First Name       | empty  |    alphabetic character    |
      | Last Name        | empty  |    alphabetic character    |
      | email address    | empty  | example-name@mail-type.com | 
      | User Name        | empty  |   @alphabets or intiger    |
      | Password         | empty  |      intiger number        |
    

    Scenario Outline: Email exist  
    Given New user 
    When User fill an information of "<email address>"
    Then Check the email address exists
    And Display a message of "email address had registered, use another email"

      Example:
      | User information | Status  |  
      | First Name       |   OK    |  
      | Last Name        |   OK    |  
      | email address    |  fail   | 
      | User Name        |   -     | 
      | Password         |   -     | 
    

    Scenario Outline: User name exist  
    Given New user 
    When User fill an information of "<User name>"
    Then Check the User name exists
    And Display a message of "User name had already taken, fill another name"

     Example:
      | User information | Status  |  
      | First Name       |   OK    |  
      | Last Name        |   OK    |  
      | email address    |   OK    | 
      | User Name        |  fail   | 
      | Password         |   -     | 


    Scenario Outline: Register User information 
    Given User get a small boxes to fill thier information of "<First name>", "<Last name>", "<email address>", "<User name>", "<Password>"
    When User fill all the needed information "<First name>", "<Last name>", "<email address>", "<User name>", "<Password>"
    Then Get a message of "Registration done"
    And User not be needed to fill user information again and again 

    Example:
      | User information |    Status      |         format             |
      | First Name       |      OK        |    alphabetic character    |
      | Last Name        |      OK        |    alphabetic character    |
      | email address    |      OK        | example-name@mail-type.com | 
      | User Name        |  MAX length 5  |   @alphabets or intiger    |
      | Password         |  MAX length 8  |   alphabetic character     |
