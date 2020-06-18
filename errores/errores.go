package errores

import (
	"encoding/json"
	"strings"

	"github.com/go-playground/validator"
)

var (
	PedidoIncorrecto     *Error
	RecursoNoEncontrado  *Error
	ErrorServidorInterno *Error
	ServicioNoDisponible *Error
)

// Error Estructura genérica para errores de Andreani S.A.
type Error struct {
	Type   string  `json:"type"`
	Title  string  `json:"title"`
	Detail string  `json:"detail"`
	Status int     `json:"status"`
	Fields []Field `json:"fields"`
}

// Field Exportable para desarrollos a medida
type Field struct {
	Name    string `json:"name"`
	Message string `json:"message"`
}

func init() {
	PedidoIncorrecto = &Error{Type: "about:blank", Title: "Error en la validacion de su pedido", Status: 400}
	RecursoNoEncontrado = &Error{Type: "about:blank", Title: "Recurso no encontrado", Status: 404}
	ErrorServidorInterno = &Error{Type: "about:blank", Title: "Error en la Respuesta", Status: 500}
	ServicioNoDisponible = &Error{Type: "about:blank", Title: "Servicio no disponible momentaneamete, intente nuevamente", Status: 503}
}

// Default - Permite settear solo el detalle y los errores del campo List
func (er *Error) Default(d string, e ...error) Error {
	er.Detail = d
	er.Fields = *errores2List(e)
	return *er
}

// All - Permite settear todos los valores del error
func (er *Error) All(t string, ti string, d string, s int, e ...error) Error {
	er.Type = t
	er.Title = ti
	er.Detail = d
	er.Status = s
	er.Fields = *errores2List(e)
	return *er
}

func errores2List(errs []error) *[]Field {
	var fieldList []Field

	for _, err := range errs {
		var field Field
		if ute, ok := err.(*json.UnmarshalTypeError); ok {
			field.Name = strings.ToLower(ute.Field)
			field.Message = ute.Value
			fieldList = append(fieldList, field)
		}

		if validatorErrors, ok := err.(validator.ValidationErrors); ok {
			for _, e := range validatorErrors {
				field.Name = strings.ToLower(e.Field())
				if e.Param() != "" {
					field.Message = e.Tag() + ": " + e.Param()
				} else {
					field.Message = e.Tag()
				}
				fieldList = append(fieldList, field)
			}
		}
	}

	return &fieldList

}
