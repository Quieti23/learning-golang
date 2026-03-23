package repository

import (
	"database/sql"
	"errors"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"month02blogapi/model"
)

func newMySQLRepositoryForTest(t *testing.T) (*MySQLPostRepository, sqlmock.Sqlmock) {
	t.Helper()

	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("create sqlmock: %v", err)
	}

	return NewMySQLPostRepository(db), mock
}

func TestMySQLPostRepositoryList(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	createdAt1 := time.Date(2026, 3, 23, 10, 0, 0, 0, time.UTC)
	updatedAt1 := createdAt1.Add(5 * time.Minute)
	createdAt2 := createdAt1.Add(10 * time.Minute)
	updatedAt2 := createdAt2.Add(5 * time.Minute)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author", "created_at", "updated_at"}).
		AddRow(2, "second", "content 2", "eson", createdAt2, updatedAt2).
		AddRow(1, "first", "content 1", "eson", createdAt1, updatedAt1)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		ORDER BY id DESC
	`)).WillReturnRows(rows)

	posts, err := repo.List()
	if err != nil {
		t.Fatalf("List() error = %v", err)
	}

	if len(posts) != 2 {
		t.Fatalf("List() len = %d, want 2", len(posts))
	}

	if posts[0].ID != 2 || posts[1].ID != 1 {
		t.Fatalf("List() order = [%d, %d], want [2, 1]", posts[0].ID, posts[1].ID)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryListReturnsQueryError(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		ORDER BY id DESC
	`)).WillReturnError(errors.New("query failed"))

	_, err := repo.List()
	if err == nil || err.Error() != "query failed" {
		t.Fatalf("List() error = %v, want query failed", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryListReturnsScanError(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author", "created_at", "updated_at"}).
		AddRow("bad-id", "first", "content", "eson", time.Now(), time.Now())

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		ORDER BY id DESC
	`)).WillReturnRows(rows)

	_, err := repo.List()
	if err == nil {
		t.Fatal("List() error = nil, want scan error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryCreate(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	input := model.Post{
		Title:   "first post",
		Content: "hello mysql",
		Author:  "eson",
	}

	createdAt := time.Date(2026, 3, 23, 11, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(2 * time.Minute)

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO posts (title, content, author, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`)).
		WithArgs(input.Title, input.Content, input.Author).
		WillReturnResult(sqlmock.NewResult(3, 1))

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author", "created_at", "updated_at"}).
		AddRow(3, input.Title, input.Content, input.Author, createdAt, updatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		WHERE id = ?
	`)).
		WithArgs(int64(3)).
		WillReturnRows(rows)

	created, err := repo.Create(input)
	if err != nil {
		t.Fatalf("Create() error = %v", err)
	}

	if created.ID != 3 {
		t.Fatalf("Create() id = %d, want 3", created.ID)
	}

	if created.Title != input.Title || created.Content != input.Content || created.Author != input.Author {
		t.Fatalf("Create() returned unexpected post = %+v", created)
	}

	if created.CreatedAt != createdAt || created.UpdatedAt != updatedAt {
		t.Fatalf("Create() timestamps = %v / %v, want %v / %v", created.CreatedAt, created.UpdatedAt, createdAt, updatedAt)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryCreateReturnsExecError(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	input := model.Post{Title: "first post", Content: "hello mysql", Author: "eson"}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO posts (title, content, author, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`)).
		WithArgs(input.Title, input.Content, input.Author).
		WillReturnError(errors.New("insert failed"))

	_, err := repo.Create(input)
	if err == nil || err.Error() != "insert failed" {
		t.Fatalf("Create() error = %v, want insert failed", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryCreateReturnsQueryError(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	input := model.Post{Title: "first post", Content: "hello mysql", Author: "eson"}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO posts (title, content, author, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`)).
		WithArgs(input.Title, input.Content, input.Author).
		WillReturnResult(sqlmock.NewResult(3, 1))

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		WHERE id = ?
	`)).
		WithArgs(int64(3)).
		WillReturnError(errors.New("select failed"))

	_, err := repo.Create(input)
	if err == nil || err.Error() != "select failed" {
		t.Fatalf("Create() error = %v, want select failed", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryCreateReturnsLastInsertIDError(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	input := model.Post{Title: "first post", Content: "hello mysql", Author: "eson"}

	mock.ExpectExec(regexp.QuoteMeta(`
		INSERT INTO posts (title, content, author, created_at, updated_at)
		VALUES (?, ?, ?, NOW(), NOW())
	`)).
		WithArgs(input.Title, input.Content, input.Author).
		WillReturnResult(sqlmock.NewErrorResult(errors.New("last insert id failed")))

	_, err := repo.Create(input)
	if err == nil || err.Error() != "last insert id failed" {
		t.Fatalf("Create() error = %v, want last insert id failed", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryGetByID(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	createdAt := time.Date(2026, 3, 23, 12, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(5 * time.Minute)

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author", "created_at", "updated_at"}).
		AddRow(8, "title", "content", "eson", createdAt, updatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		WHERE id = ?
	`)).
		WithArgs(8).
		WillReturnRows(rows)

	post, err := repo.GetByID(8)
	if err != nil {
		t.Fatalf("GetByID() error = %v", err)
	}
	if post.ID != 8 || post.Title != "title" || post.Content != "content" || post.Author != "eson" {
		t.Fatalf("GetByID() returned unexpected post = %+v", post)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryGetByIDReturnsNotFound(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		WHERE id = ?
	`)).
		WithArgs(99).
		WillReturnError(sql.ErrNoRows)

	_, err := repo.GetByID(99)
	if !errors.Is(err, ErrPostNotFound) {
		t.Fatalf("GetByID() error = %v, want %v", err, ErrPostNotFound)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryGetByIDReturnsOtherError(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		WHERE id = ?
	`)).
		WithArgs(7).
		WillReturnError(errors.New("db unavailable"))

	_, err := repo.GetByID(7)
	if err == nil || err.Error() != "db unavailable" {
		t.Fatalf("GetByID() error = %v, want db unavailable", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryUpdate(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	input := model.Post{
		ID:      5,
		Title:   "updated",
		Content: "updated content",
		Author:  "eson",
	}

	createdAt := time.Date(2026, 3, 23, 9, 0, 0, 0, time.UTC)
	updatedAt := createdAt.Add(30 * time.Minute)

	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE posts
		SET title = ?, content = ?, author = ?, updated_at = NOW()
		WHERE id = ?
	`)).
		WithArgs(input.Title, input.Content, input.Author, input.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	rows := sqlmock.NewRows([]string{"id", "title", "content", "author", "created_at", "updated_at"}).
		AddRow(input.ID, input.Title, input.Content, input.Author, createdAt, updatedAt)

	mock.ExpectQuery(regexp.QuoteMeta(`
		SELECT id, title, content, author, created_at, updated_at
		FROM posts
		WHERE id = ?
	`)).
		WithArgs(input.ID).
		WillReturnRows(rows)

	updated, err := repo.Update(input)
	if err != nil {
		t.Fatalf("Update() error = %v", err)
	}

	if updated.ID != input.ID || updated.Title != input.Title || updated.Content != input.Content || updated.Author != input.Author {
		t.Fatalf("Update() returned unexpected post = %+v", updated)
	}

	if updated.CreatedAt != createdAt || updated.UpdatedAt != updatedAt {
		t.Fatalf("Update() timestamps = %v / %v, want %v / %v", updated.CreatedAt, updated.UpdatedAt, createdAt, updatedAt)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryUpdateReturnsNotFound(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	post := model.Post{ID: 77, Title: "missing", Content: "missing", Author: "eson"}

	mock.ExpectExec(regexp.QuoteMeta(`
		UPDATE posts
		SET title = ?, content = ?, author = ?, updated_at = NOW()
		WHERE id = ?
	`)).
		WithArgs(post.Title, post.Content, post.Author, post.ID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	_, err := repo.Update(post)
	if !errors.Is(err, ErrPostNotFound) {
		t.Fatalf("Update() error = %v, want %v", err, ErrPostNotFound)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryDelete(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM posts WHERE id = ?`)).
		WithArgs(5).
		WillReturnResult(sqlmock.NewResult(0, 1))

	if err := repo.Delete(5); err != nil {
		t.Fatalf("Delete() error = %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryDeleteReturnsNotFound(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM posts WHERE id = ?`)).
		WithArgs(404).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.Delete(404)
	if !errors.Is(err, ErrPostNotFound) {
		t.Fatalf("Delete() error = %v, want %v", err, ErrPostNotFound)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}

func TestMySQLPostRepositoryDeleteReturnsExecError(t *testing.T) {
	repo, mock := newMySQLRepositoryForTest(t)

	mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM posts WHERE id = ?`)).
		WithArgs(12).
		WillReturnError(errors.New("delete failed"))

	err := repo.Delete(12)
	if err == nil || err.Error() != "delete failed" {
		t.Fatalf("Delete() error = %v, want delete failed", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet sql expectations: %v", err)
	}
}
