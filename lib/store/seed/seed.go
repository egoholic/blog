package seed

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type Tuple struct {
	Fields    map[string]string
	TableName string
}

func FieldsFor(tn string) []string {
	switch tn {
	case "accounts":
		return AccountFieldNames
	case "rubrics":
		return RubricFieldNames
	case "publications":
		return PublicationFieldNames
	case "publication_authors":
		return PublicationAuthorFieldNames
	}
	return []string{}
}

var (
	DB               *sql.DB
	Random           *rand.Rand
	logins           []string
	rubricSlugs      []string
	publicationSlugs []string

	AccountFieldNames           = []string{"login", "first_name", "last_name", "bio"}
	RubricFieldNames            = []string{"slug", "meta_keywords", "meta_description", "title", "description"}
	PublicationFieldNames       = []string{"slug", "meta_keywords", "meta_description", "title", "content", "created_at", "rubric_slug", "popularity"}
	PublicationAuthorFieldNames = []string{"publication_slug", "author_login"}
	Sentences                   = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
		"Dolor sed viverra ipsum nunc aliquet bibendum enim.",
		"In massa tempor nec feugiat.",
		"Nunc aliquet bibendum enim facilisis gravida.",
		"Nisl nunc mi ipsum faucibus vitae aliquet nec ullamcorper.",
		"Amet luctus venenatis lectus magna fringilla.",
		"Volutpat maecenas volutpat blandit aliquam etiam erat velit scelerisque in.",
		"Egestas egestas fringilla phasellus faucibus scelerisque eleifend.",
		"Sagittis orci a scelerisque purus semper eget duis.",
		"Nulla pharetra diam sit amet nisl suscipit.",
		"Sed adipiscing diam donec adipiscing tristique risus nec feugiat in.",
		"Fusce ut placerat orci nulla.",
		"Pharetra vel turpis nunc eget lorem dolor.",
		"Tristique senectus et netus et malesuada.",
		"Etiam tempor orci eu lobortis elementum nibh tellus molestie.",
		"Neque egestas congue quisque egestas.",
		"Egestas integer eget aliquet nibh praesent tristique.",
		"Vulputate mi sit amet mauris.",
		"Sodales neque sodales ut etiam sit.",
		"Dignissim suspendisse in est ante in.",
		"Volutpat commodo sed egestas egestas.",
		"Felis donec et odio pellentesque diam.",
		"Pharetra vel turpis nunc eget lorem dolor sed viverra.",
		"Porta nibh venenatis cras sed felis eget.",
		"Aliquam ultrices sagittis orci a.",
		"Dignissim diam quis enim lobortis.",
		"Aliquet porttitor lacus luctus accumsan.",
		"Dignissim convallis aenean et tortor at risus viverra adipiscing at.",
	}
)

func init() {
	logins = []string{}
	rubricSlugs = []string{}
	publicationSlugs = []string{}
	Random = rand.New(rand.NewSource(time.Now().Unix()))
}

func has(collection []string, v string) bool {
	for _, elem := range collection {
		if elem == v {
			return true
		}
	}
	return false
}

var firstNames = []string{"Aaron", "Robert", "Rob", "Richard", "Rich", "Rick", "Dirk", "Kirk", "Thomas", "Derek", "Samuel", "Sam", "Sammy", "Kennet", "Peter", "Rodger", "Rodrigo", "Ivan", "Mark", "Kventin", "Oleg", "Andrey", "Sergey", "Vladimir", "Volodymyr", "Voldemar", "Ulrih", "Rodrigo", "Esteban", "Gielermo", "Francis", "Frank", "Kristoph", "Ann", "Caren", "Julia", "Anastatia", "Margaret", "Sally", "John", "Joanna"}

func CreatedAt() string {
	t := time.Now()
	y, m, d := t.Date()
	h := t.Hour()
	M := t.Minute()
	s := t.Second()
	return fmt.Sprintf("%d-%d-%d %d:%d:%d", y, m, d, h, M, s)
}

