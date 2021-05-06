package dto

import "Gin/model"

type RecordDto struct {
	Date           string `json:"date"`
	Site           int    `json:"site"`
	Time           int    `json:"time"`
	ExperimentName string `json:"experimentName"`
	ExperimentID   string `json:"experimentID"`
	Lab            string `json:"lab"`
}

func ToRecordDto(record model.Record) RecordDto {
	return RecordDto{
		Date:           record.Date,
		Site:           record.Site,
		Time:           record.Time,
		ExperimentName: record.ExperimentName,
		ExperimentID:   record.ExperimentID,
		Lab:            record.Lab,
	}
}
