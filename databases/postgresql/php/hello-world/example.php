<?php
if ($argc < 2) {
    fwrite(STDERR, "URI is missing in the environment.\n");
    exit(1); 
}

$values = parse_url($argv[1]);
$host = $values['host'];
$user = $values['user'];
$password = $values['pass'];
$dbname = preg_replace('/^\//', '', $values['path']);
$port = $values['port'];

$conn = new PDO("pgsql:host=$host; port=$port; dbname=$dbname; sslmode=require; user=$user; password=$password");
$data = $conn->query("select version();")->fetchColumn();
echo $data, "\n";
?>