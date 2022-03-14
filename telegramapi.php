<?php
$messaggio=$_GET["m"];
$data= $messaggio;
function sendMessage($chatID, $messaggio, $token) {

    $url = "https://api.telegram.org/bot" . $token . "/sendMessage?chat_id=" . $chatID;
    $url = $url . "&text=" . urlencode($messaggio);
    $ch = curl_init();
    $optArray = array(
            CURLOPT_URL => $url,
            CURLOPT_RETURNTRANSFER => true
    );
    curl_setopt_array($ch, $optArray);
    $result = curl_exec($ch);
    curl_close($ch);
    return $result;
    echo $result;
}

 sendMessage("##YOUR_TELEGRAM_USER_ID##",$data,"##YOUR_BOT_API##");

?>