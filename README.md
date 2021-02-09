# Кодогенератор для паттерна options

Иногда бывает утомительно описывать огромное количество опциональных параметров,
но по факту это не очень сложная кодогенерация.

## Пример

``` bash
./options-codegen --input testdata --struct Simple
```

**options.go**

``` golang
package options

type Simple struct {
    //options:ignore
    StringVal string
    IntVal    int
}
```

**simple_options.go**

``` golang
// DO NOT EDIT!!!

package options

type Option func(options *Simple)

func OptionIntVal(option int) Option {
    return func(options *Simple) {
        options.IntVal = option
    }
}
```
