package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/sclevine/agouti"
	. "github.com/sclevine/agouti/matchers"
)

var (
	page    *agouti.Page
	err     error
	command = []string{"java", "-jar", "selenium-server.jar", "-port", ""}
	driver  = agouti.NewWebDriver("http://localhost:8080", command)
	opt     = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "progress", // can define default values
	}
	cmd = exec.Command("go", "run", "../../../targets/web/main.go")
)

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func blogHasPublications() error {
	// return godog.ErrPending
	return nil
}

func iVisitHomePage() error {
	// return godog.ErrPending
	return nil
}

func iSeeRecentPublications(arg1 int) error {
	// return godog.ErrPending
	return nil
}

func iSeeMostPopularPublications(arg1 int) error {
	// return godog.ErrPending
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		if err := cmd.Run(); err != nil {
			log.Fatal(err)
		}
		driver.Start()
		page, err = driver.NewPage(agouti.Browser("firefox"))
	})
	s.AfterScenario(func(interface{}, error) {
		driver.Stop()
		page.Destroy()
	})
	s.Step(`^blog has publications$`, blogHasPublications)
	s.Step(`^I visit home page$`, iVisitHomePage)
	s.Step(`^I see (\d+) recent publications$`, iSeeRecentPublications)
	s.Step(`^I see (\d+) most popular publications$`, iSeeMostPopularPublications)
}
