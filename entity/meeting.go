/* define meeting */
package entity

type Meeting struct {
	Title         string
	Sponsor       string
	StartTime     string
	EndTime       string
	Participators []string
}

func (meeting *Meeting) InitMeeting(t, s, st, et string, p []string) {
	meeting.Title = t
	meeting.Sponsor = s
	meeting.StartTime = st
	meeting.EndTime = et
	meeting.Participators = p
}

func (meeting *Meeting) getTitle() string {
	return meeting.Title
}

func (meeting *Meeting) setTitle(t string) {
	meeting.Title = t
}

func (meeting *Meeting) getSponsor() string {
	return meeting.Sponsor
}

func (meeting *Meeting) setSponsor(s string) {
	meeting.Sponsor = s
}

func (meeting *Meeting) getStartTime() string {
	return meeting.StartTime
}

func (meeting *Meeting) setStartTime(st string) {
	meeting.StartTime = st
}

func (meeting *Meeting) getEndTime() string {
	return meeting.EndTime
}

func (meeting *Meeting) setEndTime(et string) {
	meeting.EndTime = et
}

func (meeting *Meeting) getParticipators() []string {
	return meeting.Participators
}

func (meeting *Meeting) setParticipators(p []string) {
	meeting.Participators = p
}

func (meeting *Meeting) isParticipator(n string) int {
	for i := 0; i < len(meeting.Participators); i++ {
		if meeting.Participators[i] == n {
			return i
		}
	}
	return -1
}
