document.getElementById('getPrice').addEventListener('click', function() {
    const speech = new SpeechSynthesisUtterance();
    speech.text = 'Уточнение';
    speech.lang = 'ru-RU'; // Установка русского языка озвучки
    window.speechSynthesis.speak(speech) 
    goStart();
});


function speak(message) {
    const speech = new SpeechSynthesisUtterance();
    speech.text = message;
    speech.lang = 'ru-RU'; // Установка русского языка озвучки
    window.speechSynthesis.speak(speech);
}

function goStart(){
	fetch('https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT')
        .then(response => response.json())
        .then(data => {
            const price = Math.round(data.price); // Округление до целого числа
            const message = `Текущая цена биткоина составляет ${price} долларов.`;
            speak(message);
        })
        .catch(error => console.error('Error fetching data:', error));
}

var noSleep = new NoSleep();

document.getElementById('activateNoSleep').addEventListener('click', function() {
    const speech = new SpeechSynthesisUtterance();
    speech.text = 'Запуск автоматического воспроизведения';
    speech.lang = 'ru-RU'; // Установка русского языка озвучки
    window.speechSynthesis.speak(speech) 
    noSleep.enable(); // Включить предотвращение спящего режима
	setInterval(goStart, 60000);
    this.remove(); // Удалить кнопку после активации
});