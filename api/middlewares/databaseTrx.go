package middlewares

import (
	"go-todoApp/constants"
	"go-todoApp/libs"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DatabaseTransactionMiddleware struct {
	handler libs.RequestHandler
	logger  libs.Logger
	db      libs.Database
}

func statusInList(status int, statusList []int) bool {
	for _, i := range statusList {
		if i == status {
			return true
		}
	}
	return false
}

func NewDatabaseTransactionMiddleware(
	handler libs.RequestHandler,
	logger libs.Logger,
	db libs.Database,
) DatabaseTransactionMiddleware {
	return DatabaseTransactionMiddleware{
		handler: handler,
		logger:  logger,
		db:      db,
	}
}

func (m DatabaseTransactionMiddleware) SetUp() {
	m.logger.Info("Initializing database transaction middleware...")

	m.handler.Gin.Use(func(c *gin.Context) {
		transactionHandle := m.db.Begin()
		m.logger.Info("Starting transaction with database...")
		defer func() {
			if r := recover(); r != nil {
				transactionHandle.Rollback()
			}
		}()

		c.Set(constants.DBTransaction, transactionHandle)
		c.Next()

		if c.Writer.Status() == http.StatusInternalServerError {
			m.logger.Info("Internal Server Error occurred. Rolling back...")
			transactionHandle.Rollback()
		}

		if statusInList(c.Writer.Status(), []int{http.StatusOK, http.StatusCreated}) {
			m.logger.Info("Commiting transaction...")
			if err := transactionHandle.Commit().Error; err != nil {
				m.logger.Error("Error(s) during commit: ", err)
			}
		}
	})
}
