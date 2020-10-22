### Usage


composer.json中增加：

```json
"autoload":{
        "psr-4":{
          "GPBMetadata\\":"GPBMetadata/",
          "Services\\":"Services/"
        }
   }

composer dumpautoload

```



```bash
cd pbfiles/ && protoc --php_out=../ Product.proto && cd ../



php Demo/Demo.php
```