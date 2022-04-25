Тут пытаюсь сделать ТГ бота по [курсу](https://www.youtube.com/watch?v=73OsSlsuhFY&list=PLFAQFisfyqlWDwouVTUztKX2wUjYQ4T3l) 

## Ветка Keyboards

[Общее описание про клавиши](https://core.telegram.org/bots#keyboards)

Из описания получается, что когда бот отправляет сообщение и может отрисовать заранее подготовленную клавиатуру. 

Класивтура приезажает в параметре `reply_markup` в струкуре [sendMessage](https://core.telegram.org/bots/api#sendmessage)

может содержать несколько варианторв 
* [ReplyKeyboardMarkup](https://core.telegram.org/bots/api#replykeyboardmarkup) - клава с кнопками. Содержит массив массивов кнопок [KeyboardButton](https://core.telegram.org/bots/api#keyboardbutton) 

* [ReplyKeyboardRemove](https://core.telegram.org/bots/api#replykeyboardremove) - удаляет текущую клавиатуру и ставит вместо нее дефолтную. 

* [ForceReply](https://core.telegram.org/bots/api#forcereply) - судя по всему нужная штука когда надо пройти по цепочке вопросов ответов. Только пока не очень понял, где там сами кнопки. 

* [InlineKeyboardMarkup](https://core.telegram.org/bots/api#inlinekeyboardmarkup) - кнопки прям под сообщением. Содержит массив массивово кнопок [InlineKeyboardButton](https://core.telegram.org/bots/api#inlinekeyboardbutton). Выдают колбек, который надо еще дополнительно обрабатывать. 

Что получилось по кнопкам
1. С событием `/start` вылезает клавиатура. 
2. В обычных сообщениях не должна. 
3. Сама клавиатура реализована в `events/telegram/commands.go` в виде ужасного JSON. Это надо будет дико переделать


TODO:
1. реализовать нормальную клавиаутуру. просто засунуть в событие JSON очень не очень идея. 
2. покрыть тестиками все
2. переехать на БД
3. сделать разбор разных типов сообщений
4. добавить разлинчные команды.
5. сделать отдельную сборочную среду



