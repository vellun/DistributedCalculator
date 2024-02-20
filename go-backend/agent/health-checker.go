package agent

import (
	"distributed-calculator/orchestrator/pkg/database"
	"fmt"
	"time"
)

// Переменные берутся из файла конфига
var (
	// Интервал времени между проверками состояния агентов
	health_check_interval time.Duration = 40 * time.Second
	// Максимально допустимое время неактивности агента чтобы он перешел в статус missing
	max_inactive_time time.Duration = 120 * time.Second
	// Максимально допустимое время неактивности агента чтобы он перешел в статус dead
	very_max_inactive_time time.Duration = 200 * time.Second
)

var health_ticker = time.NewTicker(health_check_interval)

// Горутина для переодической проверки состояния агентов и отлавливания умерших
func RunHealthChecker() {
	fmt.Println("Health checker is running")
	for {
		go func() {
			<-health_ticker.C // Если интервал прошел

			for _, agent := range Resources.Agents { // Чекаем всех агентов

				if agent.Status == "dead" {
					ReplaceDeadAgent(agent) // Если помер, отправляется на замену
					fmt.Printf("Агент %d был заменен\n", agent.Id)
				}

				// Если с момента последней зафиксированной активности агента прошло больше назначенного времени
				sub := time.Now().Sub(time.Unix(int64(agent.Last_active), 0)).Seconds()
				if sub > max_inactive_time.Seconds() && agent.Status == "running" {
					fmt.Printf("Агент %d долго неактивен. Статус переходит в 'missing'\n", agent.Id)
					agent.Status = "missing"
					database.UpdateStatus(agent.Id, "missing") // Меняем статус агента в бд
				}
				// Если агент пропал совсем надолго
				if sub > very_max_inactive_time.Seconds() && agent.Status == "missing" {
					fmt.Printf("Агент %d слишком долго неактивен. Статус переходит в 'dead'\n", agent.Id)
					agent.Status = "dead"
					database.UpdateStatus(agent.Id, "dead") // Меняем статус агента в бд(во время следующей проверки умерший агент будет заменен)
				}
				if agent.Status == "running" {
					fmt.Printf("Health checker: Агент %d в порядке. Продолжаем работу\n", agent.Id)
				}
			}

		}()
		time.Sleep(health_check_interval + 2)
	}
}

// Замена агента, с которым потеряна связь
func ReplaceDeadAgent(agent *Agent) {
	for i := range Resources.Agents {
		if i+1 == agent.Id {
			Resources.Agents = append(Resources.Agents[:i], Resources.Agents[i+1:]...)
			break
		}
	}
	new := &Agent{Id: agent.Id, Last_active: time.Now().Unix(), Status: "running"}
	// Добавляем нового агента и присваиваем ему id старого
	Resources.Agents = append(Resources.Agents, new)
	database.UpdateStatus(agent.Id, "running") // В бд обновляем у старого агента статус на running
	go new.RunAgent()                          // Запуск нового агента
	fmt.Printf("New agent %d is running", new.Id)
}
