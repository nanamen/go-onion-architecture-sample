package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/nanamen/go-echo-rest-sample/domain/model"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

// TODO:mock生成を共通化する
type mockUserUseCase struct{}

func (u *mockUserUseCase) GetUsers(ctx context.Context) ([]*model.User, error) {
	return getMockUsers(5), nil
}

func (u *mockUserUseCase) GetUser(ctx context.Context, id int) (*model.User, error) {
	return getMockUser(id), nil
}

func (u *mockUserUseCase) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return getMockUser(1), nil
}

func (u *mockUserUseCase) UpdateUser(ctx context.Context, id int) (*model.User, error) {
	mu := getMockUser(id)
	mu.Name = mu.Name + "_updated"
	return mu, nil
}

func (u *mockUserUseCase) DeleteUser(ctx context.Context, id int) error {
	return nil
}

func getMockUsers(n int) []*model.User {
	ret := []*model.User{}
	for i := 0; i < n; i++ {
		u := getMockUser(int(i))
		ret = append(ret, u)
	}
	return ret
}

func getMockUser(id int) *model.User {
	u := &model.User{
		ID:        id,
		Name:      fmt.Sprintf("name_%d", id),
		CreatedAt: time.Date(2015, 9, 13, 12, 35, 42, 123456789, time.Local),
		UpdatedAt: time.Date(2015, 9, 13, 12, 35, 42, 123456789, time.Local),
	}
	return u
}

func getMockUserNoID() *model.User {
	u := &model.User{
		Name:      fmt.Sprintf("name_%d", 1),
		CreatedAt: time.Date(2015, 9, 13, 12, 35, 42, 123456789, time.Local),
		UpdatedAt: time.Date(2015, 9, 13, 12, 35, 42, 123456789, time.Local),
	}
	return u
}

func TestUserHandler_GetUsers(t *testing.T) {
	// expected
	expected := getMockUsers(5)

	// set stub
	usecase := &mockUserUseCase{}
	h := NewUserHandler(usecase)

	e := echo.New()
	req := httptest.NewRequest(echo.GET, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, h.GetUsers(c)) {
		users := []*model.User{}
		if err := json.Unmarshal(rec.Body.Bytes(), &users); err != nil {
			t.Fatal(err)
		}

		t.Log(rec.Code)
		assert.Equal(t, http.StatusOK, rec.Code)

		for i, u := range users {
			t.Log(u)
			assert.Equal(t, expected[i], u)
		}
	}

}

// Table Driven Tests
type getUserTest struct {
	ID   int
	User *model.User
}

var getUserTests = []getUserTest{
	{math.MaxInt8, getMockUser(math.MaxInt8)},
	{math.MaxInt16, getMockUser(math.MaxInt16)},
	{math.MaxInt32, getMockUser(math.MaxInt32)},
	{math.MaxInt64, getMockUser(math.MaxInt64)},
}

func TestUserHandler_GetUser(t *testing.T) {
	// set stub
	usecase := &mockUserUseCase{}
	h := NewUserHandler(usecase)

	for _, test := range getUserTests {
		// set request
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprint(test.ID))

		// assertions
		if assert.NoError(t, h.GetUser(c)) {
			user := &model.User{}
			if err := json.Unmarshal(rec.Body.Bytes(), &user); err != nil {
				t.Fatal(err)
			}
			t.Log(rec.Code)
			assert.Equal(t, http.StatusOK, rec.Code)
			t.Log(user)
			assert.Equal(t, test.User, user)
		}
	}
}

func TestUserHandler_CreateUser(t *testing.T) {
	// expected
	expected := getMockUser(1)

	// set stub
	usecase := &mockUserUseCase{}
	h := NewUserHandler(usecase)

	e := echo.New()
	jsonBytes, err := json.Marshal(getMockUserNoID())
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(echo.POST, "/users", strings.NewReader(string(jsonBytes)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	// Assertions
	if assert.NoError(t, h.CreateUser(c)) {
		user := &model.User{}
		if err := json.Unmarshal(rec.Body.Bytes(), &user); err != nil {
			t.Fatal(err)
		}

		t.Log(rec.Code)
		assert.Equal(t, http.StatusCreated, rec.Code)
		t.Log(user)
		assert.Equal(t, expected, user)
	}
}

type updateUserTest struct {
	ID       int
	UserName string
}

var updateUserTests = []updateUserTest{
	{math.MaxInt8, fmt.Sprintf("name_%d_updated", math.MaxInt8)},
	{math.MaxInt16, fmt.Sprintf("name_%d_updated", math.MaxInt16)},
	{math.MaxInt32, fmt.Sprintf("name_%d_updated", math.MaxInt32)},
	{math.MaxInt64, fmt.Sprintf("name_%d_updated", math.MaxInt64)},
}

func TestUserHandler_UpdateUser(t *testing.T) {
	// set stub
	usecase := &mockUserUseCase{}
	h := NewUserHandler(usecase)

	for _, test := range updateUserTests {
		// set request
		e := echo.New()
		req := httptest.NewRequest(echo.PUT, "/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/users/:id")
		c.SetParamNames("id")
		c.SetParamValues(fmt.Sprint(test.ID))

		// assertions
		if assert.NoError(t, h.UpdateUser(c)) {
			user := &model.User{}
			if err := json.Unmarshal(rec.Body.Bytes(), &user); err != nil {
				t.Fatal(err)
			}
			t.Log(rec.Code)
			assert.Equal(t, http.StatusOK, rec.Code)
			t.Log(user)
			assert.Equal(t, test.UserName, user.Name)
		}
	}
}

func TestUserHandler_DeleteUser(t *testing.T) {
	// set stub
	usecase := &mockUserUseCase{}
	h := NewUserHandler(usecase)

	// set request
	e := echo.New()
	req := httptest.NewRequest(echo.DELETE, "/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues(fmt.Sprint(1))

	// assertions
	if assert.NoError(t, h.DeleteUser(c)) {
		t.Log(rec.Code)
		assert.Equal(t, http.StatusNoContent, rec.Code)
	}
}
