package seed

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	. "github.com/egoholic/blog/config"
	_ "github.com/lib/pq"
)

func Seed(initCtx context.Context) (err error) {
	connStr, err := Config.DBCredentials().ConnectionString()
	if err != nil {
		return
	}
	db, err := sql.Open("postgres", connStr)
	defer db.Close()
	if err != nil {
		return
	}
	ctx, cancel := context.WithTimeout(initCtx, 10*time.Second)
	defer cancel()

	query := `INSERT INTO publications (slug,            meta_keywords,    meta_description,    title,               content,                 created_at) VALUES
																		 ('publication-1', 'publication, 1', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06'),
																		 ('publication-2', 'publication, 2', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06'),
																		 ('publication-3', 'publication, 3', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06'),
																		 ('publication-4', 'publication, 4', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06'),
																		 ('publication-5', 'publication, 5', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06'),
																		 ('publication-6', 'publication, 6', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06'),
																		 ('publication-7', 'publication, 7', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06'),
																		 ('publication-8', 'publication, 8', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06'),
																		 ('publication-9', 'publication, 9', 'First publication', 'FIRST PUBLICATION', 'My first publication.', '2019-06-06');`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return
	}
	fmt.Println("-- table `publications` has been populated")

	query = `INSERT INTO rubrics (slug,       meta_keywords, meta_description, title,          description) VALUES
															 ('rubric-1', 'rubric, 1',   'First Rubric',  'FIRST RUBRIC', 'My first rubric.'),
															 ('rubric-2', 'rubric, 2',   'First Rubric',  'FIRST RUBRIC', 'My first rubric.'),
															 ('rubric-3', 'rubric, 3',   'First Rubric',  'FIRST RUBRIC', 'My first rubric.'),
															 ('rubric-4', 'rubric, 4',   'First Rubric',  'FIRST RUBRIC', 'My first rubric.'),
															 ('rubric-5', 'rubric, 5',   'First Rubric',  'FIRST RUBRIC', 'My first rubric.'),
															 ('rubric-6', 'rubric, 6',   'First Rubric',  'FIRST RUBRIC', 'My first rubric.'),
															 ('rubric-7', 'rubric, 7',   'First Rubric',  'FIRST RUBRIC', 'My first rubric.'),
															 ('rubric-8', 'rubric, 8',   'First Rubric',  'FIRST RUBRIC', 'My first rubric.');`
	_, err = db.ExecContext(ctx, query)
	if err != nil {
		return
	}
	fmt.Println("-- table `rubrics` has been populated")
	return
}
