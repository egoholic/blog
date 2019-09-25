Feature: Blog previewing
    As a visitor
    I want to visit home page
    In order to investigate recent and most popular content

    Background:
        Given there is a blog
        And the blog has next rubrics:
          | slug       | meta_keywords          | meta_description                             | title      | description                                                          |
          | interviews | interview, VIP, celebs | Interviews with known persons.               | Interviews | Know what best practitioners and achievers think about the industry. |
          | how-tos    | how to                 | One-bite sized instructions to achieve more. | How Tos    | One-bite sized instructions to achieve more.                         |
          | ideas      | idea, creativity       | Insightful ideas of using chaos around you.  | Ideas      | Insightful ideas of using chaos around you.                          |
          | releases   | release, news          | Releases of new possibilities.               | Releases   | Releases of new possibilities.                                       |
        And the blog has next publications:
          | slug                         | meta_keywords                    | meta_description                                                 | title                                                            | content                                          | created_at          | rubric_slug | popularity |
          | interview-with-peter-thiel   | interview, Peter Thiel           | Interview with Piter Thiel                                       | Interview with Piter Thiel                                       | - Hello Piter! - Hello!                          | 2019-09-01 12:54:34 | interviews  | 65         |
          | how-to-invest-in-yourself    | investments, funds               | How to invest in yourself                                        | How to invest in yourself                                        | It is a good idea!                               | 2019-09-12 17:22:09 | how-tos     | 100        |
          | interview-with-jack-black    | interview, Jack Black            | Interview with Jack Black                                        | Interview with Jack Black                                        | - Hello Jack! - Bye!                             | 2019-10-03 18:12:12 | interviews  | 24         |
          | bottlenecks-metaphora        | bottleneck, metaphora            | What "Bottleneck" metaphora teaches us?                          | What "Bottleneck" metaphora teaches us?                          | A lot!                                           | 2019-10-07 11:13:15 | ideas       | 34         |
          | how-to-invent-great-titles   | invention, title, writing        | How to invent great titles?                                      | How to invent great titles?                                      | There are a lot ways to.                         | 2019-10-15 13:34:23 | how-tos     | 59         |
          | interview-with-bill-gates    | interview, Bill Gates            | Interview with Bill Gates                                        | Interview with Bill Gates                                        | - How are you? - Fine.                           | 2019-10-23 15:54:11 | interviews  | 45         |
          | swarm-landing-pages          | release, swarm, landing          | Swarm - evolutional approach to landing pages                    | Swarm - evolutional approach to landing pages                    | Swarm is the best and only!                      | 2019-10-29 12:09:01 | releases    | 27         |
          | interview-with-edward-deming | interview, Edward Deming         | Interview with Edward Deming                                     | Interview with Edward Deming                                     | - How is going? - It is going normal.            | 2019-11-07 09:23:03 | interviews  | 121        |
          | landing-pages-styles-market  | release, landing, styles         | Landing pages styles market released!                            | Landing pages styles market released!                            | Now designers can make money togather with Stoa! | 2019-11-09 12:04:17 | releases    | 102        |
          | onboarding-and-outboarding   | release, onboarding, outboarding | We'are happy to present new onboarding and outboarding features! | We are happy to present new onboarding and outboarding features! | Onboarding and outbording are out!               | 2019-11-21 19:00:59 | releases    | 48         |

    Scenario: visitor investigates recent publications
        When I visit home page
        Then I see next recent publications:
          | slug                         | title                                                            |
          | interview-with-bill-gates    | Interview with Bill Gates                                        |
          | swarm-landing-pages          | Swarm - evolutional approach to landing pages                    |
          | interview-with-edward-deming | Interview with Edward Deming                                     |
          | landing-pages-styles-market  | Landing pages styles market released!                            |
          | onboarding-and-outboarding   | We are happy to present new onboarding and outboarding features! |


    Scenario: visitor investigates most popular puplications
        When I visit home page
        Then I see next most popular publications:
          | slug                         | title                                 |
          | interview-with-edward-deming | Interview with Edward Deming          |
          | landing-pages-styles-market  | Landing pages styles market released! |
          | how-to-invest-in-yourself    | How to invest in yourself             |
          | interview-with-peter-thiel   | Interview with Piter Thiel            |
          | how-to-invent-great-titles   | How to invent great titles?           |

