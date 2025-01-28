package service

import (
	"fmt"
	"syncal/internal/utils/utilRepeatType"
	"time"
)

type dataBody struct {
	title       string
	name        string
	last        string
	location    string
	start       time.Time
	end         time.Time
	endGroup    time.Time
	description string
	typeOf      utilRepeatType.RepeatType
	to          string
}

func (d *dataBody) buildMessage() string {
	if d.typeOf == utilRepeatType.Monthly {
		return d.createComplexMsg("mensual")
	} else if d.typeOf == utilRepeatType.Weekly {
		return d.createComplexMsg("semanal")
	} else if d.typeOf == utilRepeatType.Daily {
		return d.createComplexMsg("diaria")
	} else {
		return d.createSimpleMsg()
	}
}

func (d *dataBody) createSimpleMsg() string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f9f9f9; }
			.container { max-width: 600px; margin: 20px auto; background: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1); }
			.header { text-align: center; }
			.title { font-size: 24px; font-weight: bold; color: #333333; }
			.content { margin-top: 20px; font-size: 16px; line-height: 1.6; color: #555555; }
			.actions { margin-top: 30px; text-align: center; }
			.button { display: inline-block; margin: 0 10px; padding: 10px 20px; font-size: 16px; color: #ffffff; text-decoration: none; border-radius: 4px; }
			.accept { background-color: #4caf50; }
			.reject { background-color: #f44336; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<p class="title">%s</p>
			</div>
			<div class="content">
				<p>Fuiste invitado por <strong>%s %s</strong> a <strong>%s</strong>, una %s.</p>
				<p>Horario: <strong>%s</strong> a <strong>%s</strong></p>
				<p>%s</p>
			</div>
			<div class="actions">
				<a href="localhost:8080/api/participant/accept?utilEmail=%s" class="button accept">Aceptar</a>
				<a href="localhost:8080/api/participant/reject?utilEmail=%s" class="button reject">Rechazar</a>
			</div>
		</div>
	</body>
	</html>
	`, d.title, d.name, d.last, d.title, d.location, d.start, d.end, d.description,
		d.to, d.to)
}

func (d *dataBody) createComplexMsg(recurrence string) string {
	return fmt.Sprintf(`
	<!DOCTYPE html>
	<html>
	<head>
		<style>
			body { font-family: Arial, sans-serif; margin: 0; padding: 0; background-color: #f9f9f9; }
			.container { max-width: 600px; margin: 20px auto; background: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0px 4px 6px rgba(0, 0, 0, 0.1); }
			.header { text-align: center; }
			.title { font-size: 24px; font-weight: bold; color: #333333; }
			.content { margin-top: 20px; font-size: 16px; line-height: 1.6; color: #555555; }
			.actions { margin-top: 30px; text-align: center; }
			.button { display: inline-block; margin: 0 10px; padding: 10px 20px; font-size: 16px; color: #ffffff; text-decoration: none; border-radius: 4px; }
			.accept { background-color: #4caf50; }
			.reject { background-color: #f44336; }
		</style>
	</head>
	<body>
		<div class="container">
			<div class="header">
				<p class="title">%s</p>
			</div>
			<div class="content">
				<p>Fuiste invitado por <strong>%s %s</strong> a <strong>%s</strong>, una reunión <strong>%s</strong>.</p>
				<p>Comienza el <strong>%s</strong>, termina el <strong>%s</strong>.</p>
				<p>Horario: <strong>%s</strong> a <strong>%s</strong></p>
				<p>%s</p>
			</div>
			<div class="actions">
				<a href="localhost:8080/api/participant/accept?utilEmail=%s" class="button accept">Aceptar</a>
				<a href="localhost:8080/api/participant/reject?utilEmail=%s" class="button reject">Rechazar</a>
			</div>
		</div>
	</body>
	</html>
	`, d.title, d.name, d.last, d.title, recurrence, d.start.Format("02/01/2006"), d.endGroup.Format("02/01/2006"), d.start.Format("15:04"), d.end.Format("15:04"), d.description, d.to, d.to)
}
