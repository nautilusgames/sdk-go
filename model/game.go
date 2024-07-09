package model

import "time"

type GameResponse struct {
	Games []Game `json:"games"`
}

type Game struct {
	ID           string `json:"id"`
	ConditionIds any    `json:"conditionIds"`
	URL          string `json:"url"`
	BackendPath  string `json:"backendPath"`
	GameID       string `json:"gameId"`
	Status       string `json:"status"`
	Categories   struct {
		Values []struct {
			ID          string `json:"id"`
			Name        string `json:"name"`
			Description string `json:"description"`
		} `json:"values"`
	} `json:"categories"`
	CreatedAt    time.Time `json:"createdAt"`
	Orientations struct {
		Values []string `json:"values"`
	} `json:"orientations"`
	TenantID         string `json:"tenantId"`
	Version          string `json:"version"`
	StaffGameVersion string `json:"staffGameVersion"`
	Labels           struct {
		Values []string `json:"values"`
	} `json:"labels"`
	StaffGameID string `json:"staffGameId"`
	Assets      struct {
		Values []struct {
			ID          string `json:"id"`
			Lang        string `json:"lang"`
			Name        string `json:"name"`
			Slug        string `json:"slug"`
			Description string `json:"description"`
			Icon        struct {
				Value string `json:"value"`
				Files []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					Size string `json:"size"`
					Path string `json:"path"`
				} `json:"files"`
			} `json:"icon"`
			Thumbnail struct {
				Value string `json:"value"`
				Files []struct {
					ID   string `json:"id"`
					Name string `json:"name"`
					Type string `json:"type"`
					Size string `json:"size"`
					Path string `json:"path"`
				} `json:"files"`
			} `json:"thumbnail"`
		} `json:"values"`
	} `json:"assets"`
	IsDemo       bool   `json:"isDemo"`
	Weight       string `json:"weight"`
	CustomLabels struct {
		Values []struct {
			Value string `json:"value"`
		} `json:"values"`
	} `json:"customLabels"`
	StaffGameStatus string `json:"staffGameStatus"`
	IsAutoVersion   bool   `json:"isAutoVersion"`
}
