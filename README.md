# string-unpacker

Otus golang home work #2

## Задание

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)
* "qwe\4\5" => "qwe45" (*)
* "qwe\45" => "qwe44444" (*)
* "qwe\\5" => "qwe\\\\\" (*)
  
## Usage example

```shell
export GO111MODULE=on
git clone git@github.com:imorph/string-unpacker.git
cd string-unpacker
go build
./string-unpacker
```