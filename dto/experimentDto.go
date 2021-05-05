package dto

import "Gin/model"

type ExperimentDto struct {
	Name        string `json:"name"`
	Eid         string `json:"eid"`
	ELabel      int    `json:"eLabel"`
	Lab         string `json:"lab"`
	TimeMorning int    `json:"timeMorning"`
	TimeNoon    int    `json:"timeNoon"`
	TimeEvening int    `json:"timeEvening"`
	SiteSize    int    `json:"SiteSize"`
}

func ToExperimentDto(experiment model.Experiment) ExperimentDto {
	return ExperimentDto{
		Name:        experiment.Name,
		Eid:         experiment.Eid,
		ELabel:      experiment.ELabel,
		Lab:         experiment.Lab,
		TimeMorning: experiment.TimeMorning,
		TimeNoon:    experiment.TimeNoon,
		TimeEvening: experiment.TimeEvening,
		SiteSize:    experiment.SiteSize,
	}
}
