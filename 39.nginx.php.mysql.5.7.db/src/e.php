<?php  

$servidor = "mysql.5.7";
$db_usuario = "root";
$db_senha = "my-password";
$banco = "tecnocrata_org";

$mysqli = @mysqli_connect($servidor,$db_usuario,$db_senha,$banco);

  echo "<pre>";
  if ($mysqli) {
    var_dump($mysqli);
  } else {
    // echo $this->erroMysqli." - ".$this->erroMysqli_Mensagem = $this->cod_MySQL_Message [$this->erroMysqli];
    echo "Não deu certo";
  }
  echo "</pre>";

  

?>