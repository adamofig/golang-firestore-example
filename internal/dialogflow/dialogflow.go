package dialogflow

import (
	"context"
	"dataclouader/entity"
	"dataclouader/firestoreService"
	"dataclouader/internal/logger"
	"errors"
	"fmt"
	"net/http"

	dialogflow "cloud.google.com/go/dialogflow/apiv2"
	"github.com/gin-gonic/gin"
)

// dialogflow "cloud.google.com/go/dialogflow/apiv2"
// dialogflowpb "google.golang.org/genproto/googleapis/cloud/dialogflow/v2"

func webhook(ctx *gin.Context) {

	//color.Info.Println("Iniciando petición web hoob")
	logger.Info("Iniciando petición a webhook")
	var request entity.Request
	// var entityMap map[string]interface{}

	// Creé un objeto que podría variar dependiendo de las respuestas, este muestra exactamente lo que viene
	// err1 := ctx.BindJSON(&entityMap)
	// if err1 != nil {

	// 	utils.Error("Se encontró un error", err1)
	// }
	//utils.DebugJson(entityMap)

	err := ctx.BindJSON(&request)
	if err != nil {
		logger.Error(err)
	}

	logger.DebugJson(request)

	logger.Info("Intención: ", request.QueryResult.Intent.DisplayName, "  Lenguaje: ", request.QueryResult.LanguageCode,
		"  Confianza: ", request.QueryResult.IntentDetectionConfidence, "  Texto: ", request.QueryResult.QueryText)

	var intent string = request.QueryResult.Intent.DisplayName

	var message = ""

	// TODO Necesito un enum para todos las intenciones
	if intent == "MostrarPalabras" {
		logger.Log("Se tiene la intención de Mostrar palabras buscando en firebase")
		collection, _ := firestoreService.GetAllDocsInCollection("words")

		for i, v := range collection {
			logger.Debug("ind", i, "valor : ", v)
			// message += v["word"].(string) + ","
			message += fmt.Sprint(v["word"], ", ")
		}
	}

	if intent == "MostrarPalabras - custom" {
		word := request.QueryResult.Parameters["any"]
		logger.Log("Se tiene que buscar el significado de la palabra", word)

		message = "Tu palabra es, la que dijiste, " + word.(string)

	}
	//utils.DebugJson(request)

	response := map[string]interface{}{"fulfillmentText": message}

	logger.DebugJson(response)
	// Aqui voy a probar los restultados de la consulta.

	//projectId := "buhoh-learning"

	//DetectIntentText(projectId, request.Session)

	ctx.JSON(http.StatusOK, response)
}

func DetectIntentText(projectID, sessionID, text, languageCode string) (string, error) {
	ctx := context.Background()

	sessionClient, err := dialogflow.NewSessionsClient(ctx)
	if err != nil {
		return "", err
	}
	defer sessionClient.Close()

	if projectID == "" || sessionID == "" {
		return "", errors.New(fmt.Sprintf("Received empty project (%s) or session (%s)", projectID, sessionID))
	}

	sessionPath := fmt.Sprintf("projects/%s/agent/sessions/%s", projectID, sessionID)
	textInput := dialogflowpb.TextInput{Text: text, LanguageCode: languageCode}
	queryTextInput := dialogflowpb.QueryInput_Text{Text: &textInput}
	queryInput := dialogflowpb.QueryInput{Input: &queryTextInput}
	request := dialogflowpb.DetectIntentRequest{Session: sessionPath, QueryInput: &queryInput}

	response, err := sessionClient.DetectIntent(ctx, &request)
	if err != nil {
		return "", err
	}

	queryResult := response.GetQueryResult()
	fulfillmentText := queryResult.GetFulfillmentText()
	return fulfillmentText, nil
}
