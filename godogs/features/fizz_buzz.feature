Feature: As a human, I want to get FizzBuzz

  Scenario: FizzBuzz should success with valid parameters
    Given I reset client
    When I set request query
      | fizz_modulo | 3   |
      | buzz_modulo | 5   |
      | limit       | 15  |
      | fizz_string | foo |
      | buzz_string | bar |
    And I GET http://localhost:8080/v1/fizz_buzz
    Then response should contain sadfsdf
