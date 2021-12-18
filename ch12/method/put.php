<?php
$http_method = $_SERVER['REQUEST_METHOD'];
switch ($http_method) {
    case 'PUT':
        do_put();
        break;
    default:
}

function do_put(): void
{
    $contents = file_get_contents("php://input");
    $contents = json_decode($contents, true);
    $response = [
        "new_data" => $contents['data']
    ];
    echo json_encode($response);
}
