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
	//. "github.com/sclevine/agouti/matchers"
)

var (
	logger  = log.New(os.Stdout, "-> ", 0)
	page    *agouti.Page
	err     error
	command = []string{"java", "-jar", "selenium-server.jar", "-port", ""}
	driver  = agouti.ChromeDriver()
	opt     = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "cucumber", // can define default values
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
	return godog.ErrPending
}

func iVisitHomePage() error {
	return godog.ErrPending
}

func iSeeRecentPublications(arg1 int) error {
	return godog.ErrPending
}

func iSeeMostPopularPublications(arg1 int) error {
	return godog.ErrPending
}

func FeatureContext(s *godog.Suite) {
	s.BeforeScenario(func(interface{}) {
		if err := cmd.Start(); err != nil {
			_, err2 := cmd.CombinedOutput()
			logger.Fatalf("Error1: %s\n\tcmd: %s\n", err.Error(), err2.Error())
		}
		logger.Println("ok1")
		err = driver.Start()
		if err != nil {
			logger.Fatalf("Error2: %s\n", err.Error())
		}
		logger.Println("ok2")

		page, err = driver.NewPage(agouti.Browser("firefox"))
		if err != nil {
			logger.Fatalf("Error3: %s\n", err.Error())
		}
		logger.Println("ok3")

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
