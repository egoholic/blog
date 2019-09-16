package previewing_test

import (
	"flag"
	"os"
	"testing"

	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"github.com/tebeka/selenium"

	"github.com/llonchj/browsersteps"
)

var (
	Godogs int
	opt    = godog.Options{
		Output: colors.Colored(os.Stdout),
		Format: "progress", // can define default values
	}
)

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}
func FeatureContext(s *godog.Suite) {
	browsersteps.NewBrowserSteps(s,
		selenium.Capabilities{"browserName": "chrome"},
		"")
	s.Step(`^blog has several publications$`, blogHasSeveralPublications)
	s.Step(`^I visit home page$`, visitHomePage)
	s.Step(`^I see (\d+) recent publications`, seeRecentPublications)
	s.Step(`^I see (\d+) most popular publications`, seeTopPublications)
}
func blogHasSeveralPublications() error {
	return nil
}
func visitHomePage() error {
	return nil
}
func seeRecentPublications(_ int) error {
	return nil
}
func seeTopPublications(_ int) error {
	return nil
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
