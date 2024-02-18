package agent

type CompResources struct { // Структура со всеми вычислительными ресурсами, которые у нас есть
	Agents []*Agent
}

func NewResources() *CompResources {
	return &CompResources{[]*Agent{}}
}

func (cr *CompResources) Init() { // При инициализации создается 3 агента
	for i := 1; i < 4; i++ {
		agent := NewAgent(i)
		cr.Agents = append(cr.Agents, agent)
	}
}
