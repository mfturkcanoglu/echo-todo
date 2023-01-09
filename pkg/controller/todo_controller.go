package controller

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/mfturkcanoglu/echo-todo/pkg/entity"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

type todoController struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewTodoController(db *sql.DB, l *zap.Logger) *todoController {
	return &todoController{
		db:     db,
		logger: l,
	}
}

// api/todo/:id
func (t *todoController) Get(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.RestResponse{
			Success: false,
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	todo, err := todoById(id, t.db, t.logger)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.RestResponse{
			Success: false,
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	return c.JSON(http.StatusOK, entity.RestResponse{
		Success: true,
		Data:    todo,
		Code:    http.StatusOK,
	})
}

func (t *todoController) GetAll(c echo.Context) error {
	fmt.Println(t.db)
	todos, err := todos(t.db, t.logger)
	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.RestResponse{
			Code:    http.StatusBadRequest,
			Success: true,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, entity.RestResponse{
		Code:    http.StatusOK,
		Data:    todos,
		Success: true,
	})
}

func (t *todoController) Post(c echo.Context) error {
	var todo entity.Todo
	if err := c.Bind(&todo); err != nil {
		return c.JSON(http.StatusBadRequest, entity.RestResponse{
			Success: false,
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	todo.CreatedAt = time.Now()
	todo.Deleted = false

	id, err := createTodo(todo, t.db, t.logger)

	if err != nil {
		return c.JSON(http.StatusBadRequest, entity.RestResponse{
			Success: false,
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		})
	}

	todo.Id = id

	return c.JSON(http.StatusCreated, entity.RestResponse{
		Success: true,
		Message: "created",
		Code:    http.StatusCreated,
		Data:    todo,
	})
}

func todoById(id int, db *sql.DB, l *zap.Logger) (todo entity.Todo, err error) {
	row := db.QueryRow("SELECT * FROM todo WHERE id = $1", id)
	if rowErr := row.Scan(&todo.Id, &todo.Text, &todo.CreatedAt, &todo.Deleted); rowErr != nil {
		if rowErr == sql.ErrNoRows {
			err = fmt.Errorf("todoById %d: no such todo", id)
			return
		}
		err = fmt.Errorf("todoById %d: %v", id, err)
		return
	}
	return
}

func todos(db *sql.DB, l *zap.Logger) (todos []entity.Todo, err error) {
	rows, err := db.Query("SELECT * FROM todo;")

	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var todo entity.Todo
		if rowErr := rows.Scan(&todo.Id, &todo.Text, &todo.CreatedAt, &todo.Deleted); rowErr != nil {
			return
		}
		todos = append(todos, todo)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		err = rowsErr
		return
	}
	return
}

func createTodo(todo entity.Todo, db *sql.DB, l *zap.Logger) (id int64, err error) {
	err = db.QueryRow(`INSERT INTO todo (text, created_at, deleted) VALUES ($1, $2, $3) RETURNING id`,
		todo.Text, todo.CreatedAt, todo.Deleted).Scan(&id)

	if err != nil {
		return
	}
	return
}
