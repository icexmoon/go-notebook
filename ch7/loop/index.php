<?php
$name = isset($_GET['name']) ? $_GET['name'] : '';
$people = array(
    'Han Meimei' => array(
        'name' => 'Han Meimei',
        'age' => '20',
        'career' => 'student'
    ),
    'Brus Lee' => array(
        'name' => 'Brus Lee',
        'age' => '30',
        'career' => 'engineer'
    ),
    'Jack Chen' => array(
        'name' => 'Jack Chen',
        'age' => '50',
        'career' => 'actor'
    )
);
$result = array('status' => 'success', 'result' => array());
if (array_key_exists($name, $people)) {
    $result['result'] = $people[$name];
} else {
    $result['status'] = 'fail';
}
sleep(10);
echo json_encode($result);