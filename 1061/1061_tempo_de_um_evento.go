package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Leitura das 4 linhas
	strDiaInicio, _ := reader.ReadString('\n')
	strHoraInicio, _ := reader.ReadString('\n')
	strDiaFim, _ := reader.ReadString('\n')
	strHoraFim, _ := reader.ReadString('\n')

	// Extraindo os dias
	diaInicioStr := strings.Split(strDiaInicio, " ")[1]
	diaFimStr := strings.Split(strDiaFim, " ")[1]
	diaInicio, _ := strconv.Atoi(strings.TrimSpace(diaInicioStr))
	diaFim, _ := strconv.Atoi(strings.TrimSpace(diaFimStr))

	// Removendo espaços e separando horas
	strHoraInicio = strings.ReplaceAll(strings.TrimSpace(strHoraInicio), " ", "")
	horaInicioSplit := strings.Split(strHoraInicio, ":")
	horaInicio, _ := strconv.Atoi(horaInicioSplit[0])
	minutoInicio, _ := strconv.Atoi(horaInicioSplit[1])
	segundoInicio, _ := strconv.Atoi(horaInicioSplit[2])

	strHoraFim = strings.ReplaceAll(strings.TrimSpace(strHoraFim), " ", "")
	horaFimSplit := strings.Split(strHoraFim, ":")
	horaFim, _ := strconv.Atoi(horaFimSplit[0])
	minutoFim, _ := strconv.Atoi(horaFimSplit[1])
	segundoFim, _ := strconv.Atoi(horaFimSplit[2])

	inicioTotal := diaInicio*86400 + horaInicio*3600 + minutoInicio*60 + segundoInicio
	fimTotal := diaFim*86400 + horaFim*3600 + minutoFim*60 + segundoFim

	duracao := fimTotal - inicioTotal

	dias := duracao / 86400
	duracao %= 86400
	horas := duracao / 3600
	duracao %= 3600
	minutos := duracao / 60
	segundos := duracao % 60

	fmt.Printf("%d dia(s)\n", dias)
	fmt.Printf("%d hora(s)\n", horas)
	fmt.Printf("%d minuto(s)\n", minutos)
	fmt.Printf("%d segundo(s)\n", segundos)

}
