Feature: As a human, I want to retrieve FizzBuzz stats

  Scenario: Retreive stats should success
    Given I reset client

    # --------------------------------------------------------------------------------
    # Record requests
    # --------------------------------------------------------------------------------
    When I set request query
      | fizz_modulo | 3   |
      | buzz_modulo | 5   |
      | limit       | 15  |
      | fizz_string | foo |
      | buzz_string | bar |
    And I GET http://localhost:8080/v1/fizz_buzz
    Then response status code should be 200
    And I GET http://localhost:8080/v1/fizz_buzz
    Then response status code should be 200

    When I set request query
      | fizz_modulo | 3       |
      | buzz_modulo | 10      |
      | limit       | 20      |
      | fizz_string | buzz    |
      | buzz_string | leclair |
    And I GET http://localhost:8080/v1/fizz_buzz
    Then response status code should be 200
    And I GET http://localhost:8080/v1/fizz_buzz
    Then response status code should be 200

    When I set request query
      | fizz_modulo | 3      |
      | buzz_modulo | 10     |
      | limit       | 20     |
      | fizz_string | buzz   |
      | buzz_string | aldrin |
    And I GET http://localhost:8080/v1/fizz_buzz
    Then response status code should be 200
    And I GET http://localhost:8080/v1/fizz_buzz
    Then response status code should be 200


    # --------------------------------------------------------------------------------
    # Retrieve stats
    # --------------------------------------------------------------------------------
    When I GET http://localhost:8080/v1/stats
    Then response status code should be 200
    Then json response should resemble
    """
    {
      "details": [
          "missing fizz modulo",
          "missing buzz modulo",
          "missing limit",
          "missing fizz string",
          "missing buzz string"
      ],
      "status": 400
    }
    """
