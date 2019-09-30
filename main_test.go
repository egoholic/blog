package main

import (
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

func thereIsABlog() (err error) {
	return cmd.Start()
}

func blogHasNextRubrics(rubrics *gherkin.DataTable) error {
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

func blogHasNextPublications(publications *gherkin.DataTable) error {
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

func iVisitHomePage() (err error) {
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

func iSeeNextRecentPublications(publications *gherkin.DataTable) error {
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

func iSeeNextMostPopularPublications(publications *gherkin.DataTable) error {
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
		if err != nil {
			page.Screenshot(fmt.Sprintf("tmp/screenshots/screenshot-%d.png", time.Now().Unix()))
			logger.Fatal(err.Error())
		}
		driver.Stop()
		page.Destroy()
		err = stopBlogApp()
		if err != nil {
			logger.Fatal(err.Error())
		}
	})

	s.Step(`there is a blog`, thereIsABlog)
	s.Step(`^the blog has next rubrics:$`, blogHasNextRubrics)
	s.Step(`^the blog has next publications:$`, blogHasNextPublications)

	s.Step(`^I visit home page$`, iVisitHomePage)
	s.Step(`^I see next recent publications:$`, iSeeNextRecentPublications)
	s.Step(`^I see next most popular publications:$`, iSeeNextMostPopularPublications)
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
