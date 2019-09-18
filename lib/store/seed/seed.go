package seed

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	repository "github.com/egoholic/blog/blog/previewing/repository/postgresql"

	. "github.com/egoholic/blog/config"
	_ "github.com/lib/pq"
)

var (
	__LAST_ID = map[string]int{}
)

func init() {
	__LAST_ID["publications"] = 0
	__LAST_ID["rubrics"] = 0
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
	query := fmt.Sprintf(`INSERT INTO publications (slug,             meta_keywords,     meta_description,   title,              content,              popularity,  created_at) VALUES
																		             ('publication-%d', 'publication, %d', '%dth publication', '%dth PUBLICATION', 'My %d publication.', 1,           CURRENT_DATE + INTERVAL '%d day' - INTERVAL '1000 day');`, pid, pid, pid, pid, pid, pid)
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
func Seed() (err error) {
	var (
		publicationsNumber = 20
		rubricsNumber      = 5
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
	_, err = Many(publicationsNumber, db, CreatePublication)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("-- %d publications created!\n", publicationsNumber)
	_, err = Many(rubricsNumber, db, CreateRubric)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}
	fmt.Printf("-- %d rubrics created!\n", rubricsNumber)

	logger := log.New(os.Stdout, "blog", 0)

	repo := repository.New(context.TODO(), db, logger)

	pop, err := repo.PopularPublications()
	fmt.Printf("\n\n%#v\n\terr:%s\n", pop, err)
	rec, err := repo.PopularPublications()
	fmt.Printf("\n\n%#v\n\terr:%s\n", rec, err)
	return
}
