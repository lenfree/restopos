package models

type AppStatus struct {
        Status string `json:"status,up"`
}

func Status() AppStatus{
        return  AppStatus{Status: "up"}
}
