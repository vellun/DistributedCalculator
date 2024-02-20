package agent

import "distributed-calculator/orchestrator/pkg/database"

type CompResources struct { // Структура со всеми вычислительными ресурсами(агентами)
	Agents []*Agent
}

var Resources *CompResources = NewResources()

func NewResources() *CompResources {
	return &CompResources{[]*Agent{}}
}

func (cr *CompResources) Init() { // При инициализации из базы берется 3 дефолтных агента
	db_agents, _ := database.GetAllAgents()
	for _, ag := range db_agents {
		// Формируем список с агентами, чтобы в дальнейшем отслеживать их состояние
		cr.Agents = append(cr.Agents, &Agent{Id: ag.Id, Last_active: int64(ag.Id), Status: ag.Status})
	}
}
