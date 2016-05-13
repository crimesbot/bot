package errs

import (
	"fmt"
	"os"
	"time"

	"github.com/crimesbot/bot/session"

	"github.com/crimesbot/bot/settings"
	"github.com/pkg/errors"
)

func Catch(e error, s *session.Session) {
	if e != nil {
		LogError(e, s)
	}
}

// This methods should send to admins
// every error caught and send something to user
func LogError(e error, s *session.Session) {

	fmt.Println("---------- ERROR ----------")
	defer fmt.Println("---------------------------")

	f, err := os.OpenFile(settings.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		errors.Fprint(os.Stdout, errors.Wrap(err, "openning error log file"))
		return
	}
	defer f.Close()

	t := time.Now()
	l := fmt.Sprintf("%s - System error\n", t.Format(time.Stamp))
	f.WriteString(l)
	fmt.Printf(l)

	errors.Fprint(f, e)
	errors.Fprint(os.Stdout, e)
}
