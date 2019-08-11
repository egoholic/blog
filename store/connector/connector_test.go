package connector_test

import (
	"database/sql"

	"github.com/egoholic/blog/config"
	. "github.com/egoholic/blog/store/connector"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("connector", func() {
	Describe("New()", func() {
		It("returns connector", func() {
			credsKeeper := &config.DBCredentials{Host: "localhost", Port: 5432, User: "postgres", Password: "", DBName: ""}
			conn := New(credsKeeper)
			Expect(conn).NotTo(BeNil())
			Expect(conn).To(BeAssignableToTypeOf(&Connector{}))
		})
	})

	Describe("Connector", func() {
		Describe(".Connection()", func() {
			It("returns connection", func() {
				credsKeeper := &config.DBCredentials{Host: "localhost", Port: 5432, User: "postgres", Password: "", DBName: ""}
				conn := New(credsKeeper)
				db := conn.Connection()
				Expect(db).NotTo(BeNil())
				Expect(db).To(BeAssignableToTypeOf(&sql.DB{}))
			})
		})
	})
})
