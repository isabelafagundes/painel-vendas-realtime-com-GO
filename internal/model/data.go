package model

import "time"

type DataHora struct {
	Valor string
}

const (
	LayoutISOData     = "2006-01-02"          // 2025-12-06
	LayoutISODataHora = "2006-01-02 15:04:05" // 2025-12-06 22:30:00
	LayoutBRData      = "02/01/2006"          // 06/12/2025
	LayoutBRDataHora  = "02/01/2006 15:04:05" // 06/12/2025 22:30:00
)

func DataHoje() *DataHora {
	return &DataHora{
		Valor: time.Now().Format(LayoutISOData),
	}
}

func DataHojeBR() *DataHora {
	return &DataHora{
		Valor: time.Now().Format(LayoutBRData),
	}
}

func DataAgora() *DataHora {
	return &DataHora{
		Valor: time.Now().Format(LayoutISODataHora),
	}
}

func DataAgoraBR() *DataHora {
	return &DataHora{
		Valor: time.Now().Format(LayoutBRDataHora),
	}
}

func NovaDataHora(t time.Time) *DataHora {
	return &DataHora{
		Valor: t.Format(LayoutISODataHora),
	}
}

func NovaDataHoraComLayout(t time.Time, layout string) *DataHora {
	return &DataHora{
		Valor: t.Format(layout),
	}
}

func DataHoraDeString(valor string) *DataHora {
	return &DataHora{
		Valor: valor,
	}
}

func (d *DataHora) Time(layout string) (time.Time, error) {
	return time.Parse(layout, d.Valor)
}
