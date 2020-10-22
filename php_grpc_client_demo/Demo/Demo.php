<?php
require 'vendor/autoload.php';

//用于连接 服务端
$client = new \Services\ProductClient('127.0.0.1:8081', [
    'credentials' => Grpc\ChannelCredentials::createInsecure()
]);

//实例化 TestRequest 请求类
$request = new \Services\ProductRequest();
$request->setProdId(22);

//调用远程服务
$get = $client->GetProdStock($request)->wait();

list($res, $status) = $get;

$stock = $res->getProdStock();

var_dump($stock);exit;