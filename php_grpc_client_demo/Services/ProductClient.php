<?php
namespace Services;

use Services\ProductRequest;

class ProductClient extends \Grpc\BaseStub{

    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    public function GetProdStock(ProductRequest $argument,$metadata=[],$options=[]) {

        return $this->_simpleRequest('/services.ProductService/GetProdStock', $argument,['\Services\ProductResponse', 'decode'],$metadata, $options);
    }
}