Feature: rubric previewing
  As a visitor
  I want to preview rubrics
  So that I can preview corresponding publications

  Scenario: rubric description investigation
    Given I'm on the home page
    When I open a rubric
    Then I see rubric's description

  Scenario: rubric's publications previewing
    Given blog has several publications
    And I'm on the home page
    When I open a rubric
    Then I can preview rubric's publications
