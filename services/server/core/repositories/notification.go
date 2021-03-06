package repositories

import (
	"fmt"
	"time"

	"github.com/ilovelili/dongfeng-core/services/server/core/models"
)

// NotificationRepository notification repository
type NotificationRepository struct{}

// NewNotificationRepository init notification repository
func NewNotificationRepository() *NotificationRepository {
	return &NotificationRepository{}
}

// Select select notification logs
func (r *NotificationRepository) Select(uid string, adminonly bool, excluderead bool) (notifications []*models.Notification, err error) {
	query := fmt.Sprintf("CALL spSelectNotifications('%s', %d, %d)", uid, resolveAdminOnly(adminonly), resolveExcludeRead(excluderead))
	err = session().Find(query, nil).All(&notifications)
	if norows(err) {
		err = nil
	}

	return
}

// Upsert upsert notification
func (r *NotificationRepository) Upsert(notifications []*models.Notification) (err error) {
	tx, err := session().Begin()
	if err != nil {
		return
	}

	// upsert by loop
	for _, notification := range notifications {
		query := Table("notifications").Alias("n").
			Where().Eq("n.id", notification.ID).
			Sql()

		var n models.Notification
		err := session().Find(query, nil).Single(&n)
		if err != nil || 0 == n.ID {
			notification.Time = time.Now()
			err = session().InsertTx(tx, notification)
		} else {
			n.Read = 1
			n.Time = time.Now()
			err = session().UpdateTx(tx, n)
		}

		if err != nil {
			session().Rollback(tx)
			return err
		}
	}

	return session().Commit(tx)
}

func resolveAdminOnly(adminonly bool) int {
	return resolveBool(adminonly)
}

func resolveExcludeRead(excluderead bool) int {
	return resolveBool(excluderead)
}

func resolveBool(b bool) int {
	if b {
		return 1
	}
	return 0
}
