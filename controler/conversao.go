package controler

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

//PegarDataAtualTime pega a data atual
func PegarDataAtualTime() string {
	//Pega data atual servido
	t := time.Now()
	//Get date current and zone Recife
	location, err := time.LoadLocation("America/Recife")
	if err != nil {
		fmt.Println(err)
	}
	//Format patten BR
	dataRecife := t.In(location).Format("02/01/2006 15:04:05")
	//dataRecife := t.In(location)

	return dataRecife
}

//ConverterDataAtual pega a data atual
func ConverterDataAtual(date string) string {

	const (
		layoutISO = "2006-01-02 15:04:05"
		layoutUS  = "02/01/2006 15:04:05"
	)

	//date := "2020-05-14"
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t) // 1999-12-31 00:00:00 +0000 UTC
	retornoData := t.Format(layoutUS)

	return retornoData
}

//ConverterDataFormatBanco pega a data atual
func ConverterDataFormatBanco(date string) string {

	const (
		layoutISO = "2006-01-02 15:04:05"
		layoutUS  = "02/01/2006 15:04:05"
	)

	//date := "2020-05-14"
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t) // 1999-12-31 00:00:00 +0000 UTC
	//retornoData := t.Format(layoutUS)

	return t.String()
}

//PegarDataAtualStringTeste pega a data atual
func PegarDataAtualStringTeste() time.Time {
	//Pega data atual servido
	t := time.Now()
	//Get date current and zone Recife
	location, err := time.LoadLocation("America/Recife")
	if err != nil {
		fmt.Println(err)
	}
	//Format patten BR
	//dataRecife := t.In(location).Format("02/01/2006 15:04:05")
	dataRecife := t.In(location)

	return dataRecife
}

//PegarDataAtualString pega a data atual
func PegarDataAtualString() string {
	//Pega data atual servido
	t := time.Now()
	//Get date current and zone Recife
	location, err := time.LoadLocation("America/Recife")
	if err != nil {
		fmt.Println(err)
	}
	//Format patten BR
	dataRecife := t.In(location).Format("02/01/2006 15:04:05")

	return dataRecife
}

//PegarDataAtualTime1 pega a data atual
func PegarDataAtualTime1() time.Time {
	//Pega data atual servido
	t := time.Now()
	//Get date current and zone Recife
	location, err := time.LoadLocation("America/Recife")
	if err != nil {
		fmt.Println("Erro na conversao", err)
	}
	//Format patten BR
	dataRecife := t.In(location).Format("02/01/2006 15:04:05")

	layout := "02/01/2006 15:04:05"
	t2, err := time.Parse(layout, dataRecife)

	fmt.Println(" Dados da data ", t2)

	//Pega data atual servido
	return t2
}

//FormatDateTime padrao
func FormatDateTime(date time.Time) (time.Time, error) {

	data, err := time.Parse("02/01/2006", date.String())

	if err != nil {
		fmt.Println(err)
	}

	return data, err

}

//FormatDateString padrao
func FormatDateString(date string) string {
	//Pega data atual servido
	t := time.Now()
	//Get date current and zone Recife
	location, err := time.LoadLocation("America/Recife")
	if err != nil {
		fmt.Println(err)
	}
	//Format patten BR
	dataRecife := t.In(location).Format("02/01/2006 15:04:05")

	return dataRecife
}

//PegarDataAtualStringAntigo pega a data atual
func PegarDataAtualStringAntigo() string {
	timeString := time.Now().Format("02/01/2006 03:04:05")
	return timeString
}

//PegarDataAtualStringNew pega a data atual
func PegarDataAtualStringNew() string {
	timeString := time.Now().Format("2006-01-02 15:04:05")
	return timeString
}

//StringForInt converter string para int
func StringForInt(id string) int {
	idInt, erro := strconv.Atoi(id)

	if erro != nil {
		fmt.Println(erro)
		return idInt
	}

	return idInt
}

//StringForBool converter string para int
func StringForBool(dados string) bool {
	var err error
	var b bool
	if dados == "true" {
		b, err = strconv.ParseBool("true")
	} else {
		b, err = strconv.ParseBool("false")
	}

	if err != nil {
		fmt.Println(err)
		return b
	}

	return b
}

//FormatStringDate teste
func FormatStringDate(dateSring string) string {

	fmt.Println(" String dateSring  ", dateSring)

	dateSring = strings.ReplaceAll(dateSring, "T", " ")
	dateSring = strings.ReplaceAll(dateSring, "Z", "")

	fmt.Println(" String formatada  ", dateSring)

	return dateSring

}
