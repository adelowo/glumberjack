package glumberjack

// func TestDateBackupName(t *testing.T) {
//
// 	currentTime = func() time.Time {
// 		ti, err := time.Parse(time.DateOnly, "2006-1-12")
//
// 		require.NoError(t, err)
//
// 		return ti
// 	}
//
// 	tt := []struct {
// 		name       string
// 		layout     string
// 		backupFile string
// 	}{
// 		{
// 			name:       "no layout provided. using default layout",
// 			layout:     "",
// 			backupFile: "file-2006-1-12.png",
// 		},
// 	}
//
// 	for _, v := range tt {
// 		t.Run(v.name, func(t *testing.T) {
// 			l := &Logger{}
//
// 			op := DateBackupName(v.layout)
//
// 			op(l)
//
// 			backedUpFile := l.backupNameFunc("file.png")
// 			require.Equal(t, backedUpFile, v.backupFile)
// 		})
// 	}
// }
