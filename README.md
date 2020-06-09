# questack-api

## resources 

### Question 

|name|type|description|
|:---|:---|:---|
|id|int|primary key|
|title|string||
|author|string|questioner's name|
|content|string|question body|


### Owner

|name|type|description|
|:---|:---|:---|
|id|int|primary key|
|name|string|user's name|
|mail|string|mail address|
|description|string|presentation|
|stacks|[]Stack|question boxes|



### Stack

|name|type|description|
|:---:|:---:|:---|
|id|int|primary key|
|owner|Owner||
|name|string||
|questions|[]Question||


