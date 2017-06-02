// generated by jsonenums -type=InternalState; DO NOT EDIT

package errs

import (
	"encoding/json"
	"fmt"
)

var (
	_InternalStateNameToValue = map[string]InternalState{
		"Mochasippi":   Mochasippi,
		"Breve":        Breve,
		"Papi":         Papi,
		"Instant":      Instant,
		"Turkish":      Turkish,
		"Irish":        Irish,
		"Crema":        Crema,
		"Cappuccino":   Cappuccino,
		"Unknown":      Unknown,
		"Latte":        Latte,
		"Flatwhite":    Flatwhite,
		"Melange":      Melange,
		"Cortado":      Cortado,
		"Galao":        Galao,
		"Kopisusu":     Kopisusu,
		"Affogato":     Affogato,
		"Macchiato":    Macchiato,
		"Bicerin":      Bicerin,
		"Bombón":       Bombón,
		"Mocha":        Mocha,
		"Caphesuada":   Caphesuada,
		"Carajillo":    Carajillo,
		"Espresso":     Espresso,
		"Eiskaffee":    Eiskaffee,
		"Frappuccino":  Frappuccino,
		"Iced":         Iced,
		"Indianfilter": Indianfilter,
		"Kopiluwak":    Kopiluwak,
		"Kopitubruk":   Kopitubruk,
		"Vienna":       Vienna,
		"Yuanyang":     Yuanyang,
		"None":         None,
		"Americano":    Americano,
		"Cubano":       Cubano,
		"Zorro":        Zorro,
		"Doppio":       Doppio,
		"Romano":       Romano,
		"Guillermo":    Guillermo,
		"Ristretto":    Ristretto,
		"Antoccino":    Antoccino,
	}

	_InternalStateValueToName = map[InternalState]string{
		Mochasippi:   "Mochasippi",
		Breve:        "Breve",
		Papi:         "Papi",
		Instant:      "Instant",
		Turkish:      "Turkish",
		Irish:        "Irish",
		Crema:        "Crema",
		Cappuccino:   "Cappuccino",
		Unknown:      "Unknown",
		Latte:        "Latte",
		Flatwhite:    "Flatwhite",
		Melange:      "Melange",
		Cortado:      "Cortado",
		Galao:        "Galao",
		Kopisusu:     "Kopisusu",
		Affogato:     "Affogato",
		Macchiato:    "Macchiato",
		Bicerin:      "Bicerin",
		Bombón:       "Bombón",
		Mocha:        "Mocha",
		Caphesuada:   "Caphesuada",
		Carajillo:    "Carajillo",
		Espresso:     "Espresso",
		Eiskaffee:    "Eiskaffee",
		Frappuccino:  "Frappuccino",
		Iced:         "Iced",
		Indianfilter: "Indianfilter",
		Kopiluwak:    "Kopiluwak",
		Kopitubruk:   "Kopitubruk",
		Vienna:       "Vienna",
		Yuanyang:     "Yuanyang",
		None:         "None",
		Americano:    "Americano",
		Cubano:       "Cubano",
		Zorro:        "Zorro",
		Doppio:       "Doppio",
		Romano:       "Romano",
		Guillermo:    "Guillermo",
		Ristretto:    "Ristretto",
		Antoccino:    "Antoccino",
	}
)

func init() {
	var v InternalState
	if _, ok := interface{}(v).(fmt.Stringer); ok {
		_InternalStateNameToValue = map[string]InternalState{
			interface{}(Mochasippi).(fmt.Stringer).String():   Mochasippi,
			interface{}(Breve).(fmt.Stringer).String():        Breve,
			interface{}(Papi).(fmt.Stringer).String():         Papi,
			interface{}(Instant).(fmt.Stringer).String():      Instant,
			interface{}(Turkish).(fmt.Stringer).String():      Turkish,
			interface{}(Irish).(fmt.Stringer).String():        Irish,
			interface{}(Crema).(fmt.Stringer).String():        Crema,
			interface{}(Cappuccino).(fmt.Stringer).String():   Cappuccino,
			interface{}(Unknown).(fmt.Stringer).String():      Unknown,
			interface{}(Latte).(fmt.Stringer).String():        Latte,
			interface{}(Flatwhite).(fmt.Stringer).String():    Flatwhite,
			interface{}(Melange).(fmt.Stringer).String():      Melange,
			interface{}(Cortado).(fmt.Stringer).String():      Cortado,
			interface{}(Galao).(fmt.Stringer).String():        Galao,
			interface{}(Kopisusu).(fmt.Stringer).String():     Kopisusu,
			interface{}(Affogato).(fmt.Stringer).String():     Affogato,
			interface{}(Macchiato).(fmt.Stringer).String():    Macchiato,
			interface{}(Bicerin).(fmt.Stringer).String():      Bicerin,
			interface{}(Bombón).(fmt.Stringer).String():       Bombón,
			interface{}(Mocha).(fmt.Stringer).String():        Mocha,
			interface{}(Caphesuada).(fmt.Stringer).String():   Caphesuada,
			interface{}(Carajillo).(fmt.Stringer).String():    Carajillo,
			interface{}(Espresso).(fmt.Stringer).String():     Espresso,
			interface{}(Eiskaffee).(fmt.Stringer).String():    Eiskaffee,
			interface{}(Frappuccino).(fmt.Stringer).String():  Frappuccino,
			interface{}(Iced).(fmt.Stringer).String():         Iced,
			interface{}(Indianfilter).(fmt.Stringer).String(): Indianfilter,
			interface{}(Kopiluwak).(fmt.Stringer).String():    Kopiluwak,
			interface{}(Kopitubruk).(fmt.Stringer).String():   Kopitubruk,
			interface{}(Vienna).(fmt.Stringer).String():       Vienna,
			interface{}(Yuanyang).(fmt.Stringer).String():     Yuanyang,
			interface{}(None).(fmt.Stringer).String():         None,
			interface{}(Americano).(fmt.Stringer).String():    Americano,
			interface{}(Cubano).(fmt.Stringer).String():       Cubano,
			interface{}(Zorro).(fmt.Stringer).String():        Zorro,
			interface{}(Doppio).(fmt.Stringer).String():       Doppio,
			interface{}(Romano).(fmt.Stringer).String():       Romano,
			interface{}(Guillermo).(fmt.Stringer).String():    Guillermo,
			interface{}(Ristretto).(fmt.Stringer).String():    Ristretto,
			interface{}(Antoccino).(fmt.Stringer).String():    Antoccino,
		}
	}
}

// MarshalJSON is generated so InternalState satisfies json.Marshaler.
func (r InternalState) MarshalJSON() ([]byte, error) {
	if s, ok := interface{}(r).(fmt.Stringer); ok {
		return json.Marshal(s.String())
	}
	s, ok := _InternalStateValueToName[r]
	if !ok {
		return nil, fmt.Errorf("invalid InternalState: %d", r)
	}
	return json.Marshal(s)
}

// UnmarshalJSON is generated so InternalState satisfies json.Unmarshaler.
func (r *InternalState) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("InternalState should be a string, got %s", data)
	}
	v, ok := _InternalStateNameToValue[s]
	if !ok {
		return fmt.Errorf("invalid InternalState %q", s)
	}
	*r = v
	return nil
}