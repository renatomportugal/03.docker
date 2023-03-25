# Dicas Liinux

## Quantidade de Arquivos e Pastas

```CMD
Quantidade de Arquivos
ls -la|grep -e "^-"|wc -l

Quantidade de Pastas
ls -la|grep -e "^d"|wc -l

Fazer um comando
alias fid='echo "Arquivos: $(ls -la |grep "^-"|wc -l)  Diretórios: $(ls -la|grep "^d"|wc -l)"'

fid

```
