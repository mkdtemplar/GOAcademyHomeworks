package main

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"regexp"
	"testing"

	"github.com/Rosaniline/gorm-ut/pkg/model"
	"github.com/go-test/deep"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

type Suite struct {
	suite.Suite
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repository Repo
	tStories   *topstories
}

func (s *Suite) SetupSuite() {
	var (
		db  *sql.DB
		err error
	)

	db, s.mock, err = sqlmock.New()
	require.NoError(s.T(), err)

	s.DB, err = gorm.Open(sqlite.Open("Stories.db"), &gorm.Config{})
	require.NoError(s.T(), err)

	s.repository = CreateRepository(s.DB)
}
func (s *Suite) AfterTest(_, _ string) {
	require.NoError(s.T(), s.mock.ExpectationsWereMet())
}

func TestInit(t *testing.T) {
	suite.Run(t, new(Suite))
}

func (s *Suite) Test_repository_Get() {
	var (
		id      = uuid.NewV4()
		name    = "test-name"
		stories []*topstories
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`SELECT * FROM "topstories"`)).
		WithArgs(id.String()).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).
			AddRow(id.String(), name))

	res, err := s.repository.ReadAll(stories)

	require.NoError(s.T(), err)
	require.Nil(s.T(), deep.Equal(&model.Person{ID: id, Name: name}, res))
}

func (s *Suite) Test_repository_Create() {
	var (
		id   = uuid.NewV4()
		name = "test-name"
	)

	s.mock.ExpectQuery(regexp.QuoteMeta(
		`INSERT INTO "person" ("id","name")
			VALUES ($1,$2) RETURNING "person"."id"`)).
		WithArgs(id, name).
		WillReturnRows(
			sqlmock.NewRows([]string{"id"}).AddRow(id.String()))

	err := s.repository.Create(id, name)

	require.NoError(s.T(), err)
}
