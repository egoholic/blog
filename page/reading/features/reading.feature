Feature: "About Us" page reading
  As a visitor
  I want to investigate "About Us" page
  In order to know more about the authors and blog's purpose

  Scenario: desciption investigation
    Given I'm on the home page
    When I visit "About Us" page
    Then I see blog description

  Scenario: authors list investigation
    Given I'm on the home page
    When I visit "About Us" page
    Then I see blog authors' short bios

Feature: "Contacts" page reading
  As a visitor
  I want to investigate "Contacts" page
  In order to be able to contact with blog authors

  Scenario: contacts list investigation
    Given I'm on the home page
    When I visit "Contacts" page
    Then I see contacts list
