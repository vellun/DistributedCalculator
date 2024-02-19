package agent

import "distributed-calculator/orchestrator/pkg/database"

type CompResources struct { // Структура со всеми вычислительными ресурсами
	Agents []*Agent
}

var Resources *CompResources = NewResources()

func NewResources() *CompResources {
	return &CompResources{[]*Agent{}}
}

func (cr *CompResources) Init() { // При инициализации создается 3 агента
	// for i := 1; i < 4; i++ {
	// 	agent := NewAgent(i)
	// 	cr.Agents = append(cr.Agents, agent)
	// }
	db_agents, _ := database.GetAllAgents()
	for _, ag := range db_agents {
		cr.Agents = append(cr.Agents, &Agent{Id: ag.Id, Last_active: int64(ag.Id), Status: ag.Status})
	}
}
