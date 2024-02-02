package streaming

import (
	"net/http"

	"github.com/JoVi0li/ocean-server/internal/shared"
	"github.com/JoVi0li/ocean-server/internal/shared/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var service Service
var waitList = make(chan uuid.UUID, 2)
var calls = make(map[uuid.UUID]Connection)

func Configure() {
	service = Service{
		Repository: &RepositoryPostgres{
			Connection: database.Conn,
		},
	}
}

func VoiceCall(ctx *gin.Context) {

	token, err := shared.GetToken(ctx)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	decodedTk, err := shared.DecodeTokenClaims(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	id, idErr := uuid.Parse(decodedTk.ID)
	if idErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  shared.ErrorIdInvalid.Error(),
		})
		return
	}

	waitList <- id
	usersId := make([]uuid.UUID, 0, 2)

	for userId := range waitList {
		usersId = append(usersId, userId)
		if len(usersId) == 2 {
			call, err := service.StartCall(ctx, [2]uuid.UUID(usersId))
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"sucess": false,
					"data":   nil,
					"error":  err.Error(),
				})
				return
			}
			ctx.JSON(http.StatusCreated, gin.H{
				"sucess": true,
				"data":   call,
				"error":  nil,
			})
			return
		}
	}
}

func Connect(ctx *gin.Context) {
	if ok := ctx.IsWebsocket(); !ok {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  nil,
		})
		return
	}

	token, err := shared.GetToken(ctx)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	decodedTk, err := shared.DecodeTokenClaims(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	_, idErr := uuid.Parse(decodedTk.ID)
	if idErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  shared.ErrorIdInvalid.Error(),
		})
		return
	}

	voiceCallId := ctx.Param("id")
	parsedVCId, idVCErr := uuid.Parse(voiceCallId)
	if idVCErr != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  idVCErr.Error(),
		})
		return
	}

	vc, err := service.GetCallById(ctx, parsedVCId)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"sucess": false,
			"data":   nil,
			"error":  err.Error(),
		})
		return
	}

	if conn, ok := calls[vc.ID]; ok {
		if len(conn.Clients) > 1 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"sucess": false,
				"data":   nil,
				"error":  "Connection invalid",
			})
			return
		}
		conn.Clients[ctx] = make(chan []byte)
		defer delete(conn.Clients, ctx)
	} else {
		newConn := Connection{
			VoiceCallID: vc.ID,
			Clients:     Clients{},
		}
		calls[vc.ID] = newConn
		calls[vc.ID].Clients[ctx] = make(chan []byte)
	}

	ctx.Header("Content-Type", "audio/webm")

	go func() {
		for {
			// _, message, err := ctx.Request.Web.ReadMessage()
			// if err != nil {
			// 	log.Println("read:", err)
			// 	close(ch)
			// 	break
			// }
			// // Transmitir a mensagem recebida para todos os outros clientes
			// for _, client := range clients {
			// 	if client != ch {
			// 		client <- message
			// 	}
			// }
		}
	}()
}
