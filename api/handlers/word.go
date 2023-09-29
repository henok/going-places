package handlers

import (
    "net/http"
    "going-places-api/db"
    "going-places-api/config"
    "going-places-api/utils"
    "going-places-api/logger"
    "github.com/labstack/echo/v4"
)

func AddWord(c echo.Context) error {
    word := c.QueryParam("word")

    reasonToGo, openAiErr := utils.FetchCompletionFromOpenAI(config.ADD_WORD_PROMPT_TEMPLATE + word)
    if openAiErr != nil {
        logger.Log.Errorf("Error fetching optional prompt response from open-ai. Possibly no api-key is set: %v", openAiErr)
    }

    err := db.AddToSet(config.COLLECTION_NAME, word +" - " + reasonToGo)
    if err != nil {
        // Log the error
        logger.Log.Errorf("Error adding to set: %v", err)

        // Return a JSON-formatted error response
        errorResponse := map[string]string{
            "error": "Failed to add word to collection",
            "details": err.Error(),
        }
        return c.JSON(http.StatusInternalServerError, errorResponse)
    }

    // Create a map or struct to represent the successful JSON response
    response := map[string]string{
        "message": "thank you for your word",
        "word": word,
        "reason" : reasonToGo,
    }

    // Return the successful JSON response
    return c.JSON(http.StatusOK, response)
}


func RemoveWord(c echo.Context) error {
    word, err := db.PopRandomFromSet(config.COLLECTION_NAME)
    if err != nil {
        // Log the error
        logger.Log.Errorf("Error retrieving from set: %v", err)

        // Return a JSON-formatted error response
        errorResponse := map[string]string{
            "error": "Failed to retrieve a word from collection",
            "details": err.Error(),
        }
        return c.JSON(http.StatusInternalServerError, errorResponse)
    }

    reasonToLeave, openAiErr := utils.FetchCompletionFromOpenAI(config.REMOVE_WORD_PROMPT_TEMPLATE + word)
    if openAiErr != nil {
        logger.Log.Errorf("Error fetching optional prompt response from open-ai. Possibly no api-key is set: %v", openAiErr)
    }

    // Create a map or struct to represent the successful JSON response
    response := map[string]string{
        "message": "We are happy to give you back a word",
        "word": word,
        "reason": reasonToLeave,
    }

    // Return the successful JSON response
    return c.JSON(http.StatusOK, response)
}

