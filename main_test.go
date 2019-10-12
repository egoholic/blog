package main

import (
	"context"
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	. "github.com/egoholic/blog/config"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/DATA-DOG/godog/gherkin"
	. "github.com/egoholic/blog/lib/store/seed"
	rubricPreviewingRepo "github.com/egoholic/blog/rubric/previewing/repository/postgresql"
	"github.com/sclevine/agouti"
)

var (
	logger = log.New(os.Stdout, "blog-test", 0)
	page   *agouti.Page
	err    error
	driver = agouti.ChromeDriver()
	opt    = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "cucumber",
	}
	cmd = exec.Command("go", "run", "targets/web/main.go", "-logpath", "tmp/log/test.log", "-port", strconv.Itoa(Port), "-dbname", "stoa_blogging_test_acceptance", "-pidpath", "tmp/pids/web.pid")
)

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
	DB, err = sql.Open("postgres", DBConnectionString)
	if err != nil {
		panic(err)
	}
}
func blogHadTheFollowingRubrics(rubrics *gherkin.DataTable) error {
	rubricsToInsert := make([]*Tuple, len(rubrics.Rows)-1)
	header := rubrics.Rows[0].Cells
	for i, row := range rubrics.Rows[1:] {
		attrs := map[string]string{}
		values := row.Cells
		for attrIdx, attrName := range header {
			attrs[attrName.Value] = values[attrIdx].Value
		}
		rubricsToInsert[i] = Must(NewRubric(attrs))
	}
	return InsertList(rubricsToInsert)
}
func theBlogHadTheFollowingAuthors(authors *gherkin.DataTable) error {
	authorsToInsert := make([]*Tuple, len(authors.Rows)-1)
	header := authors.Rows[0].Cells
	for i, row := range authors.Rows[1:] {
		attrs := map[string]string{}
		values := row.Cells
		for attrIdx, attrName := range header {
			attrs[attrName.Value] = values[attrIdx].Value
		}
		authorsToInsert[i] = Must(NewAccount(attrs))
	}
	return InsertList(authorsToInsert)
}
func iVisitedAuthorPage(login string) error {
	expectedURL := fmt.Sprintf("http://localhost:%d/a/%s", Port, login)
	err := page.Navigate(expectedURL)
	if err != nil {
		return err
	}
	url, err := page.URL()
	if err != nil {
		return err
	}
	if url != expectedURL {
		return fmt.Errorf("expected page URL: '%s', got: '%s'", expectedURL, url)
	}
	return nil
}
func iSawAuthor(fullName string) error {
	pageHeaderElems, err := page.AllByID("bhv-author__full_name").Elements()
	if err != nil {
		return err
	}
	if len(pageHeaderElems) == 0 {
		return errors.New("header element not found")
	}
	pageHeader, err := pageHeaderElems[0].GetText()
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if pageHeader != fullName {
		return fmt.Errorf("wrong page header, expected: '%s', got: '%s'", fullName, pageHeader)
	}
	return nil
}
func blogHadTheFollowingPublications(publications *gherkin.DataTable) error {
	ln := len(publications.Rows) - 1
	publicationsToInsert := make([]*Tuple, ln)
	publicationAuthorsToInsert := []*Tuple{}
	header := publications.Rows[0].Cells
	for i, prow := range publications.Rows[1:] {
		attrs := map[string]string{}
		values := prow.Cells
		for attrIdx, attrName := range header {
			attrs[attrName.Value] = values[attrIdx].Value
		}
		loginsStr, ok := attrs["author_logins"]
		if ok {
			logins := strings.Split(loginsStr, ", ")
			delete(attrs, "author_logins")
			slug := attrs["slug"]
			for _, al := range logins {
				authorAttrs := map[string]string{}
				authorAttrs["publication_slug"] = slug
				authorAttrs["author_login"] = al
				publicationAuthorsToInsert = append(publicationAuthorsToInsert, Must(NewPublicationAuthor(authorAttrs)))
			}
		}
		publicationsToInsert[i] = Must(NewPublication(attrs))
	}
	err := InsertList(publicationsToInsert)
	if err != nil {
		return err
	}
	return InsertList(publicationAuthorsToInsert)
}
func iVisitedTheHomePage() (err error) {
	rubricPreviewingRepo.New(context.TODO(), DB, logger)
	expected := fmt.Sprintf("http://localhost:%d/", Port)
	err = page.Navigate(expected)
	if err != nil {
		return
	}
	url, err := page.URL()
	if err != nil {
		return
	}
	if url != expected {
		return fmt.Errorf("expected: `%s`, got: `%s`", expected, url)
	}
	selector := page.Find(".bhv-main-title")
	text, err := selector.Text()
	if err != nil {
		return err
	}
	if text != "BLOG" {
		return fmt.Errorf("expected blog title to be: `%s` got: `%s`\n", "BLOG", text)
	}
	cnt, err := selector.Count()
	if err != nil {
		return
	}
	if cnt != 1 {
		return fmt.Errorf("expected to find 1, got: %d", cnt)
	}
	return
}
func iSawTheFollowingRecentPublications(publications *gherkin.DataTable) error {
	recentSelectors := page.All("#ns-recent .bhv-publication__title")
	elements, err := recentSelectors.Elements()
	if err != nil {
		return err
	}
	rows := publications.Rows[1:]
	ln := len(elements)
	expectedLn := len(rows)
	if ln != expectedLn {
		return fmt.Errorf("expected: '%d' publications, got: '%d'", expectedLn, ln)
	}
	for i, row := range rows {
		elem := elements[i]
		href, err := elem.GetAttribute("href")
		if err != nil {
			return err
		}
		expectedHref := fmt.Sprintf("http://localhost:%d/p/%s", Port, row.Cells[0].Value)
		if href != expectedHref {
			return fmt.Errorf("Expected 'href' attribute to be equal: '%s', got: '%s'", expectedHref, href)
		}
		linkText, err := elem.GetText()
		if err != nil {
			return err
		}
		expectedText := row.Cells[1].Value
		if linkText != expectedText {
			return fmt.Errorf("expected link text to be '%s', got: '%s'", linkText, expectedText)
		}
	}
	return nil
}
func iSawTheFollowingMostPopularPublications(publications *gherkin.DataTable) error {
	popularSelectors := page.All("#ns-popular .bhv-publication__title")
	elements, err := popularSelectors.Elements()
	if err != nil {
		return err
	}
	rows := publications.Rows[1:]
	ln := len(elements)
	expectedLn := len(rows)
	if ln != expectedLn {
		return fmt.Errorf("expected: '%d' publications, got: '%d'", expectedLn, ln)
	}
	for i, row := range rows {
		elem := elements[i]
		href, err := elem.GetAttribute("href")
		if err != nil {
			return err
		}
		expectedHref := fmt.Sprintf("http://localhost:%d/p/%s", Port, row.Cells[0].Value)
		if href != expectedHref {
			return fmt.Errorf("Expected 'href' attribute to be equal: '%s', got: '%s'", expectedHref, href)
		}
		linkText, err := elem.GetText()
		if err != nil {
			return err
		}
		expectedText := row.Cells[1].Value
		if linkText != expectedText {
			return fmt.Errorf("expected link text to be '%s', got: '%s'", linkText, expectedText)
		}
	}
	return nil
}
func iSawTheFollowingRubrics(rubrics *gherkin.DataTable) error {
	rubricSelectors := page.All(".bhv-rubric__title")
	elements, err := rubricSelectors.Elements()
	if err != nil {
		return err
	}
	rows := rubrics.Rows[1:]
	ln := len(elements)
	expectedLn := len(rows)
	if ln != expectedLn {
		return fmt.Errorf("expected: '%d' rubrics, got: '%d'", expectedLn, ln)
	}
	for i, row := range rubrics.Rows[1:] {
		elem := elements[i]
		href, err := elem.GetAttribute("href")
		if err != nil {
			return err
		}
		expectedHref := fmt.Sprintf("http://localhost:%d/r/%s", Port, row.Cells[0].Value)
		if href != expectedHref {
			return fmt.Errorf("Expected 'href' attribute to be equal: '%s', got: '%s'", expectedHref, href)
		}
		linkText, err := elem.GetText()
		if err != nil {
			return err
		}
		expectedText := row.Cells[1].Value
		if linkText != expectedText {
			return fmt.Errorf("expected link text to be '%s', got: '%s'", linkText, expectedText)
		}
	}
	return nil
}
func iVisitedRubricPage(slug string) error {
	expectedURL := fmt.Sprintf("http://localhost:%d/r/%s", Port, slug)
	err = page.Navigate(expectedURL)
	if err != nil {
		return err
	}
	url, err := page.URL()
	if err != nil {
		return err
	}
	if url != expectedURL {
		return fmt.Errorf("expected to visit: '%s', visited: '%s'", expectedURL, url)
	}
	return nil
}
func iSawRubric(title string) error {
	elems, err := page.Find("#bhv-rubric__title").Elements()
	if err != nil {
		return err
	}
	elem := elems[0]

	t, err := elem.GetText()
	if err != nil {
		return err
	}
	if t != title {
		return fmt.Errorf("Expected rubric title: '%s', got: '%s'", title, t)
	}
	return nil
}
func iSawTheFollowingPublications(publications *gherkin.DataTable) error {
	selector := page.All(".bhv-publication__title")
	elements, err := selector.Elements()
	if err != nil {
		return err
	}
	ln := len(elements)
	rows := publications.Rows[1:]
	expectedLn := len(rows)
	if ln != expectedLn {
		return fmt.Errorf("wrong number of publications, expected: '%d', got: '%d'", expectedLn, ln)
	}
	for i, row := range rows {
		elem := elements[i]
		href, err := elem.GetAttribute("href")
		if err != nil {
			return err
		}
		expectedHref := fmt.Sprintf("http://localhost:%d/p/%s", Port, row.Cells[0].Value)
		if href != expectedHref {
			return fmt.Errorf("Expected 'href' attribute to be equal: '%s', got: '%s'", expectedHref, href)
		}
		linkText, err := elem.GetText()
		if err != nil {
			return err
		}
		expectedText := row.Cells[1].Value
		if linkText != expectedText {
			return fmt.Errorf("expected link text to be '%s', got: '%s'", linkText, expectedText)
		}
	}
	return nil
}
func iVisitedPublicationPage(slug string) error {
	expectedURL := fmt.Sprintf("http://localhost:%d/p/%s", Port, slug)
	err := page.Navigate(expectedURL)
	if err != nil {
		return err
	}
	pageURL, err := page.URL()
	if err != nil {
		return err
	}
	if pageURL != expectedURL {
		return fmt.Errorf("expected page url to be: '%s', got: '%s'", expectedURL, pageURL)
	}
	return nil
}
func iReadPublication(expectedTitle string) error {
	row := DB.QueryRow("SELECT title, slug, content FROM publications WHERE title = $1 LIMIT 1;", expectedTitle)
	var publication struct{ Title, Slug, Content string }
	err := row.Scan(&publication.Title, &publication.Slug, &publication.Content)
	if err != nil {
		return err
	}
	title, err := page.FindByID("bhv-publication__title").Text()
	if err != nil {
		return err
	}
	if title != expectedTitle {
		return fmt.Errorf("expected title: '%s', got: '%s'", expectedTitle, title)
	}
	content, err := page.FindByID("bhv-publication__content").Text()
	if err != nil {
		return err
	}
	if content != publication.Content {
		return fmt.Errorf("expected content: '%s', got: '%s'", publication.Content, content)
	}
	return nil
}
func iSeeThatPageNotFound() error {
	selector := page.AllByID("bhv-system-page-not-found")
	elements, err := selector.Elements()
	if err != nil {
		return err
	}
	if len(elements) == 0 {
		return fmt.Errorf("expected to find: '%s', but element not found", selector.String())
	}
	header := elements[0]
	text, err := header.GetText()
	if err != nil {
		return err
	}

	expectedText := "PAGE NOT FOUND"
	if text != expectedText {
		return fmt.Errorf("expected text: '%s', got: '%s'", expectedText, text)
	}
	return nil
}
func FeatureContext(s *godog.Suite) {
	s.BeforeSuite(func() {
		err = driver.Start()
		if err != nil {
			logger.Printf("Error: Can't run driver: %s\n", err.Error())
		}
		logger.Println("Driver runned!")
		page, err = driver.NewPage(agouti.Browser("firefox"))
		if err != nil {
			logger.Printf("Error: Can't run client: %s\n", err.Error())
		}
		logger.Println("Client runned!")
		err = cmd.Start()
		if err != nil {
			logger.Printf("Error: Can't run blog server: %s\n", err.Error())
		}
	})
	s.AfterSuite(func() {
		err = driver.Stop()
		if err != nil {
			logger.Println("Error: ", err.Error())
		}
		err = stopBlogApp()
		if err != nil {
			logger.Println("Error: ", err.Error())
		}
	})
	s.BeforeScenario(func(interface{}) {
		err = Truncate("accounts", "rubrics", "publications", "publication_authors")
		if err != nil {
			logger.Printf("Error: Can't clean up DB: %s\n", err.Error())
		}
		page, err = driver.NewPage(agouti.Browser("firefox"))
		if err != nil {
			logger.Printf("Error: Can't run client: %s\n", err.Error())
		}
	})
	s.AfterScenario(func(_ interface{}, err error) {
		if err != nil {
			url, err2 := page.URL()
			if err2 != nil {
				url = ""
			}
			url = strings.ReplaceAll(url, "/", "--")
			page.Screenshot(fmt.Sprintf("tmp/screenshots/screenshot-%d[%s].png", time.Now().Unix(), url))
			logger.Println("Error: ", err.Error())
		}
		err = page.Destroy()
		if err != nil {
			logger.Println("Error: ", err.Error())
		}
	})

	s.Step(`^the blog had the following rubrics:$`, blogHadTheFollowingRubrics)
	s.Step(`^the blog had the following publications:$`, blogHadTheFollowingPublications)
	s.Step(`^I visited the home page$`, iVisitedTheHomePage)
	s.Step(`^I saw the following recent publications:$`, iSawTheFollowingRecentPublications)
	s.Step(`^I saw the following most popular publications:$`, iSawTheFollowingMostPopularPublications)
	s.Step(`^I saw the following rubrics:$`, iSawTheFollowingRubrics)
	s.Step(`^I visited "([^"]*)" rubric page$`, iVisitedRubricPage)
	s.Step(`^I saw the following publications:$`, iSawTheFollowingPublications)
	s.Step(`^I visited "([^"]*)" publication page$`, iVisitedPublicationPage)
	s.Step(`^I read "([^"]*)" publication$`, iReadPublication)
	s.Step(`^the blog had the following authors:$`, theBlogHadTheFollowingAuthors)
	s.Step(`^I visited "([^"]*)" author page$`, iVisitedAuthorPage)
	s.Step(`^I saw "([^"]*)" author$`, iSawAuthor)
	s.Step(`^I saw the following publications:$`, iSawTheFollowingPublications)
	s.Step(`^I see that page not found$`, iSeeThatPageNotFound)
	s.Step(`^I saw "([^"]*)" rubric$`, iSawRubric)
}

func stopBlogApp() (err error) {
	fmt.Println("stopping blog app")
	file, err := os.Open(PIDFilePath)
	if err != nil {
		return
	}
	pidStr := make([]byte, 16)
	_, err = file.Read(pidStr)
	if err != nil {
		return
	}
	pid, err := strconv.Atoi(strings.TrimRight(string(pidStr), "\x00"))
	if err != nil {
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return
	}
	err = process.Kill()
	if err != nil {
		return
	}
	fmt.Println("pid file dropped")
	return
}
