i don't really know if anyone checks these tasks and if they even relevant to getting internship, but for those who is not me checking this repository:    
 
- В папке basic_task1 находится задание с github.com/beevik/ntp.  
- В папке calendar_api находится задание №11, где нужно поднять http сервер, задача мне показалась самой интересной из всех.  Применил подход чистой архитекуры и разделил приложение на 3 слоя:  
*handler*, *service* и *repository*, про базу данных ничего сказано не было, поэтому для хранения я использовал мапу, которую и разместил в папке repository  
Остальные задачи находятся в папке tasks/cmd. Для unix утилит я написал простенький makefile и тесты на bash, которые сверяют системные bash утилиты с имплементированными мной  
Ну надо сказать, что на языке программирования *С* я уже писал свои cat и grep с поддержкой почти всех флагов, so i did those tasks not very enthusiastically  
