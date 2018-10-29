//provide interface of the operation on data
/* warning: without fully test */
/* remember to tell me if there are bugs */
package entity

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Username string
	Password string
	Mail     string
	Phone    string
}

type Meeting struct {
	Title       string
	Sponsor     string
	StartTime   string
	EndTime     string
	Participant []string
}

var (
	current_user  User
	total_user    []User
	total_meeting []Meeting
	userLib       string = "user.json"
	meetingLib    string = "meeting.json"
)

func Init() {
	ReadUserFile()
	ReadMeetingFile()
}

func UpdataLib() {
	WriteUserFile()
	WriteMeetingFile()
}

func GetAllUser() []User {
	return total_user
}

func GetAllMeeting() []Meeting {
	return total_meeting
}

func ReadUserFile() {
	file, err := os.Open(userLib)
	if err != nil {
		return
	}
	state, _ := file.Stat()
	if state.Size() == 0 {
		return
	}
	buffer := make([]byte, state.Size())
	_, err = file.Read(buffer)
	if err != nil {
		return
	}
	buffer = []byte(os.ExpandEnv(string(buffer)))
	err = json.Unmarshal(buffer, &total_user)
	if err != nil {
		return
	}
}

func ReadMeetingFile() {
	file, err := os.Open(meetingLib)
	if err != nil {
		return
	}
	state, _ := file.Stat()
	if state.Size() == 0 {
		return
	}
	buffer := make([]byte, state.Size())
	_, err = file.Read(buffer)
	//buffer, err = StripComments(buffer)
	if err != nil {
		return
	}
	buffer = []byte(os.ExpandEnv(string(buffer)))
	err = json.Unmarshal(buffer, &total_meeting)
	if err != nil {
		return
	}
}

func WriteUserFile() {
	userRec, err := json.Marshal(total_user)
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create(userLib)
	defer f.Close()
	f.WriteString(string(userRec))
}

func WriteMeetingFile() {
	meetingRec, err := json.Marshal(total_user)
	if err != nil {
		fmt.Println(err)
	}
	f, _ := os.Create(meetingLib)
	defer f.Close()
	f.WriteString(string(meetingRec))
}

func IsUserExist_Login(name string, psw string) bool {
	for i := 0; i < len(total_user); i++ {
		if total_user[i].Username == name && total_user[i].Password == psw {
			return true
		}
	}
	return false
}

func UserCheck(new_user User) bool {
	for i := 0; i < len(total_user); i++ {
		if total_user[i].Username == new_user.Username {
			return false
		}
	}
	return true
}

func UsernameCheck(name string) bool {
	for i := 0; i < len(total_user); i++ {
		if total_user[i].Username == name {
			return true
		}
	}
	return false
}

func CreateUser(name string, psw string, ma string, ph string) int {
	user := User{
		Username: name,
		Password: psw,
		Mail:     ma,
		Phone:    ph,
	}
	if UserCheck(user) {
		_, err := json.Marshal(user)

		if err != nil {
			return 1
		}

		total_user = append(total_user, user)

		return 0
	} else {
		return 2
	}
}

func DeleteUser(name string) bool {
	var i int
	for i = 0; i < len(total_user); i++ {
		if total_user[i].Username == name {
			break
		}
	}
	total_user[i] = total_user[len(total_user)-1]
	total_user = total_user[0 : len(total_user)-1]
	return true
}

func MeetingCheck(t string) int {
	for i := 0; i < len(total_meeting); i++ {
		if total_meeting[i].Title == t {
			return i
		}
	}
	return -1
}

func CreateMeeting(t string, s string, st string, et string, p []string) bool {
	if len(p) == 0 {
		return false
	}
	meeting := Meeting{
		Title:       t,
		Sponsor:     s,
		StartTime:   st,
		EndTime:     et,
		Participant: p,
	}
	total_meeting = append(total_meeting, meeting)
	return true
}

func DeleteMeeting(t string, name string) int {
	if name == "" || UsernameCheck(name) {
		pos := MeetingCheck(t)
		if total_meeting[pos].Sponsor == name {
			total_meeting[pos] = total_meeting[len(total_meeting)-1]
			total_meeting = total_meeting[0 : len(total_meeting)-1]
			return 0
		} else {
			return 1
		}
	} else {
		return 2
	}
}

func AddMeetingParticipant(t string, player string) int {
	if UsernameCheck(player) {
		pos := MeetingCheck(t)
		if pos == -1 {
			return 1
		}
		total_meeting[pos].Participant = append(total_meeting[pos].Participant, player)
		return 0
	} else {
		return 2
	}
}

func DeleteMeetingParticipant(t string, player string) int {
	if UsernameCheck(player) {
		pos := MeetingCheck(t)
		if pos == -1 {
			return 1
		}
		var i, size int = 0, len(total_meeting[pos].Participant)
		for i = 0; i < size; i++ {
			if total_meeting[pos].Participant[i] == player {
				break
			}
		}
		if i == size {
			return 1
		} else {
			if len(total_meeting[pos].Participant) == 1 {
				total_meeting[pos] = total_meeting[len(total_meeting)-1]
				total_meeting = total_meeting[0 : len(total_meeting)-1]
			} else {
				total_meeting[pos].Participant[i] = total_meeting[pos].Participant[size-1]
				total_meeting[pos].Participant = total_meeting[pos].Participant[0 : size-1]
			}
		}
		return 0
	} else {
		return 2
	}
}

func DeleteAllMeeting(name string, meetingId []string) int {
	if UsernameCheck(name) {
		for i := 0; i < len(meetingId); i++ {
			flag := DeleteMeeting(meetingId[i], "")
			if flag != 0 {
				return flag
			}
		}
		return 0
	} else {
		return 2
	}
}
