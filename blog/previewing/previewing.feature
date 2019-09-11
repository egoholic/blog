Feature: blog previewing
  As a visitor
  I want to visit home page
  In order to investigate recent and most popular content

  Scenario: visitor investigates recent publications
    Given blog has several publications
    When I visit home page
    Then I see 5 recent publications

  Scenario: visitor investigates most popular puplications
    Given blog has several publications
    When I visit home page
    Then I see 5 most popular publications