func Sentence() string {
	return Sentences[Random.Intn(len(Sentences)-1)]
}
func Paragraph() string {
	var (
		sb strings.Builder
		n  = Random.Intn(6)
	)
	for i := 0; i < n-1; i++ {
		sb.WriteString(Sentence())
		sb.WriteRune(' ')
	}
	sb.WriteString(Sentence())

	return sb.String()
}
func Paragraphs(n int) string {
	var sb strings.Builder
	for i := 0; i < n-1; i++ {
		sb.WriteString(Paragraph())
		sb.WriteRune('\n')
	}
	sb.WriteString(Paragraph())
	return sb.String()
}

func FirstName() string {
	return firstNames[Random.Intn(len(firstNames)-1)]
}

var lastNames = []string{"Peterson", "Johnson", "Falcon", "Black", "White", "Brown", "Silver", "Gold", "Golt", "Colt", "Kaas", "Ivanov", "Sidorov", "Petrov", "Vernidub", "Melnik", "Melnyk", "Melnychenko", "Marchanko", "Petrenko", "Shevchenko", "Washindton", "Miller"}

func LastName() string {
	return lastNames[Random.Intn(len(lastNames)-1)]
}

func LoginFor(firstName, lastName string) string {
	login := strings.ToLower(fmt.Sprintf("%s.%s-%d", firstName, lastName, len(logins)))
	logins = append(logins, login)
	return login
}

func Login() string {
	return logins[Random.Intn(len(logins)-1)]
}

func Bio() string {
	return Paragraphs(1 + Random.Intn(2))
}

var publicationTitles = []string{"How to write great articles", "Best practices for content marketing", "10 secrets of attraction", "Introcuction to Content Marketing", "How to avoid problems", "Business drivers", "The art of focusing", "Content Marketing trends", "Top blogers' secrets", "How to Sell", "How to improve relations with your clients", "Blogging Templates", "Great Story"}

func PublicationTitle() string {
	return publicationTitles[Random.Intn(len(publicationTitles)-1)]
}
func PublicationContent() string {
	return Paragraphs(5 + Random.Intn(5))
}
func Popularity() string {
	return fmt.Sprintf("%d", Random.Intn(100000))
}

var rubricTitles = []string{"HowTos", "Interviews", "Feature Releases", "Best Practices", "Retrospectivas", "Opinions", "Reports"}

func RubricTitle() string {
	return rubricTitles[Random.Intn(len(rubricTitles)-1)]
}

func RubricDescription() string {
	return Paragraphs(1 + Random.Intn(2))
}

func slug(s string, n int) string {
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ReplaceAll(s, "'", "-")
	s = strings.ReplaceAll(s, "_", "-")
	s = strings.ReplaceAll(s, "/", "-or-")
	s = strings.ReplaceAll(s, "&", "-and-")
	s = strings.ReplaceAll(s, "?", "")
	s = strings.ReplaceAll(s, "=", "-is-")
	s = strings.ReplaceAll(s, "\"", "-")
	return fmt.Sprintf("%s-%d", strings.ToLower(s), n)
}

func PublicationSlugFor(t string) string {
	s := slug(t, len(publicationSlugs))
	publicationSlugs = append(publicationSlugs, s)
	return s
}

func PublicationSlug() string {
	return publicationSlugs[Random.Intn(len(publicationSlugs)-1)]
}

func RubricSlugFor(t string) string {
	s := slug(t, len(rubricSlugs))
	rubricSlugs = append(rubricSlugs, s)
	return s
}

func RubricSlug() string {
	l := len(rubricSlugs)
	if l == 0 {
		return ""
	}
	return rubricSlugs[Random.Intn(l-1)]
}

var keywords = []string{"content", "marketing", "blogging", "best-practices", "business", "company", "product", "secret", "sale", "recommendation", "idea"}

