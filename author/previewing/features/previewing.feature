Feature: author previewing
  As a visitor
  I want to preview author
  So that I can get more information about his/her person
  And preview his/her publications

  Scenario: author's bio investigation
    Given I'm on the "About Us" page
    When I open an author
    Then I see author's bio

  Scenario: author's publications preview
    Given: blog has several publications
    And I'm on the "About Us" page
    When I open an author
    Then I see a list of author's publications
