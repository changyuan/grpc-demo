### Usage

`grpc_php_plugin` 文件为编译的的插件，用来生成PHP端的文件。

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
