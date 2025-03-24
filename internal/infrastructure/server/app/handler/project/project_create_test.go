package project

import (
	"context"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/db"
	"github.com/shatilovlex/kanban_backend_go/internal/infrastructure/server/app/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type myQuerierWithProjectCreateMockObject struct {
	db.Querier
	mock.Mock
}

func (m *myQuerierWithProjectCreateMockObject) ProjectCreate(ctx context.Context, arg db.ProjectCreateParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func TestHandler_createProject(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/project/create",
		struct{ io.Reader }{strings.NewReader("{\"name\":\"Website Development\"," +
			"\"description\":\"website for multiple clients\"}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(201)
	ctx := context.Background()
	testObj := new(myQuerierWithProjectCreateMockObject)

	testObj.On(
		"ProjectCreate",
		ctx,
		mock.AnythingOfType("db.ProjectCreateParams"),
	).Return(nil)

	h := &CreateProjectHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.NoError(t, err)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 201, resp.StatusCode)
	assert.True(t, strings.Contains(string(body), "\"name\":\"Website Development\""))
	assert.True(t, strings.Contains(string(body), "\"description\":\"website for multiple clients\""))
}

func TestHandler_createProject_ValidationError(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/project/create",
		struct{ io.Reader }{strings.NewReader("{}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(400)
	ctx := context.Background()
	testObj := new(myQuerierWithProjectCreateMockObject)
	testObj.AssertNotCalled(t, "ProjectCreate", ctx, mock.Anything)

	h := &CreateProjectHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.Error(t, err)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "Key: 'CreateProjectRequestParams.Name' "+
		"Error:Field validation for 'Name' failed on the 'required' tag\n"+
		"Key: 'CreateProjectRequestParams.Description' "+
		"Error:Field validation for 'Description' failed on the 'required' tag", err.Error())
}
