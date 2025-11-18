package main

import (
	"task-api/internal/handlers"
	"task-api/internal/repositories"
	"task-api/internal/usecases"
)

/*
Sistema de Agenda

- Armazenar para cada compromisso:
	1 - Data/Horário
	2 - Titulo do compromisso
	3 - Descrição do compromisso (opcional)
	4 - Pessoas Envolvidas

	- As pessoas devem ser armazenadas individualmente

-Funcionalidades:
	1 - Listar Compromissos
	2 - Cadastrar Compromisso
	4 - Alterar titulo e descriçao em um compromisso
	5 - Exlcuir uma pessoa em um compromisso
	6 - Excluir um comprmisso
*/

func main() {
	repos := repositories.New()
	usecases := usecases.New(repos)
	handler := handlers.New(usecases)
	handler.Listen(8080)
}
