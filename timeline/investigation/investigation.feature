Feature: Timeline investigation
  As a visitor
  I want to visit Timeline
  I order to preview publications chonologically

  Scenario: timeline investigation
    Given blog has several publications
    And I'm on home page
    When I visit "Timeline"
    Then I see publications chronologically
