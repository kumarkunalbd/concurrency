package retyrservice

type Retry interface {
	Retyr()
	Wait()
}

/*
type RetyrService struct {
	Retry retyr
}

func (rs RetyrService) Retyr() {
	rs.Retry()
}

type URLretry struct {
	attempts    int
	interval    time.Duration
	anOperation interface{}
	// // --> Which I need to work
}

func Newretyr() *URLretry {
	return &URLretry{}
}

func (rt *URLretry) Retyr() {
	for rt.attempts > 0 {
		rt.Wait()
		res := rt.anOperation
		if res == success {
			return
		} // --> Which I need to work upon
		rt.attempts--
	}
}

func (rt *URLretry) Wait() {
	time.Sleep(rt.interval)
}

func (rt *URLretry) AttemptLeft() int {
	return rt.attempts
}

type URLretry struct {
	attempts    int
	interval    time.Duration
	anOperation interface{}
	sshowstoper interface{}
	// // --> Which I need to work
}

func (rt *URLretry) Retyr() {
	fmt.Printf("Own fucntionaluty")
}*/
