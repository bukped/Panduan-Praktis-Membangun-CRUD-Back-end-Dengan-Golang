package handlers

import (
	"database/sql"
	"example/crud/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type H map[string]interface{}

func GetComputers(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, models.GetComputers(db))
	}
}

func PutComputer(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var computer models.Computer

		c.Bind(&computer)

		id, err := models.PutComputer(db, computer.Name, computer.Price)

		if err == nil {
			return c.JSON(http.StatusCreated, H{
				"created": id,
			})
		} else {
			return err
		}

	}
}

func EditComputer(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {

		var computer models.Computer
		c.Bind(&computer)

		_, err := models.EditComputer(db, computer.ID, computer.Name, computer.Price)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"updated": computer,
			})
		} else {
			return err
		}
	}
}

func DeleteComputer(db *sql.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		_, err := models.DeleteComputer(db, id)

		if err == nil {
			return c.JSON(http.StatusOK, H{
				"deleted": id,
			})
		} else {
			return err
		}

	}
}
