<?php
if ($argc < 2) {
    fwrite(STDERR, "usage: ".$argv[0]." <URI>\n");
    exit(1); 
}

$conf = new RdKafka\Conf();
$conf->set('security.protocol', 'ssl');
$conf->set('ssl.ca.location', '/certificates/ca.certificate.pem');
$conf->set('ssl.certificate.location', '/certificates/access.certificate.pem');
$conf->set('ssl.key.location', '/certificates/access.key');
$conf->set('bootstrap.servers', $argv[1]);

$rk = new RdKafka\Consumer($conf);
if ($rk->addBrokers($argv[1]) < 1) {
    echo "Failed adding brokers\n";
    exit;
}

$metadata = $rk->getMetadata(true, NULL, 1000);
$topicsMetadata = $metadata->getTopics();

if ($topicsMetadata->count() == 0) {
    echo "No topics found on the cluster.\n";
    exit;
}

echo "Found the following topics on the cluster:\n";
while ($topicsMetadata->valid()) {
    $topicMetadata = $topicsMetadata->current();
    echo "- " . $topicMetadata->getTopic() . "\n";
    $topicsMetadata->next();
}
?>