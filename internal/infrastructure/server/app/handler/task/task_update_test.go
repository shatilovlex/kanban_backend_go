package task

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

type myQuerierWithTaskUpdateMockObject struct {
	db.Querier
	mock.Mock
}

func (m *myQuerierWithTaskUpdateMockObject) TaskUpdate(ctx context.Context, arg db.TaskUpdateParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func TestHandler_updateTask(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/tasks/create",
		struct{ io.Reader }{strings.NewReader("{" +
			"\"id\":\"c142c1a1-30d4-452c-af3e-47fb821e4646\"," +
			"\"list_id\":\"7142c1a1-30d4-452c-af3e-47fb821e4646\"," +
			"\"title\":\"task\"," +
			"\"description\":\"task\"," +
			"\"sort\":4}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(200)
	ctx := context.Background()
	testObj := new(myQuerierWithTaskUpdateMockObject)

	testObj.On(
		"TaskUpdate",
		ctx,
		mock.AnythingOfType("db.TaskUpdateParams"),
	).Return(nil)

	h := &UpdateTaskHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.NoError(t, err)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, 200, resp.StatusCode)
}

func TestHandler_updateTask_ValidationError(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/project/create",
		struct{ io.Reader }{strings.NewReader("{}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(400)
	ctx := context.Background()
	testObj := new(myQuerierWithTaskUpdateMockObject)
	testObj.AssertNotCalled(t, "TaskUpdate", ctx, mock.Anything)

	h := &CreateTaskHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.Error(t, err)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "Key: 'CreateTaskRequest.Title' "+
		"Error:Field validation for 'Title' failed on the 'required' tag\n"+
		"Key: 'CreateTaskRequest.Description' "+
		"Error:Field validation for 'Description' failed on the 'required' tag\n"+
		"Key: 'CreateTaskRequest.Sort' "+
		"Error:Field validation for 'Sort' failed on the 'required' tag", err.Error())
}
