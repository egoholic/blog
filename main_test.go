package main

import (
	"context"
	"database/sql"
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
	"github.com/sclevine/agouti/api"
)

var (
	logger = log.New(os.Stdout, "-> ", 0)
	page   *agouti.Page
	err    error
	driver = agouti.ChromeDriver()
	opt    = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "cucumber",
	}
	cmd = exec.Command("go", "run", "targets/web/main.go", "-logpath", "./test.log", "-port", strconv.Itoa(Port), "-dbname", "stoa_blogging_test_acceptance")
)

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
	DB, err = sql.Open("postgres", DBConnectionString)
	if err != nil {
		panic(err)
	}
}

func thereWasABlog() (err error) {
	return cmd.Start()
}

func blogHadTheFollowingRubrics(rubrics *gherkin.DataTable) error {
	rubricsToInsert := make([]*Tuple, len(rubrics.Rows)-1)
	header := rubrics.Rows[0].Cells
	for i, rrow := range rubrics.Rows[1:] {
		attrs := map[string]string{}
		values := rrow.Cells
		for attrIdx, attrName := range header {
			attrs[attrName.Value] = values[attrIdx].Value
		}
		rubricsToInsert[i] = Must(NewRubric(attrs))
	}
	return InsertList(rubricsToInsert)
}

func blogHadTheFollowingPublications(publications *gherkin.DataTable) error {
	publicationsToInsert := make([]*Tuple, len(publications.Rows)-1)
	header := publications.Rows[0].Cells
	for i, prow := range publications.Rows[1:] {
		attrs := map[string]string{}
		values := prow.Cells
		for attrIdx, attrName := range header {
			attrs[attrName.Value] = values[attrIdx].Value
		}
		publicationsToInsert[i] = Must(NewPublication(attrs))
	}
	return InsertList(publicationsToInsert)
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
	elements, err := page.AllByClass("bhv-recent-publication").Elements()
	if err != nil {
		return err
	}
	for i, row := range publications.Rows[1:] {
		elem := elements[i]
		aElem, err := elem.GetElement(api.Selector{Using: "css selector", Value: "a"})
		if err != nil {
			return err
		}
		href, err := aElem.GetAttribute("href")
		if err != nil {
			return err
		}
		expectedHref := fmt.Sprintf("http://localhost:%d/p/%s", Port, row.Cells[0].Value)
		if href != expectedHref {
			return fmt.Errorf("Expected 'href' attribute to be equal: '%s', got: '%s'", expectedHref, href)
		}
		linkText, err := aElem.GetText()
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
	elements, err := page.AllByClass("bhv-popular-publication").Elements()
	if err != nil {
		return err
	}
	for i, row := range publications.Rows[1:] {
		elem := elements[i]
		aElem, err := elem.GetElement(api.Selector{Using: "css selector", Value: "a"})
		if err != nil {
			return err
		}
		href, err := aElem.GetAttribute("href")
		if err != nil {
			return err
		}
		expectedHref := fmt.Sprintf("http://localhost:%d/p/%s", Port, row.Cells[0].Value)
		if href != expectedHref {
			return fmt.Errorf("Expected 'href' attribute to be equal: '%s', got: '%s'", expectedHref, href)
		}
		linkText, err := aElem.GetText()
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
	elements, err := page.AllByClass("bhv-rubric").Elements()
	if err != nil {
		return err
	}
	for i, row := range rubrics.Rows[1:] {
		elem := elements[i]
		aElem, err := elem.GetElement(api.Selector{Using: "css selector", Value: "a"})
		if err != nil {
			return err
		}
		href, err := aElem.GetAttribute("href")
		if err != nil {
			return err
		}
		expectedHref := fmt.Sprintf("http://localhost:%d/r/%s", Port, row.Cells[0].Value)
		if href != expectedHref {
			return fmt.Errorf("Expected 'href' attribute to be equal: '%s', got: '%s'", expectedHref, href)
		}
		linkText, err := aElem.GetText()
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

func iVisitedRubricPage(title string) error {
	r := DB.QueryRow("SELECT r.slug FROM (SELECT slug, title FROM rubrics WHERE title = $1 LIMIT 1) AS r LIMIT 1;", title)
	var slug string
	err = r.Scan(&slug)
	if err != nil {
		return err
	}
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
	elems, err := page.FindByID("bhv-rubric-title").Elements()
	if err != nil {
		return err
	}
	rubric := elems[0]
	t, err := rubric.GetText()
	if err != nil {
		return err
	}
	if t != title {
		return fmt.Errorf("Expected rubric title: '%s', got: '%s'", title, t)
	}
	return nil
}
func iSawTheFollowingPublications(publications *gherkin.DataTable) error {
	elements, err := page.AllByClass("bhv-publication-preview").Elements()
	if err != nil {
		return err
	}
	for i, row := range publications.Rows[1:] {
		elem := elements[i]
		aElem, err := elem.GetElement(api.Selector{Using: "css selector", Value: "a"})
		if err != nil {
			return err
		}
		href, err := aElem.GetAttribute("href")
		if err != nil {
			return err
		}
		expectedHref := fmt.Sprintf("http://localhost:%d/p/%s", Port, row.Cells[0].Value)
		if href != expectedHref {
			return fmt.Errorf("Expected 'href' attribute to be equal: '%s', got: '%s'", expectedHref, href)
		}
		linkText, err := aElem.GetText()
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

func FeatureContext(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		err = driver.Start()
		if err != nil {
			logger.Fatalf("Can't run driver: %s\n", err.Error())
		}
		logger.Println("Driver runned!")
		page, err = driver.NewPage(agouti.Browser("firefox"))
		if err != nil {
			logger.Fatalf("Can't run client: %s\n", err.Error())
		}
		logger.Println("Client runned!")
		err = Truncate("accounts", "rubrics", "publications", "publication_authors")
		if err != nil {
			logger.Fatalf("Can't clean up DB: %s\n", err.Error())
		}
	})
	s.AfterScenario(func(_ interface{}, err error) {
		driver.Stop()
		page.Destroy()
		if err != nil {
			page.Screenshot(fmt.Sprintf("tmp/screenshots/screenshot-%d.png", time.Now().Unix()))
			logger.Fatal(err.Error())
		}
		err = stopBlogApp()
		if err != nil {
			logger.Fatal(err.Error())
		}
	})

	s.Step(`there was a blog`, thereWasABlog)
	s.Step(`^the blog had the following rubrics:$`, blogHadTheFollowingRubrics)
	s.Step(`^the blog had the following publications:$`, blogHadTheFollowingPublications)

	s.Step(`^I visited the home page$`, iVisitedTheHomePage)
	s.Step(`^I saw the following recent publications:$`, iSawTheFollowingRecentPublications)
	s.Step(`^I saw the following most popular publications:$`, iSawTheFollowingMostPopularPublications)
	s.Step(`^I saw the following rubrics:$`, iSawTheFollowingRubrics)
	s.Step(`^I visited "([^"]*)" rubric page$`, iVisitedRubricPage)
	s.Step(`^I saw the following publications:$`, iSawTheFollowingPublications)
}

func stopBlogApp() (err error) {
	fmt.Println("stopping blog app")
	file, err := os.Open("blog-web.pid")
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
