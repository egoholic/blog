package seed

import (
	"database/sql"
	"fmt"

	. "github.com/egoholic/blog/config"
	_ "github.com/lib/pq"
)

var (
	__LAST_PUBLICATION_ID = 0
	__LAST_RUBRIC_ID      = 0
)

func Truncate(db *sql.DB, names ...string) (err error) {
	for _, name := range names {
		_, err = db.Exec(fmt.Sprintf("TRUNCATE %s RESTART IDENTITY;", name))
		if err != nil {
			return
		}
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
	query := fmt.Sprintf(`INSERT INTO publications (slug,             meta_keywords,     meta_description,   title,              content,              created_at) VALUES
																		             ('publication-%d', 'publication, %d', '%dth publication', '%dth PUBLICATION', 'My %d publication.', CURRENT_DATE + INTERVAL '%d day' - INTERVAL '1000 day');`, __LAST_PUBLICATION_ID, __LAST_PUBLICATION_ID, __LAST_PUBLICATION_ID, __LAST_PUBLICATION_ID, __LAST_PUBLICATION_ID, __LAST_PUBLICATION_ID)
	__LAST_PUBLICATION_ID = __LAST_PUBLICATION_ID + 1
	return db.Exec(query)
}
func CreateRubric(db *sql.DB) (result sql.Result, err error) {
	query := fmt.Sprintf(`INSERT INTO rubrics (slug,   meta_keywords, meta_description, title,  description) VALUES
															              ('%dth', 'rubric, %d',  '%d Rubric',      '%dTH', 'My %dth rubric.');`, __LAST_RUBRIC_ID, __LAST_RUBRIC_ID, __LAST_RUBRIC_ID, __LAST_RUBRIC_ID, __LAST_RUBRIC_ID)
	__LAST_RUBRIC_ID = __LAST_RUBRIC_ID + 1
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
	return
}
