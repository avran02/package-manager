# Simple ssh package manager

### Установка
``` bash
git clone https://github.com/avran02/package-manager
cd package-manager
go install

# Создаём .env конфиг в домашней директории пользователя
package-manager init

# Поднимаем тестовый ssh сервер в контейнере 
docker-compose up -d
```
Конфиг уже содержит данные для подключения к docker контейнеру с поднятым ssh сервером 

# ТЗ

Сделать на GO пакетный менеджер

Должен уметь упаковывать файлы в архив, и заливать их на сервер по SSH
должен уметь скачивать файлы архивов по SSH и распаковывать.

Файл для упаковки должен иметь формат .yaml или json
в файле должны быть указаны пути по которым нужно подобрать файлы по маске

Пример файла пакета для упаковки:

packet.json
```
{
 "name": "packet-1",
 "ver": "1.10",
 "targets": [
  "./archive_this1/*.txt",
  {"path", "./archive_this2/*", "exclude": "*.tmp"},
 ]
 packets: {
  {"name": "packet-3", "ver": "<="2.0" },
 }
}
```
Пример файла для распаковки:


packages.json
```
{
 "packages": [
  {"name": "packet-1", "ver": ">=1.10"},
  {"name": "packet-2" },
  {"name": "packet-3", "ver": "<="1.10" },
 ]
}
```
Сделать commandline tools с командами:
```
pm create ./packet.json

pm update ./packages.json
```
PS: Можно использовать любые допущения которые сделают разработку тестового задания проще