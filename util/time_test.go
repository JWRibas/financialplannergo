package util

import "testing"

func TestStringToDate(t *testing.T) {
	var convertedDate = StringToDate("2023-05-01T15:40:00")

	if convertedDate.Year() != 2023 {
		t.Error("Espera que o ano seja 2023")
	}

	if convertedDate.Month() != 05 {
		t.Error("Espera que o Mês seja 2023")
	}

}
