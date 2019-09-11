Feature: publication reading
  As a visitor
  I want to open publication
  In order to read it

  Scenario: publication reading
    Given I'm on the home page
    When I open publication preview
    Then I see whole publication
