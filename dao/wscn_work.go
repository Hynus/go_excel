package dao

import (
	"go_excel/models"
)

func SaveWscnWorkRecord(record models.WscnWork) {
	notFound := models.DB().Where("date = ?", record.Date).Find(&record).RecordNotFound()
	if notFound {
		models.DB().Save(&record)
	}
}
