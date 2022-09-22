<?php
require __DIR__ . '/vendor/autoload.php';

if ($argc < 2) {
    fwrite(STDERR, "usage: ".$argv[0]." <URI>\n");
    exit(1); 
}

$u = parse_url($argv[1]);

$client = (new \OpenSearch\ClientBuilder())
    ->setHosts(['https://'.$u['host'].':'.$u['port']])
    ->setBasicAuthentication($u['user'], $u['pass'])
    ->build();

$nodeInfo = $client->nodes()->info();
foreach($nodeInfo['nodes'] as $node) {
    echo $node['name'].': '.$node['version'].PHP_EOL;
}
?>