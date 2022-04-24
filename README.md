Тут пытаюсь сделать ТГ бота по [курсу](https://www.youtube.com/watch?v=73OsSlsuhFY&list=PLFAQFisfyqlWDwouVTUztKX2wUjYQ4T3l) 

Сам бот [тут](https://github.com/JustSkiv/read-adviser-bot/tree/lessons)


Пока во получается довольно печально.
1. собираем докер с прогой. 
2.  внутри контейнера собраем бот и потом его запускаем

это не очень себе механизм. 
надо научиться делать отделную сборочную среду где бот собирается и вторую где бот уже будет работать. 

## Ветка Keyboards
Тут буду пытаться добавить кнопки. 
Нужны для дргого бота. Но пытаться будут тут. 

[Общее описание про клавиши](https://core.telegram.org/bots#keyboards)

Из описания получается, что когда бот отправляет сообщение и может отрисовать заранее подготовленную клавиатуру. 

Класивтура приезажает в параметре `reply_markup` в струкуре [sendMessage](https://core.telegram.org/bots/api#sendmessage)

может содержать несколько варианторв 
* [ReplyKeyboardMarkup](https://core.telegram.org/bots/api#replykeyboardmarkup) - клава с кнопками. Содержит массив массивов кнопок [KeyboardButton](https://core.telegram.org/bots/api#keyboardbutton) 

* [ReplyKeyboardRemove](https://core.telegram.org/bots/api#replykeyboardremove) - удаляет текущую клавиатуру и ставит вместо нее дефолтную. ***Пока не очень понимаю назначение***

* [ForceReply](https://core.telegram.org/bots/api#forcereply) - судя по всему нужная штука когда надо пройти по цепочке вопросов ответов. Только пока не очень понял, где там сами кнопки. 

* [InlineKeyboardMarkup](https://core.telegram.org/bots/api#inlinekeyboardmarkup) - кнопки прям под сообщением. Содержит массив массивово кнопок [InlineKeyboardButton](https://core.telegram.org/bots/api#inlinekeyboardbutton)


### План.
1. Создаем [ReplyKeyboardMarkup](https://core.telegram.org/bots/api#replykeyboardmarkup) с парой кнопок.
2. Втыкаем эту штуку в сообщение. 

вот собсна и план. 

Описание типов будет тут `clients/telegram/types.go`
Тут надо немного переделать `clients/telegram/telegram.go` стурктуру `SendMessage` что бы можно было в нее заливать клавиатуру. 

1. Создаем интерфес `reply_markup`. По идее это какая-то офигенская структура или интрефейс, которая может быть использована для всеми, описанными выше варианатами ответов. 
2. В функцию `SendMessage` добавляем поле `reply_markup`. 
3. Проверяем текущий функционал. 
