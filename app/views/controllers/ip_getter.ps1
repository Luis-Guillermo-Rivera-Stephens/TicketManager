$response = curl https://ipinfo.io/ip 
$text = "const URL_SERVER = 'http://$($response):8080';"

Set-Content -Path "./app/views/controllers/URL.js" -Value $text

