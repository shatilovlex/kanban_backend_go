package list

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

type myQuerierWithListRemoveMockObject struct {
	db.Querier
	mock.Mock
}

func (m *myQuerierWithListRemoveMockObject) ListRemove(ctx context.Context, id pgtype.UUID) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestHandler_removeList(t *testing.T) {
	var projectID pgtype.UUID
	projectID.Scan("7142c1a1-30d4-452c-af3e-47fb821e4646")
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/removeList",
		struct{ io.Reader }{strings.NewReader("{\"id\":\"" + projectID.String() + "\"}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(200)
	ctx := context.Background()
	testObj := new(myQuerierWithListRemoveMockObject)

	testObj.On("ListRemove", ctx, projectID).Return(nil)

	h := &RemoveListHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.NoError(t, err)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "\""+projectID.String()+"\"\n", string(body))
}

func TestHandler_removeList_ValidationError(t *testing.T) {
	req := httptest.NewRequest(
		"POST",
		"http://example.com/v1/removeList",
		struct{ io.Reader }{strings.NewReader("{\"id\":\"4\"}")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(400)
	ctx := context.Background()
	testObj := new(myQuerierWithListRemoveMockObject)

	testObj.AssertNotCalled(t, "ListRemove", ctx, mock.Anything)

	h := &RemoveListHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.Error(t, err)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "invalid length for UUID: 3", err.Error())
}
