Feature: I want to retrieve FizzBuzz stats

  Scenario: Retreive stats should success
    Given I reset HTTP client

    # --------------------------------------------------------------------------------
    # Reset record requests
    # --------------------------------------------------------------------------------
    And I reset fizzBuzz request stats
    Then response status code should be 200

    # --------------------------------------------------------------------------------
    # Record requests
    # --------------------------------------------------------------------------------
    When I set request query
      | fizz_modulo | 3   |
      | buzz_modulo | 5   |
      | limit       | 15  |
      | fizz_string | foo |
      | buzz_string | bar |
    And I get fizzBuzz
    Then response status code should be 200
    And I get fizzBuzz
    Then response status code should be 200

    When I set request query
      | fizz_modulo | 3       |
      | buzz_modulo | 10      |
      | limit       | 20      |
      | fizz_string | buzz    |
      | buzz_string | leclair |
    And I get fizzBuzz
    Then response status code should be 200
    And I get fizzBuzz
    Then response status code should be 200

    When I set request query
      | fizz_modulo | 3      |
      | buzz_modulo | 10     |
      | limit       | 20     |
      | fizz_string | buzz   |
      | buzz_string | aldrin |
    And I get fizzBuzz
    Then response status code should be 200

    # --------------------------------------------------------------------------------
    # Retrieve stats
    # --------------------------------------------------------------------------------
    When I get fizzBuzz requests stats
    Then response status code should be 200
    Then json response should resemble
    """
    {
      "requests": [
          {
              "FizzModulo": 3,
              "BuzzModulo": 5,
              "Limit": 15,
              "FizzString": "foo",
              "BuzzString": "bar"
          },
          {
              "FizzModulo": 3,
              "BuzzModulo": 10,
              "Limit": 20,
              "FizzString": "buzz",
              "BuzzString": "leclair"
          }
      ],
      "total": 2
    }
    """
