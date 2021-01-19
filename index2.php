<?php

$index=0;

while (true) {



    echo "success------------".$index.PHP_EOL;

    sleep(mt_rand(1,3));

    $index++;

    if($index>=10){

       throw new Exception("gg");
    }

//    throw new Exception("gg");

//   break;

}


