<?php
$pdo = null;
try {
    $pdo = new PDO('mysql:host=mysql.5.7;dbname=tecnocrata_org', 'root', 'my-password');
} catch (PDOException $e) {
    print $e->getMessage();
    die();
}
var_dump($pdo);