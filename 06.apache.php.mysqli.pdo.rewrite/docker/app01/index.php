<?php
//SESSÃO
if(!isset($_SESSION)){
	session_start();
}
// _Ler o caminho
$request = $_SERVER["REQUEST_URI"];
// _Verifica se tem barra final
if  (substr($request,-1) == "/")
{
  // _Retirar a barra final
  $request = substr($request,0,-1);
}

switch ($request) {
    //___AVISOS_______________________________________________________
    case '/we' :
      require __DIR__ . '/controller/controller.we.php';
      break;

    default:
    require __DIR__ . '/phpinfo.php';
    break;
}

?>