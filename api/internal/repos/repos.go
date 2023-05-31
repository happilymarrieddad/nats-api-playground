package repos

import (
	"github.com/pkg/errors"
	"xorm.io/xorm"
)

var errTransaction = errors.New("session required for transactions")

func wrapInTx(db *xorm.Engine, f func(s *xorm.Session) error) error {
	sesh := db.NewSession()
	defer sesh.Close()

	if err := f(sesh); err != nil {
		if seshErr := sesh.Rollback(); seshErr != nil {
			return errors.WithMessage(err, seshErr.Error())
		}
		return err
	}

	if err := sesh.Commit(); err != nil {
		return err
	}

	return nil
}
