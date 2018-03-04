package main

import (
	"fmt"
	"github.com/excelize"
	"go_excel/dao"
	"go_excel/models"
	"go_excel/timeutil"
	"os"
)

func main() {
	xlsx, err := excelize.OpenFile("excels/wscn_work.xlsx")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	rows := xlsx.GetRows("Sheet1")
	for i := 1; i < len(rows); i++ {
		singleRecord := models.WscnWork{
			Date:       timeutil.ParseStringTime(rows[i][0]),
			GetUpTime:  timeutil.ParseStringTime(rows[i][0] + " " + rows[i][1]),
			GetOffRoom: timeutil.ParseStringTime(rows[i][0] + " " + rows[i][2]),
			Busline:    rows[i][3],
			BusGo:      timeutil.ParseStringTime(rows[i][0] + " " + rows[i][4]),
			MetroArr:   timeutil.ParseStringTime(rows[i][0] + " " + rows[i][5]),
			OfficeArr:  timeutil.ParseStringTime(rows[i][0] + " " + rows[i][6]),
			Duration: timeutil.CalcTimeDura(timeutil.ParseStringTime(rows[i][0]+" "+rows[i][2]),
				timeutil.ParseStringTime(rows[i][0]+" "+rows[i][6])),
			Mark: rows[i][7],
		}
		dao.SaveWscnWorkRecord(singleRecord)
	}
}
