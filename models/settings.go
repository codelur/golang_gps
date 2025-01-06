package models

type Settings struct {
	ShowLatitude    bool `json:"showLatitude"`
	ShowLongitude   bool `json:"showLongitude"`
	ShowFactoryId   bool `json:"showFactoryId"`
	ShowDriveStatus bool `json:"showDriveStatus"`
	ShowState       bool `json:"showState"`
	ShowName        bool `json:"showName"`
	ShowMake        bool `json:"showMake"`
	ShowModel       bool `json:"showModel"`
	ShowOnline      bool `json:"showOnline"`
}