func MetaKeywords() string {
	acc := []string{}
	l := len(keywords)
	for i := 0; i < 3; i++ {
		acc = append(acc, keywords[Random.Intn(l-1)])
	}
	return strings.Join(acc, ", ")
}

func MetaDescription() string {
	return Sentence()
}

func NewAccount(fields map[string]string) (*Tuple, error) {
	if _, ok := fields["first_name"]; !ok {
		fields["first_name"] = FirstName()
	}
	fields["first_name"] = clean(fields["first_name"])
	if _, ok := fields["last_name"]; !ok {
		fields["last_name"] = LastName()
	}
	fields["last_name"] = clean(fields["last_name"])
	if _, ok := fields["login"]; !ok {
		fields["login"] = LoginFor(fields["first_name"], fields["last_name"])
	}
	fields["login"] = clean(fields["login"])
	if _, ok := fields["bio"]; !ok {
		fields["bio"] = Bio()
	}
	fields["bio"] = clean(fields["bio"])
	return new("accounts", fields)
}

func NewRubric(fields map[string]string) (*Tuple, error) {
	if _, ok := fields["title"]; !ok {
		fields["title"] = RubricTitle()
	}
	fields["title"] = clean(fields["title"])
	// should be after title assignment
	if _, ok := fields["slug"]; !ok {
		fields["slug"] = RubricSlugFor(fields["title"])
	}
	fields["slug"] = clean(fields["slug"])
	if _, ok := fields["meta_keywords"]; !ok {
		fields["meta_keywords"] = MetaKeywords()
	}
	fields["meta_keywords"] = clean(fields["meta_keywords"])
	if _, ok := fields["meta_description"]; !ok {
		fields["meta_description"] = MetaDescription()
	}
	fields["meta_description"] = clean(fields["meta_description"])
	if _, ok := fields["description"]; !ok {
		fields["description"] = RubricDescription()
	}
	fields["description"] = clean(fields["description"])
	return new("rubrics", fields)
}

func clean(v string) string {
	return strings.ReplaceAll(v, "'", `''`)
}

func NewPublication(fields map[string]string) (*Tuple, error) {
	if _, ok := fields["title"]; !ok {
		fields["title"] = PublicationTitle()
	}
	fields["title"] = clean(fields["title"])
	// should be after title assignment
	if _, ok := fields["slug"]; !ok {
		fields["slug"] = PublicationSlugFor(fields["title"])
	}
	fields["slug"] = clean(fields["slug"])
	if _, ok := fields["meta_keywords"]; !ok {
		fields["meta_keywords"] = MetaKeywords()
	}
	fields["meta_keywords"] = clean(fields["meta_keywords"])
	if _, ok := fields["meta_description"]; !ok {
		fields["meta_description"] = MetaDescription()
	}
	fields["meta_description"] = clean(fields["meta_description"])
	if _, ok := fields["content"]; !ok {
		fields["content"] = PublicationContent()
	}
	fields["content"] = clean(fields["content"])
	if _, ok := fields["created_at"]; !ok {
		fields["created_at"] = CreatedAt()
	}
	fields["created_at"] = clean(fields["created_at"])
	if _, ok := fields["rubric_slug"]; !ok {
		fields["rubric_slug"] = RubricSlug()
	}
	fields["rubric_slug"] = clean(fields["rubric_slug"])
	if _, ok := fields["popularity"]; !ok {
		fields["popularity"] = Popularity()
	}
	fields["popularity"] = clean(fields["popularity"])
	return new("publications", fields)
}

func NewPublicationAuthor(fields map[string]string) (*Tuple, error) {
	if _, ok := fields["publication_slug"]; !ok {
		fields["publication_slug"] = PublicationSlug()
	}
	if _, ok := fields["author_login"]; !ok {
		fields["author_login"] = Login()
	}
	return new("publication_authors", fields)
}

