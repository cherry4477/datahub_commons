package statistics

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
)

/*
The fact of arbitrary chars in tag may be a problem for stat key.
This may make some tag stat keys are duplicated with some user stat key.
So it is best to avoid $#& in tag name.
*/

func GetGeneralStatKey(words ...string) string {
	return strings.Join(words, "/")
}

func GetSubscriptionsStatKey(words ...string) string {
	return fmt.Sprintf("%s%s%s", GetGeneralStatKey(words...), "#", "subs")
}

func GetSubscriptionPlanSigningTimesStatKey(words ...string) string { // params should be (repoName, itemName, planId string)
	return fmt.Sprintf("%s%s%s", GetGeneralStatKey(words...), "#", "sgns")
}

func GetTransactionsStatKey(words ...string) string {
	return fmt.Sprintf("%s%s%s", GetGeneralStatKey(words...), "#", "txns")
}

func GetStarsStatKey(words ...string) string {
	return fmt.Sprintf("%s%s%s", GetGeneralStatKey(words...), "#", "strs")
}

func GetCommentsStatKey(words ...string) string {
	return fmt.Sprintf("%s%s%s", GetGeneralStatKey(words...), "#", "cmts")
}

// item doesn't mean data item. It means any objects.

func GetUserItemStatKey(username string, itemStatKey string) string {
	return fmt.Sprintf("%s$%s", username, itemStatKey)
}

func GetUserSubscriptionPlanSigningTimesStatKey(userName, repoName, itemName, planId string) string {
	return GetUserItemStatKey(userName, GetSubscriptionPlanSigningTimesStatKey(repoName, itemName, planId))
}

// user stats
/*
func GetUserSubscriptionsStatKey(username string) string {
	return fmt.Sprintf("%s$#%s", username, "subs")
}

func GetUserTransactionsStatKey(username string) string {
	return fmt.Sprintf("%s$#%s", username, "txns")
}

func GetUserStarsStatKey(username string) string {
	return fmt.Sprintf("%s$#%s", username, "strs")
}

func GetUserCommentsStatKey(username string) string {
	return fmt.Sprintf("%s$#%s", username, "cmts")
}
*/

func GetDateSubscriptionsStatKey(date time.Time) string {
	return fmt.Sprintf("%s>%s", date.Format("2006-01-02"), "subs")
}
func GetDateTransactionsStatKey(date time.Time) string {
	return fmt.Sprintf("%s>%s", date.Format("2006-01-02"), "txns")
}

//==========================================================
//
//==========================================================

func UpdateStat(db *sql.DB, key string, delta int) (int, error) {
	return updateOrSetStat(db, key, delta, true)
}

func SetStat(db *sql.DB, key string, delta int) (int, error) {
	return updateOrSetStat(db, key, delta, false)
}

func updateOrSetStat(db *sql.DB, key string, delta int, isUpdate bool) (int, error) {
	sqlget := `select STAT_VALUE from DH_ITEM_STAT where STAT_KEY=?`

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	stat := 0
	err = tx.QueryRow(sqlget, key).Scan(&stat)
	if err != nil {
		if err != sql.ErrNoRows {
			tx.Rollback()
			return 0, err
		}

		stat = delta
		if stat <= 0 {
			tx.Rollback()
			return 0, errors.New("stat delta can't be <= 0")
		}

		sqlinsert := `insert into DH_ITEM_STAT (STAT_KEY, STAT_VALUE) values (?, ?)`
		_, err := tx.Exec(sqlinsert, key, stat)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	} else {
		if isUpdate {
			stat = stat + delta
		} else {
			stat = delta
		}

		// needed?
		//if stat < 0 {
		//	stat = 0
		//}

		sqlupdate := `update DH_ITEM_STAT set STAT_VALUE=? where STAT_KEY=?`
		_, err := tx.Exec(sqlupdate, stat, key)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	tx.Commit()

	return stat, nil
}

func RetrieveStat(db *sql.DB, key string) (int, error) {
	stat := 0
	sqlstr := `select STAT_VALUE from DH_ITEM_STAT where STAT_KEY=?`
	err := db.QueryRow(sqlstr, key).Scan(&stat)
	switch {
	case err == sql.ErrNoRows:
		return 0, nil
	case err != nil:
		return 0, err
	default:
		return stat, nil
	}
}

// todo: maybe it is better to do this in a txn
func RemoveStat(db *sql.DB, key string) (int, error) {
	num, err := RetrieveStat(db, key)
	if err != nil {
		return 0, err
	}
	if num == 0 {
		return 0, nil
	}
	
	sqlstr := `delete from DH_ITEM_STAT where STAT_KEY=?`
	_, err = db.Exec(sqlstr, key)
	switch {
	case err == sql.ErrNoRows:
		return 0, nil
	case err != nil:
		return 0, err
	default:
		return num, nil
	}
}
