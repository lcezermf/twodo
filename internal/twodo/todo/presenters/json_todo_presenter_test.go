package presenters

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/lccezinha/twodo/internal/twodo"
)

func TestPresentTodo(t *testing.T) {
	w := httptest.NewRecorder()
	presenter := JSONTodoPresenter{w}
	todo := twodo.Todo{
		ID:          1,
		Title:       "Title",
		Description: "Description",
		// CreatedAt:   time.Now(),
		Done: false,
	}

	presenter.Present(http.StatusOK, todo)

	expectedBody := []byte(
		fmt.Sprintf(
			`{"id":%d,"title":"%s","description":"%s","done":%t}`, todo.ID, todo.Title, todo.Description, todo.Done,
		),
	)

	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)
	expectedStatus := http.StatusOK

	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("Expected: %s. Actual: %s", expectedBody, body)
	}

	if w.Result().StatusCode != expectedStatus {
		t.Errorf("Expected: %d. Actual: %d", expectedStatus, w.Result().StatusCode)
	}

	expectedContentType := "application/json"
	contentType := w.Result().Header.Get("Content-Type")

	if contentType != expectedContentType {
		t.Errorf("Expected: %s. Actual: %s", expectedContentType, contentType)
	}
}

func TestPresentTodoErrs(t *testing.T) {
	w := httptest.NewRecorder()
	presenter := JSONTodoPresenter{w}
	errs := []twodo.ValidationError{
		twodo.ValidationError{
			Field:   "Title",
			Message: "Can not be blank",
			Type:    "Required",
		},
	}

	presenter.PresentErrors(errs)

	expectedError := fmt.Sprintf(`{"field":"%s","message":"%s","type":"%s"}`, errs[0].Field, errs[0].Message, errs[0].Type)
	expectedBody := []byte(`{"errors":[` + expectedError + `]}`)

	response := w.Result()
	body, _ := ioutil.ReadAll(response.Body)
	expectedStatus := http.StatusBadRequest

	if !reflect.DeepEqual(body, expectedBody) {
		t.Errorf("Expected: %s. Actual: %s", expectedBody, body)
	}

	if w.Result().StatusCode != expectedStatus {
		t.Errorf("Expected: %d. Actual: %d", expectedStatus, w.Result().StatusCode)
	}

	expectedContentType := "application/json"
	contentType := w.Result().Header.Get("Content-Type")

	if contentType != expectedContentType {
		t.Errorf("Expected: %s. Actual: %s", expectedContentType, contentType)
	}
}
