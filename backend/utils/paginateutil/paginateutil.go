package paginateutil

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var DEFAULT_LIMIT = 10
var DEFAULT_SKIP = 0

type Paginate struct {
	Limit int `json:"limit"`
	Skip  int `json:"skip"`
}

func getIntFromQuery(ctx *gin.Context, queryName string, defaultValue int) int {
	valueStr, queryExists := ctx.GetQuery(queryName)
	if !queryExists {
		return defaultValue
	}
	value, err := strconv.Atoi(valueStr)
	if err != nil {
		return defaultValue
	}
	return value
}

func NewPaginateFromApi(ctx *gin.Context) *Paginate {
	return &Paginate{
		Limit: getIntFromQuery(ctx, "limit", DEFAULT_LIMIT),
		Skip:  getIntFromQuery(ctx, "skip", DEFAULT_SKIP),
	}
}

func ApplyGormPaginate(tx *gorm.DB, paginate *Paginate) *gorm.DB {
	if paginate == nil {
		paginate = &Paginate{Limit: DEFAULT_LIMIT, Skip: DEFAULT_SKIP}
	}
	return tx.Limit(paginate.Limit).Offset(paginate.Skip)
}
