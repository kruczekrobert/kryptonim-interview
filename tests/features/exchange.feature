Feature: Exchange endpoint tests

  Background:
    Given router "Core"

    Given models "ExchangeRates"
      | Id | CryptoCurrency | DecimalPlaces | USDRate    |
      | 1  | BEER           | 18            | 0.00002461 |
      | 2  | FLOKI          | 18            | 0.0001428  |
      | 3  | GATE           | 18            | 6.87       |
      | 4  | USDT           | 6             | 0.999      |
      | 5  | WBTC           | 8             | 57037.22   |

  Scenario: Successful exchange response
    When have request
    """
    {
      "method": "GET",
      "path": "/core/exchange?from=WBTC&to=USDT&amount=1.0"
    }
    """

    Then have json response with status "200"
    """
    {
      "response": {
        "from": "WBTC",
        "to": "USDT",
        "amount": 57094.314314
      }
    }
    """

  Scenario Outline: CryptoCurrency <case> query param does not exists
    When have request
    """
    {
      "method": "GET",
      "path": "/core/exchange?from=<from>&to=<to>&amount=1.0"
    }
    """

    Then have json response with status "400"
    """
    {
      "response": {}
    }
    """

    Examples:
      | case | from  | to    |
      | from | MATIC | USDT  |
      | to   | USDT  | MATIC |

  Scenario Outline: Missing <case> query param
    When have request
    """
    {
      "method": "GET",
      "path": "/core/exchange?<query>"
    }
    """

    Then have json response with status "400"
    """
    {
      "response": {}
    }
    """

    Examples:
      | case   | query                 |
      | amount | from=MATIC&to=GATE    |
      | from   | to=GATE&amount=1.0    |
      | to     | from=MATIC&amount=1.0 |

  Scenario Outline: Amount is invalid, case: <case>
    When have request
    """
    {
      "method": "GET",
      "path": "/core/exchange?from=WBTC&to=USDT&amount=<amount>"
    }
    """

    Then have json response with status "400"
    """
    {
      "response": {}
    }
    """

    Examples:
      | case         | amount |
      | negative     | -1.0   |
      | not a number | abc    |
      | zero         | 0      |

