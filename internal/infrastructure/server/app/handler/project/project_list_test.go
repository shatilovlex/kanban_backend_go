package project

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

type myQuerierWithProjectListMockObject struct {
	db.Querier
	mock.Mock
}

func (m *myQuerierWithProjectListMockObject) ProjectList(ctx context.Context) ([]*db.ProjectListRow, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*db.ProjectListRow), args.Error(1)
}

func TestHandler_getProjects(t *testing.T) {
	var projectID pgtype.UUID
	name := "name"
	description := "desc"
	projectID.Scan("7142c1a1-30d4-452c-af3e-47fb821e4646")
	expectedProjectList := make([]*db.ProjectListRow, 1)
	expectedProjectList[0] = &db.ProjectListRow{
		ID:          projectID,
		Name:        &name,
		Description: &description,
	}
	req := httptest.NewRequest(
		"GET",
		"http://example.com/v1/projects",
		struct{ io.Reader }{strings.NewReader("")},
	)
	w := httptest.NewRecorder()
	ctx := context.Background()
	testObj := new(myQuerierWithProjectListMockObject)
	testObj.On("ProjectList", ctx).Return(expectedProjectList, nil)

	h := &GetProjectListHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.NoError(t, err)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "[{"+
		"\"id\":\"7142c1a1-30d4-452c-af3e-47fb821e4646\","+
		"\"name\":\""+name+"\","+
		"\"description\":\""+description+"\"}]\n",
		string(body),
	)
}
