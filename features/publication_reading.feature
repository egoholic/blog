Feature: Publication reading
    As a visitor
    I want to open publication
    In order to read its content

    Scenario: visitor reads publication
        Given the blog had the following rubrics:
          | slug    | meta_keywords | meta_description                             | title   | description                                  |
          | how-tos | how to        | One-bite sized instructions to achieve more. | How Tos | One-bite sized instructions to achieve more. |
        And the blog had the following publications:
          | slug                       | meta_keywords             | meta_description            | title                       | content                  | created_at          | rubric_slug | popularity |
          | how-to-invent-great-titles | invention, title, writing | How to invent great titles? | How to invent great titles? | There are a lot ways to. | 2019-10-15 13:34:23 | how-tos     | 59         |
        When I visited "how-to-invent-great-titles" publication page
        Then I read "How to invent great titles?" publication

    Scenario: visitor tries to read unexisting publication
        When I visited "wrong-slug" publication page
        Then I see that page not found
