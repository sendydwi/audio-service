package audioservice

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	storageutils "github.com/sendydwi/audio-service/util/storage"
	"gorm.io/gorm"
)

type RestHandler struct {
	service AudioService
}

func NewRestHandler(db *gorm.DB) RestHandler {
	return RestHandler{
		service: AudioService{
			Repository: &PosgresRepository{
				database: db,
			},
			Storage: storageutils.GetStorageAccessor("local")},
	}
}

func (h *RestHandler) RegisterHandlerRoutes(r *gin.RouterGroup) {
	r.POST("audio/user/:userId/phrase/:phraseId", h.handleUploadAudioFiles)
	r.GET("audio/user/:userId/phrase/:phraseId/:audioFormat", h.handleGetAudioFiles)
}

func (h *RestHandler) handleUploadAudioFiles(c *gin.Context) {
	userId := c.Param("userId")
	phraseId := c.Param("phraseId")
	fileHeader, err := c.FormFile("audio_file")
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	log.Printf("[AudioRestHandler] upload file with user id: %s, pharse id: %s, and filename: %s\n", userId, phraseId, fileHeader.Filename)
	err = h.service.UploadAudioFile(userId, phraseId, file)
	if err != nil {
		log.Println(err.Error())
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, "upload success")
}

func (h *RestHandler) handleGetAudioFiles(c *gin.Context) {
	userId := c.Param("userId")
	phraseId := c.Param("phraseId")
	audioFormat := c.Param("audioFormat")

	log.Printf("[AudioRestHandler] upload file with user id: %s, pharse id: %s, with file extension: %s\n", userId, phraseId, audioFormat)
	audioData, contentType, err := h.service.GetAudioFile(userId, phraseId, audioFormat)
	if err != nil {
		log.Println(err.Error())
		c.Error(err)
		return
	}
	c.Data(http.StatusOK, contentType, audioData)
}
