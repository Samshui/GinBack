package dto

import "Gin/model"

type RecordDto struct {
	SelectedExperiment model.Experiment `json:"selectedExperiment"`
	Selector           model.User       `json:"selector"`
	Date               string           `json:"date"`
	Site               int              `json:"site"`
	Time               int              `json:"time"`
}

func ToRecordDto(record model.Record) RecordDto {
	return RecordDto{
		SelectedExperiment: record.SelectedExperiment,
		Selector:           record.Selector,
		Date:               record.Date,
		Site:               record.Site,
		Time:               record.Time,
	}
}
