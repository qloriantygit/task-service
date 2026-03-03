package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Message struct {
	ID int `json:"id"`
	Text string `json:"text"`
}

var messages =make(map[int]Message)
var nextID = 1
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	e := echo.New()

	e.GET("/messages", GetHandlerr)
	e.POST("/messages", PostHandler)
	e.PATCH("messages/:id", PatchHandler)
	e.DELETE("messages/:id", DeleteHandler)

	e.Start(":8080")
}

func GetHandlerr(c echo.Context) error {
	var msgSlice []Message

	for _, msg := range messages {
		msgSlice = append(msgSlice, msg )
	}

	return c.JSON(http.StatusOK, &msgSlice)
}
func PostHandler(c echo.Context) error {
	var message Message
	if err := c.Bind(&message); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "could not add the message",
		})
	}

	message.ID = nextID
	nextID++


	messages[message.ID] = message
	return c.JSON(http.StatusOK, Response{
		Status:  "succes",
		Message: "Message was succesfully added",
	})

}

func PatchHandler(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Message: "Bad ID",
		})
	}
	var updatedmessage Message
	if err := c.Bind(&updatedmessage); err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "could not update the message",
		})
	}

	if _, exist := messages[id]; !exist {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Message was not found",
		})
	}
	updatedmessage.ID = id
	messages[id] = updatedmessage


	return c.JSON(http.StatusOK, Response{
		Status: "sucess",
		Message: "message was updated",
	})
}

func DeleteHandler(c echo.Context) error {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, Response{
			Status: "Error",
			Message: "Bad ID",
		})
	}

	if _, exist := messages[id]; !exist {
		return c.JSON(http.StatusBadRequest, Response{
			Status:  "error",
			Message: "Message was not found",
		})
	}

	delete(messages, id)

	return c.JSON(http.StatusOK, Response{
		Status: "sucess",
		Message: "message was deleted",
	})

}