package scheduler

type Topic string

type Scheduler struct {
	Cores map[Topic]Core
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		Cores: make(map[Topic]Core),
	}
}

func (s *Scheduler) Stop() {
	for _, c := range s.Cores {
		c.Shut()
	}
}

func (s *Scheduler) ShutTopic(topic Topic) {
	s.Cores[topic].Shut()
}

func (s *Scheduler) AppendCore(topic Topic, c Core) {
	s.Cores[topic] = c
	s.Cores[topic].Listen()
}

func (s *Scheduler) Push(topic Topic, event interface{}) {
	s.Cores[topic].Push(event)
}
