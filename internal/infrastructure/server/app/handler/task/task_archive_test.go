package task

import (
	"context"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type myQuerierWithTaskArchiveMockObject struct {
	db.Querier
	mock.Mock
}

func (m *myQuerierWithTaskArchiveMockObject) TaskArchive(ctx context.Context, arg db.TaskArchiveParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func TestHandler_archiveTask(t *testing.T) {
	var id pgtype.UUID
	id.Scan("7142c1a1-30d4-452c-af3e-47fb821e4646")
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/project/archive",
		struct{ io.Reader }{strings.NewReader("{\"id\":\"" + id.String() + "\", \"archived\":true}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(200)
	ctx := context.Background()
	testObj := new(myQuerierWithTaskArchiveMockObject)

	testObj.On(
		"TaskArchive",
		ctx,
		mock.AnythingOfType("db.TaskArchiveParams"),
	).Return(nil)

	h := &ArchiveTaskHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.NoError(t, err)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "\""+id.String()+"\"\n", string(body))
}

func TestHandler_archiveProject_ValidationError(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/project/create",
		struct{ io.Reader }{strings.NewReader("{}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(400)
	ctx := context.Background()
	testObj := new(myQuerierWithTaskArchiveMockObject)

	testObj.AssertNotCalled(t, "TaskArchive", ctx, mock.Anything)

	h := &ArchiveTaskHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.Error(t, err)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "Key: 'ArchiveTaskRequestParams.Archived' "+
		"Error:Field validation for 'Archived' failed on the 'required' tag", err.Error())
}
