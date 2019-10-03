Feature: Author previewing
    As a visitor
    I want to open an author page
    In order to investigate author's bio and preview author's publications

    Scenario: visitor previews author
        Given the blog had the following rubrics:
          | slug       | meta_keywords          | meta_description                             | title      | description                                                          |
          | interviews | interview, VIP, celebs | Interviews with known persons.               | Interviews | Know what best practitioners and achievers think about the industry. |
          | how-tos    | how to                 | One-bite sized instructions to achieve more. | How Tos    | One-bite sized instructions to achieve more.                         |
          | ideas      | idea, creativity       | Insightful ideas of using chaos around you.  | Ideas      | Insightful ideas of using chaos around you.                          |
          | releases   | release, news          | Releases of new possibilities.               | Releases   | Releases of new possibilities.                                       |
        And the blog had the following authors:
          | login     | first_name | last_name  | bio                            |
          | jkirk     | James      | Kirk       | Good guy and friend of mine.   |
          | pbird     | Peter      | Bird       | Nice engineer and businessman. |
          | fgelbert  | Frederick  | Gelbert    | Problem solver and researcher. |
        And the blog had the following publications:
          | slug                       | meta_keywords          | meta_description           | title                      | content                 | created_at          | rubric_slug | popularity | author_logins |
          | interview-with-peter-thiel | interview, Peter Thiel | Interview with Piter Thiel | Interview with Piter Thiel | - Hello Piter! - Hello! | 2019-09-01 12:54:34 | interviews  | 65         | jkirk, pbird  |
          | how-to-invest-in-yourself  | investments, funds     | How to invest in yourself  | How to invest in yourself  | It is a good idea!      | 2019-09-12 17:22:09 | how-tos     | 100        | fgelbert      |
          | interview-with-jack-black  | interview, Jack Black  | Interview with Jack Black  | Interview with Jack Black  | - Hello Jack! - Bye!    | 2019-10-03 18:12:12 | interviews  | 24         | jkirk         |
        When I visited "jkirk" author page
        Then I saw "James Kirk" author
        And I saw the following publications:
          | slug                       | title                      |
          | interview-with-peter-thiel | Interview with Piter Thiel |
          | how-to-invest-in-yourself  | How to invest in yourself  |
          | interview-with-jack-black  | Interview with Jack Black  |
