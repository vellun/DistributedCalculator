package agent

import (
	"distributed-calculator/orchestrator/pkg/database"
	"fmt"
	"time"
)

const (
	health_check_interval  time.Duration = 10 * time.Second  // Интервал времени между проверками состояния агентов
	max_inactive_time      time.Duration = 120 * time.Second // Максимально допустимое время неактивности агента чтобы он перешел в статус missing
	very_max_inactive_time time.Duration = 180 * time.Second // Максимально допустимое время неактивности агента чтобы он перешел в статус dead
)

var health_ticker = time.NewTicker(health_check_interval)

// Горутина для переодичнской проверки состояния агентов и отлавливания умерших
func RunHealthChecker() {
	fmt.Println("Health checker is running")
	for {
		go func() {
			<-health_ticker.C // Если интервал прошел

			for _, agent := range Resources.Agents {
				// Если с момента последней зафиксированной активности агента прошло больше назначенного времени
				if time.Duration(time.Now().Sub(time.Unix(int64(agent.Last_active), 0))/time.Second) > max_inactive_time {
					fmt.Printf("Агент %d долго неактивен. Статус переходит в 'missing'\n", agent.Id)
					agent.Status = "missing"
					database.UpdateStatus(agent.Id, "missing") // Меняем статус агента в бд
				}
				// Если агент пропал совсем надолго
				if time.Duration(time.Now().Sub(time.Unix(int64(agent.Last_active), 0))/time.Second) > very_max_inactive_time {
					fmt.Printf("Агент %d долго неактивен. Статус переходит в 'dead'\n", agent.Id)
					agent.Status = "dead"
					database.UpdateStatus(agent.Id, "dead") // Меняем статус агента в бд
				} else {
					fmt.Printf("Health checker: Агент %d в порядке. Продолжаем работу\n", agent.Id)
				}
			}

		}()
		time.Sleep(health_check_interval + 2)
	}
}
