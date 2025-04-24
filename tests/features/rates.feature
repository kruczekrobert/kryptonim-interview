Feature: Rates endpoint tests

  Background:
    Given router "Core"

  Scenario: Successful rates response (3)
    Given json mock "openexchangerates"
    """
    {
      "method": "GET",
      "path": "openexchangerates-test-url/api/latest.json?app_id=my_app_id&base=USD&symbols=USD,GBP,EUR",
      "status": 200,
      "response": {
        "rates": {
          "EUR": 0.8802,
          "GBP": 0.751589,
          "USD": 1
        }
      }
    }
    """

    When have request
    """
    {
      "method": "GET",
      "path": "/core/rates?currencies=USD,GBP,EUR"
    }
    """

    Then have json response with status "200"
    """
    {
      "response": [
        {"from":"USD","rate":1.330514416788963,"to":"GBP"},
        {"from":"USD","rate":1.1361054305839582,"to":"EUR"},
        {"from":"GBP","rate":0.751589,"to":"USD"},
        {"from":"GBP","rate":0.8538843444671665,"to":"EUR"},
        {"from":"EUR","rate":0.8802,"to":"USD"},
        {"from":"EUR","rate":1.1711187896576454,"to":"GBP"}
      ]
    }
    """

    And json mock "openexchangerates" was called

  Scenario: Too few currencies in the parameter
    Given json mock "openexchangerates"
    """
    {
      "method": "GET",
      "path": "openexchangerates-test-url/api/latest.json?app_id=my_app_id&base=USD&symbols=USD,GBP,EUR",
      "status": 200,
      "response": {
        "rates": {
          "EUR": 0.8802,
          "GBP": 0.751589,
          "USD": 1
        }
      }
    }
    """
    When have request
    """
    {
      "method": "GET",
      "path": "/core/rates?currencies=USD"
    }
    """

    Then have json response with status "400"
    """
    {
      "response": {}
    }
    """

    And json mock "openexchangerates" was not called

  Scenario: Exchange rates provider error
    Given json mock "openexchangerates"
    """
    {
      "method": "GET",
      "path": "openexchangerates-test-url/api/latest.json?app_id=my_app_id&base=USD&symbols=USD,GBP,EUR",
      "status": 500,
      "response": {}
    }
    """

    When have request
    """
    {
      "method": "GET",
      "path": "/core/rates?currencies=USD,GBP,EUR"
    }
    """

    Then have json response with status "400"
    """
    {
      "response": {}
    }
    """

    And json mock "openexchangerates" was called
