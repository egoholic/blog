package seed

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"

	. "github.com/egoholic/blog/config"
	_ "github.com/lib/pq"
)

var (
	__LAST_ID = map[string]int{}
)

func init() {
	__LAST_ID["publications"] = 0
	__LAST_ID["rubrics"] = 0
	__LAST_ID["accounts"] = 0
}

func Truncate(db *sql.DB, names ...string) (err error) {
	for _, name := range names {
		_, err = db.Exec(fmt.Sprintf("TRUNCATE %s RESTART IDENTITY;", name))
		if err != nil {
			return
		}
		__LAST_ID[name] = 0
	}
	return
}
func Many(ntimes int, db *sql.DB, factory func(*sql.DB) (sql.Result, error)) (results []sql.Result, err error) {
	var result sql.Result
	for i := 0; i < ntimes; i++ {
		result, err = factory(db)
		if err != nil {
			return
		}
		results = append(results, result)
	}
	return
}
func CreatePublication(db *sql.DB) (result sql.Result, err error) {
	pid := __LAST_ID["publications"]
	rn := rand.Intn(__LAST_ID["rubrics"])

	query := fmt.Sprintf(`INSERT INTO publications (slug,             meta_keywords,     meta_description,   title,              content,              popularity,  created_at,                                             rubric_slug) VALUES
																		             ('publication-%d', 'publication, %d', '%dth publication', '%dth PUBLICATION', 'My %d publication.', 1,           CURRENT_DATE + INTERVAL '%d day' - INTERVAL '1000 day', '%dth');`, pid, pid, pid, pid, pid, pid, rn)
	__LAST_ID["publications"]++
	return db.Exec(query)
}
func CreateRubric(db *sql.DB) (result sql.Result, err error) {
	rid := __LAST_ID["rubrics"]
	query := fmt.Sprintf(`INSERT INTO rubrics (slug,   meta_keywords, meta_description, title,  description) VALUES
															              ('%dth', 'rubric, %d',  '%d Rubric',      '%dTH', 'My %dth rubric.');`, rid, rid, rid, rid, rid)
	__LAST_ID["rubrics"]++
	return db.Exec(query)
}
func CreateAccount(db *sql.DB) (result sql.Result, err error) {
	rid := __LAST_ID["accounts"]
	query := fmt.Sprintf(`INSERT INTO accounts (login,          first_name,     last_name,     bio) VALUES
															               ('account-%dth', 'Firstname-%d', 'Lastname-%d', 'I am %dth user.');`, rid, rid, rid, rid)
	__LAST_ID["accounts"]++
	return db.Exec(query)
}

func CreatePublicationAuthors(db *sql.DB) (result sql.Result, err error) {
	var sb strings.Builder
	sb.WriteString("INSERT INTO publication_authors (publication_slug, author_login) VALUES ")
	subs := []string{}
	for pi := 0; pi < __LAST_ID["publications"]; pi++ {
		for ui := 0; ui < rand.Intn(__LAST_ID["accounts"])+1; ui++ {
			subs = append(subs, (fmt.Sprintf(`('publication-%d', 'account-%dth')`, pi, ui)))
		}
	}
	sb.WriteString(strings.Join(subs, ",\n"))
	sb.WriteRune(';')
	query := sb.String()
	return db.Exec(query)
}

func Seed() (err error) {
	var (
		publicationsNumber = 20
		rubricsNumber      = 5
		accountsNumber     = 5
	)
	connStr, err := Config.DBCredentials().ConnectionString()
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", connStr)
	defer db.Close()

	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	_, err = Many(rubricsNumber, db, CreateRubric)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("-- %d rubrics created!\n", rubricsNumber)

	_, err = Many(publicationsNumber, db, CreatePublication)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("-- %d publications created!\n", publicationsNumber)

	_, err = Many(accountsNumber, db, CreateAccount)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("-- %d accounts created!\n", accountsNumber)

	_, err = CreatePublicationAuthors(db)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Println("-- publications are linked to authots!")
	return
}
