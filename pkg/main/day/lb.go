package day

type LeaderBoard struct {
	Event   string   `json:"event"`
	OwnerID string   `json:"owner_id"`
	Members []Member `json:"members"`
}

type Member struct{}
