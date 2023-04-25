package q5

import "strings"

//Pedro começou a frequentar aulas de programação. Na primeira aula, sua tarefa foi escrever um programa simples. O
//programa deveria fazer o seguinte: na sequência de caracteres fornecida, composta por letras latinas maiúsculas e
//minúsculas, ele:
//
//* deleta todas as vogais;
//* insere um caractere "." antes de cada consoante;
//* substitui todas as consoantes maiúsculas pelas correspondentes em minúsculas.
//
//As vogais são as letras "A", "O", "E", "U", "I", e o restante são consoantes. A entrada do programa é exatamente uma
//sequência de caracteres e a saída deve ser uma única sequência de caracteres, resultante após o processamento do
//programa na sequência de caracteres inicial.
//
//Ajude Pedro a lidar com esta tarefa fácil.

func ProcessString(s string) string {

	palavra_corrigida := ""

	palavra_corrigida = strings.ReplaceAll(s, "A", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "E", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "I", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "O", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "U", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "a", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "e", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "i", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "o", "")
	palavra_corrigida = strings.ReplaceAll(palavra_corrigida, "u", "")

	var nova_palavra string
	for i := 0; i < len(palavra_corrigida); i++ {
		nova_palavra += "."
		nova_palavra += string(palavra_corrigida[i])
	}
	palavra_minuscula := strings.ToLower(nova_palavra)

	return palavra_minuscula
}
