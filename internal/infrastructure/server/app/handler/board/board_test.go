package board

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

type myQuerierMockObject struct {
	db.Querier
	mock.Mock
}

func (m *myQuerierMockObject) Board(ctx context.Context, projectID pgtype.UUID) ([]*db.BoardRow, error) {
	args := m.Called(ctx, projectID)
	return args.Get(0).([]*db.BoardRow), args.Error(1)
}

func TestHandler_getBoard(t *testing.T) {
	var projectID pgtype.UUID
	projectID.Scan("7142c1a1-30d4-452c-af3e-47fb821e4646")
	req := httptest.NewRequest(
		"GET",
		"http://example.com/v1/board?project_id="+projectID.String(),
		struct{ io.Reader }{strings.NewReader("")},
	)
	w := httptest.NewRecorder()
	ctx := context.Background()
	testObj := new(myQuerierMockObject)

	testObj.On("Board", ctx, projectID).Return(make([]*db.BoardRow, 0), nil)

	h := &GetBoardHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.NoError(t, err)

	resp := w.Result()
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	assert.Equal(t, 200, resp.StatusCode)
	assert.Equal(t, "[]\n", string(body))
}

func TestHandler_getBoard_ValidationError(t *testing.T) {
	req := httptest.NewRequest(
		"GET",
		"http://example.com/v1/board?project_id=4",
		struct{ io.Reader }{strings.NewReader("")},
	)
	w := httptest.NewRecorder()
	w.WriteHeader(400)
	ctx := context.Background()
	testObj := new(myQuerierMockObject)

	testObj.AssertNotCalled(t, "Board", ctx, mock.Anything)

	h := &GetBoardHandler{handler.NewHandlerMock(ctx, testObj)}
	err := h.Handle(w, req)

	assert.Error(t, err)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, 400, resp.StatusCode)
	assert.Equal(t, "cannot parse UUID 4", err.Error())
}
