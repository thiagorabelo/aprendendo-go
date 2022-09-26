// go mod init <mod_name>
module modulo

go 1.19

// git get github.com/badoux/checkmail
// go mod tidy // limpar pacotes não usados
require github.com/badoux/checkmail v1.2.1 // indirect
