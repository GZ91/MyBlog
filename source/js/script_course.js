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
            const message = `${price}`;
            speak(message);
        })
        .catch(error => console.error('Error fetching data:', error));
}

var noSleep = new NoSleep();
let idSetInterval = null;

document.getElementById('activateNoSleep').addEventListener('click', function() {
    if (idSetInterval == null) {
        const speech = new SpeechSynthesisUtterance();
        speech.text = 'Автомат запущен';
        speech.lang = 'ru-RU'; // Установка русского языка озвучки
        window.speechSynthesis.speak(speech) 
        noSleep.enable(); // Включить предотвращение спящего режима
        idSetInterval = setInterval(goStart, 60000);
        //this.remove(); // Удалить кнопку после активации
        this.style.backgroundColor = 'red'
    } else {
        clearInterval(idSetInterval);
        const speech = new SpeechSynthesisUtterance();
        speech.text = 'Автомат остановлен';
        speech.lang = 'ru-RU'; // Установка русского языка озвучки
        window.speechSynthesis.speak(speech) 
        intervalID = null
        this.style.backgroundColor = ''
    }      
});

let checkInterval = null; // Для хранения ссылки на setInterval
let buttonCheckPrice = null;
let userPrice = null;

document.getElementById('checkPriceBtn').addEventListener('click', function() {
    buttonCheckPrice = this;
    if (checkInterval == null) {
        // Если проверка не активирована, начинаем проверку
        checkPrice(); // Проверяем цену сразу при нажатии
        checkInterval = setInterval(checkPrice, 10000); // Затем устанавливаем интервал

        // Изменяем цвет кнопки, указывая на активацию
        buttonCheckPrice.style.backgroundColor = 'red';
        userPrice = document.getElementById('priceInput').value;
        buttonCheckPrice.textContent = `Остановить ожидание (${userPrice})`;
        
        speak(`Ожидание на цену (${userPrice}) включено`)
        
    } else {
        userPrice = null;
        speak('Ожидание остановлено')
        stopChecking()
    }
});

function stopChecking(){
    clearInterval(checkInterval);
    checkInterval = null; // Сбрасываем ссылку на setInterval

    // Возвращаем цвет кнопки к исходному состоянию
    buttonCheckPrice.style.backgroundColor = ''; // Исходный цвет кнопки
    buttonCheckPrice.textContent = 'Ожидать ценовое превышение';
}

function checkPrice() {

    fetch('https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT')
        .then(response => response.json())
        .then(data => {
            const btcPrice = parseFloat(data.price);
            if (btcPrice > userPrice) {
                let btcPriceVoice = Math.round(data.price)
                speak(`Текущая цена Bitcoin (${btcPriceVoice}) выше указанной (${userPrice})!`);
                stopChecking();
            }
        })
        .catch(error => {
            console.error('Ошибка при получении данных: ', error);
            alert('Произошла ошибка при получении данных с Binance.');
        });
}


//MIN
let checkIntervalMin = null; // Для хранения ссылки на setInterval
let buttonCheckPriceMin = null;
let userPriceMin = null;

document.getElementById('checkMinPriceBtn').addEventListener('click', function() {
    buttonCheckPriceMin = this;
    if (checkIntervalMin == null) {
        // Если проверка не активирована, начинаем проверку
        checkPriceMin(); // Проверяем цену сразу при нажатии
        checkIntervalMin = setInterval(checkPriceMin, 10000); // Затем устанавливаем интервал

        // Изменяем цвет кнопки, указывая на активацию
        buttonCheckPriceMin.style.backgroundColor = 'red';
        userPriceMin = document.getElementById('priceInputMin').value;
        buttonCheckPriceMin.textContent = `Остановить ожидание (${userPriceMin})`;
        
        speak(`Ожидание на цену (${userPriceMin}) включено`)
        
    } else {
        userPriceMin = null;
        speak('Ожидание остановлено')
        stopCheckingMin()
    }
});

function stopCheckingMin(){
    clearInterval(checkIntervalMin);
    checkIntervalMin = null; // Сбрасываем ссылку на setInterval

    // Возвращаем цвет кнопки к исходному состоянию
    buttonCheckPriceMin.style.backgroundColor = ''; // Исходный цвет кнопки
    buttonCheckPriceMin.textContent = 'Ожидать ценовое понижение';
}

function checkPriceMin() {

    fetch('https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT')
        .then(response => response.json())
        .then(data => {
            const btcPrice = parseFloat(data.price);
            if (btcPrice < userPriceMin) {
                let btcPriceVoice = Math.round(data.price)
                speak(`Текущая цена Bitcoin (${btcPriceVoice}) ниже указанной (${userPriceMin})!`);
                stopCheckingMin();
            }
        })
        .catch(error => {
            console.error('Ошибка при получении данных: ', error);
            alert('Произошла ошибка при получении данных с Binance.');
        });
}