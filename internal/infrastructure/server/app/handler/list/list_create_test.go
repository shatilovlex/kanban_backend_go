package list

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

type myQuerierWithListAddMockObject struct {
	db.Querier
	mock.Mock
}

func (m *myQuerierWithListAddMockObject) ListAdd(ctx context.Context, arg db.ListAddParams) error {
	args := m.Called(ctx, arg)
	return args.Error(0)
}

func TestHandler_createList(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/addList",
		struct{ io.Reader }{strings.NewReader("{\"projectId\":\"7142c1a1-30d4-452c-af3e-47fb821e4646\"," +
			"\"name\":\"New List\",\"sort\":4}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(201)
	ctx := context.Background()
	testObj := new(myQuerierWithListAddMockObject)

	testObj.On("ListAdd", ctx, mock.AnythingOfType("db.ListAddParams")).Return(nil)

	h := &CreateListHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.NoError(t, err)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 201, resp.StatusCode)
	assert.Equal(t, "", string(body))
}

func TestHandler_createList_ValidationError(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/addList",
		struct{ io.Reader }{strings.NewReader("{}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(400)
	ctx := context.Background()
	testObj := new(myQuerierWithListAddMockObject)

	testObj.On("ListAdd", ctx, mock.AnythingOfType("db.ListAddParams")).Return(nil)

	h := &CreateListHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.Error(t, err)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "Key: 'CreateListRequestParams.Name' "+
		"Error:Field validation for 'Name' failed on the 'required' tag\n"+
		"Key: 'CreateListRequestParams.Sort' "+
		"Error:Field validation for 'Sort' failed on the 'required' tag", err.Error())
}