func Insert(t *Tuple) (err error) {
	header, tuple := tupleHeaderAndValues(FieldsFor(t.TableName), t.Fields)
	_, err = DB.Exec(fmt.Sprintf("INSERT INTO %s %s VALUES %s;", t.TableName, header, tuple))
	return err
}

func InsertMany(tuples ...*Tuple) (err error) {
	return InsertList(tuples)
}

func InsertList(tuples []*Tuple) (err error) {
	var (
		tableName  string
		header     string
		values     string
		manyValues = []string{}
	)
	for _, t := range tuples {
		tableName = t.TableName
		header, values = tupleHeaderAndValues(FieldsFor(t.TableName), t.Fields)
		manyValues = append(manyValues, values)
	}
	values = strings.Join(manyValues, ", ")
	q := fmt.Sprintf("INSERT INTO %s %s VALUES %s;", tableName, header, values)
	_, err = DB.Exec(q)
	return err
}

func Truncate(tableNames ...string) (err error) {
	for _, tn := range tableNames {
		q := fmt.Sprintf("TRUNCATE TABLE %s CONTINUE IDENTITY CASCADE;", tn)
		_, err = DB.Exec(q)
		if err != nil {
			return
		}
	}
	return
}

func new(tableName string, fields map[string]string) (*Tuple, error) {
	fieldNames := FieldsFor(tableName)
	result := map[string]string{}
	for fname, fval := range fields {
		if has(fieldNames, fname) {
			result[fname] = fval
		} else {
			return &Tuple{
				TableName: tableName,
				Fields:    result,
			}, fmt.Errorf("wrong field name: `%s` for `%s` table\n\texpected any of: `%s`", fname, tableName, strings.Join(fieldNames, ", "))
		}
	}
	return &Tuple{
		TableName: tableName,
		Fields:    result,
	}, nil
}

func headerAndLiteral(values []string) string {
	var (
		sb strings.Builder
		ln = len(values)
	)
	sb.WriteRune('(')
	for _, value := range values[:ln-1] {
		sb.WriteString(fmt.Sprintf("'%s',", value))
	}
	sb.WriteString(fmt.Sprintf("'%s'", values[ln-1]))
	sb.WriteRune(')')
	return sb.String()
}

func tupleHeaderAndValues(fieldNames []string, fields map[string]string) (string, string) {
	var (
		l      = len(fields)
		names  = make([]string, l)
		values = make([]string, l)
	)
	for i, name := range fieldNames {
		names[i] = name
		values[i] = fmt.Sprintf("'%s'", strings.ReplaceAll(fields[name], "'", string('\'')))
	}

	return fmt.Sprintf("(%s)", strings.Join(names, ", ")),
		fmt.Sprintf("(%s)", strings.Join(values, ", "))
}

func Must(t *Tuple, err error) *Tuple {
	if err != nil {
		panic(err)
	}
	return t
}

func Seed() {
	Truncate("accounts", "rubrics", "publications", "publication_authors")
	err := InsertMany(Must(NewAccount(map[string]string{})), Must(NewAccount(map[string]string{})), Must(NewAccount(map[string]string{})))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	err = InsertMany(Must(NewRubric(map[string]string{})), Must(NewRubric(map[string]string{})), Must(NewRubric(map[string]string{})), Must(NewRubric(map[string]string{})))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	err = InsertMany(Must(NewPublication(map[string]string{})), Must(NewPublication(map[string]string{})), Must(NewPublication(map[string]string{})), Must(NewPublication(map[string]string{})))
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
	}
	acc := []*Tuple{}
	fields := map[string]string{}
	for _, s := range publicationSlugs {

		fields["publication_slug"] = s
		fields["author_login"] = logins[Random.Intn(len(logins))]
		acc = append(acc, Must(NewPublicationAuthor(fields)))

	}
	InsertList(acc)
}
