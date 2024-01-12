package db

// func GetSettings(UID int) model.Settings {
// 	var res model.Settings
// 	db := DBSETTINGS.Open()
// 	defer db.Close()
// 	rows, err := db.Query("SELECT id, bg_color FROM settings WHERE owner_id = ?", UID)
// 	if err != nil {
// 		fmt.Println(err)
// 		fmt.Println("from getSettings")
// 		return res
// 	}
// 	for rows.Next() {

// 		err := rows.Scan(&res.Id, &res.BgColor)
// 		if err != nil {

// 			fmt.Println(err)
// 			fmt.Println("from getSettings on first scan")
// 			return res
// 		}
// 	}

// 	return res
// }
