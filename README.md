## `EN` **GO counter v.0.0.1 by m3**

### Task: 

The program is reads lines with URLs from stdin. For each URL need to send HTTP-request by GET-method and count an quantity of word "Go" in body response. In the end the program showes quantity of word "Go" from all requests.

### Solution:
 - Make an buffered channel "buf" on 5 variables.
 - Make an goroutine and send an url and channel.
 - Recieve all sums of word "go" from "buf" to "total" variable.
 
 File "urls" attached.
 
### Quick start:

```bash
$ cat urls | go run main.go
```

## `RU` **Счетчик GO вер.0.0.1 от m3**

### Задание:

Программа читает из stdin строки, содержащие URL. На каждый URL нужно отправить HTTP-запрос методом GET и посчитать кол-во вхождений строки "Go" в теле ответа. В конце работы приложение выводит на экран общее кол-во найденных строк "Go" во всех переданных URL.

### Решение задачи:

 - Создаем буфферизованный канал buf на 5 значений.
 - При считывании создаем горутину и отправляем url и канал buf
 - Общую сумму получаем в total из канала buf

Файл urls прилагается

### Запуск:

```bash
$ cat urls | go run main.go
```
