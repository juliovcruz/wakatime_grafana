package model

import "fmt"

func (l Language) convertToSQL() string {
	return fmt.Sprintf(
		"INSERT INTO public.languages "+
			"(id, email, hours, \"name\", \"date\")"+
			" VALUES('%s'::uuid, '%s', %v, '%s', '%s');",
		l.ID, l.Email, l.Hours, l.Name, l.Date.Format("2006-01-02 15:04:05.000"),
	)
}

func ConvertToBackupSQL(languages []Language) string {
	str := ""
	for _, lang := range languages {
		str = str + lang.convertToSQL() + "\n"
	}
	return str
}
