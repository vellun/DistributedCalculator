package models

type Agent struct {
	Id        int    `json:"id"`
	Status    string `json:"status"` //running/missing/dead
	Last_active int    `json:"last_active"`
	Goroutines int  // Количество горутин, которые сейчас задействует агент
}
