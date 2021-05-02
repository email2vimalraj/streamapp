package api

import (
	"net/http"

	db "github.com/email2vimalraj/streamapp/db/sqlc"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type createStreamRequest struct {
	Username   string `json:"username" binding:"required"`
	StreamName string `json:"stream_name" binding:"required,min=6"`
	StreamLink string `json:"stream_link" binding:"required"`
}

func (server *Server) createStream(ctx *gin.Context) {
	var req createStreamRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateStreamParams{
		StreamName: req.StreamName,
		StreamLink: req.StreamLink,
		Username:   req.Username,
	}

	stream, err := server.store.CreateStream(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			case "foreign_key_violation", "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, stream)
}
