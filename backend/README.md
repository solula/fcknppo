# Бэкенд проекта Waterfall

### Запуск в Docker
```shell
docker-compose -f docker-compose.yml -p waterfall-backend up -d --build
```

### Первоначальное описание основных сущностей

Первый этап:
```go
type Chapter struct {
	Uuid            string
	Number          int
	Title           string
	PartUuid        string
	TextFile        string
	ImageFiles      []string
	MainImageIdx    *int
	FramingSvgFiles []string
	BottomSvgFile   *string
}

type Part struct {
	Uuid           string
	Number         int
	Title          string
	CoverImageFile string
	Annotation     *string
}

type File struct {
	Uuid     string
	FileName string
	Size     int64
	Bucket   string
}
```

Второй этап:
```go
type Comment struct {
	Uuid       string
	Text       string
	OwnerUuid  *string
	OwnerName  *string
	ParentUuid string
}

type User struct {
	Uuid        string
	Name        string
	Roles       []string
	Permissions []string
}
```

Третий этап:
```go
type Notification struct {
	Uuid       string
	Title      string
	Message    string
	Importance string
}

type NotificationRecipient struct {
	Uuid             string
	NotificationUuid string
	RecipientUuid    string
	ViewedAt         *time.Time
}
```

Четверный этап:
```go
type ChatMessage struct {
	Uuid           string
	Text           string
	ViewedAt       *time.Time
	SenderUuid     string
	RecipientUuid  string
}
```



---

Хочу посоветоваться с вами насчет хранения файлов.

Допустим, есть сущность "пост". Доступ к посту определяется типом пользователя (есть те, которые доступны, например, только админу или платным пользователям). К посту может быть прикреплено несколько фотографий, которые хранятся в s3 (сами данные поста хранятся в postgres, также в postgres хранятся метаданные файлов). 

Вопросы:

1. Каким образом должно происходить создание, удаление и обновление поста? Особенно интересует обновление: должны ли ставшие ненужными файлы во время обновления безвозвратно удаляться?
2. Есть бизнес-логика, определяющая доступ к постам. В случае, если пост доступен, то доступны и файлы. Если пост недоступен, то его файлы для пользователя также должны быть недоступны. Как это реализовать?
3. Про транзакции

Мое решение:

Я храню файлы с указанием того, какой сущности они принадлежат (например, посту). При необходимости прикрепить файл к сущности происходит следующее:

1. Файл загружается на сервер во временный бакет (s3).
2. Когда создается пост с указанием ссылки на этот файл, этот файл переносится в бакет поста и в базе данных проставляется ссылка на этот пост. Таким образом, чтобы узнать права доступа к файлам, нужно посмотреть на права доступа к сущности, к которой они прикреплены (доступна ли сущность или нет).

При обновлении происходит сходная операция, но есть также шаг с удалением ненужных файлов. Все файлы, uuid'ы которых не указаны при обновлении сущности, удаляются. 

---

```go
// File файл
type File struct {
	Uuid        string            // Uuid
	Filename    string            // Имя файла
	MIMEType    string            // MIME тип файла
	Description string            // Описание
	CreatorUuid *string           // Uuid создателя файла
	ObjectRef   *models.ObjectRef // Ссылка на сущность, которой файл принадлежит (на dto.ChapterAttachment)
	Temp        bool              // Признак временного файла
	Type        *file_type.Type   // Тип файла (картинка, аватар, аттачмент, ...)
}

// Chapter глава
type Chapter struct {
Uuid        string  // Uuid
Number      int     // Номер
Title       string  // Название
PartUuid    string  // Часть, которой глава принадлежит
ReleaseUuid *string // Uuid релиза
}

type ChapterText struct {
Uuid        string // Uuid
ChapterUuid string // Uuid глав
Version     int    // Версия
Text        string // Текст
}

type ChapterAttachment struct {
Uuid        string // Uuid
RowNum      int    // Номер строки
ChapterUuid string // Uuid главы
}
```