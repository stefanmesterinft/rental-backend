// package routes

// import (
// 	"net/http"
// 	"strconv"

// 	"github.com/gin-gonic/gin"
// 	"rental.com/api/models"
// )

// func getEvents(context *gin.Context) {
// 	events, err := models.GetAllEvents()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch events. Try again later"})
// 	}
// 	context.JSON(http.StatusOK, events)
// }

// func getEvent(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
// 		return
// 	}

// 	event, err := models.GetEventById(eventId)

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch event"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, event)
// }

// func createEvent(context *gin.Context) {
// 	var event models.Event
// 	err := context.ShouldBindJSON(&event)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the data"})
// 		return
// 	}

// 	userId := context.GetInt64("userId")
// 	event.UserID = userId

// 	err = event.Save()

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not save event. Try again later"})
// 		return
// 	}

// 	context.JSON(http.StatusCreated, gin.H{"message": "Event created!", "event": event})
// }

// func updateEvent(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
// 		return
// 	}

// 	userId := context.GetInt64("userId")
// 	event, err := models.GetEventById(eventId)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event. Try again later"})
// 		return
// 	}

// 	if event.UserID != userId {
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to update event"})
// 		return
// 	}

// 	var updateEvent models.Event
// 	err = context.ShouldBindJSON(&updateEvent)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data. Try again later"})
// 		return
// 	}

// 	updateEvent.ID = eventId
// 	err = updateEvent.Update()
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch the data. Try again later"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "Event updated succesfully"})
// }

// func deleteEvent(context *gin.Context) {
// 	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse event id"})
// 		return
// 	}

// 	userId := context.GetInt64("userId")
// 	event, err := models.GetEventById(eventId)

// 	if err != nil {
// 		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not fetch the event. Try again later"})
// 		return
// 	}

// 	if event.UserID != userId {
// 		context.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized to delete event"})
// 		return
// 	}

// 	err = event.Delete()

// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not delete the event. Try again later"})
// 		return
// 	}

// 	context.JSON(http.StatusOK, gin.H{"message": "Event deleted succesfully"})
// }
