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


let ArrayMax = new Array();

document.getElementById('addMechanismMax').addEventListener('click', function() {
    var container = document.getElementById('containerMax');
  
    // Создаем поле ввода
    var input = document.createElement('input');
    input.type = 'number';
    input.placeholder = 'Укажите цену биткоина';
  
    // Создаем кнопку для запуска процедуры
    var button = document.createElement('button');
    button.textContent = 'Запустить ожидание';
    
    // Добавляем обработчик нажатия для кнопки
    button.addEventListener('click', function() {
      var targetPrice = input.value;
      let thisElem = null;
      for (let i = 0; i < ArrayMax.length; i++) {
        if (this == ArrayMax[i].button){
            if (ArrayMax[i].checkInterval == null){
                // Если проверка не активирована, начинаем проверку
                const checkPriceFunc = function(){
                    fetch('https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT')
                    .then(response => response.json())
                    .then(data => {
                        const btcPrice = parseFloat(data.price);
                        if (btcPrice > ArrayMax[i].userPrice) {
                            let btcPriceVoice = Math.round(data.price)
                            speak(`Текущая цена Bitcoin (${btcPriceVoice}) выше (${ArrayMax[i].userPrice})!`);
                            ArrayMax[i].userPrice = null;
                            clearInterval(ArrayMax[i].checkInterval);
                            ArrayMax[i].checkInterval = null;
                            // Возвращаем цвет кнопки к исходному состоянию
                            ArrayMax[i].button.style.backgroundColor = ''; // Исходный цвет кнопки
                            ArrayMax[i].button.textContent = 'Ожидать ценовое превышение';
                        }
                    })
                    .catch(error => {
                        console.error('Ошибка при получении данных: ', error);
                        alert('Произошла ошибка при получении данных с Binance.');
                    });
                }; // Проверяем цену сразу при нажатии
                checkPriceFunc();
                ArrayMax[i].checkInterval = setInterval(checkPriceFunc, 10000); // Затем устанавливаем интервал
        
                // Изменяем цвет кнопки, указывая на активацию
                ArrayMax[i].button.style.backgroundColor = 'red';
                ArrayMax[i].userPrice = ArrayMax[i].input.value;
                ArrayMax[i].button.textContent = `Остановить ожидание MAX(${ArrayMax[i].userPrice})`;
                
                speak(`Ожидание на цену (${ArrayMax[i].userPrice}) включено`)
             }else{
                ArrayMax[i].userPrice = null;
                speak('Ожидание остановлено')
                clearInterval(ArrayMax[i].checkInterval);
                ArrayMax[i].checkInterval = null;
                // Возвращаем цвет кнопки к исходному состоянию
                ArrayMax[i].button.style.backgroundColor = ''; // Исходный цвет кнопки
                ArrayMax[i].button.textContent = 'Ожидать ценовое превышение';
             }
            break
        }
      }
      
      
    });

    let objMax = {
        input: input,
        button: button,
        checkInterval: null,
        userPrice: 0,
    }  
    ArrayMax.push(objMax)
    // Добавляем элементы в DOM
    container.appendChild(input);
    container.appendChild(button);
  });


let ArrayMin = new Array();

document.getElementById('addMechanismMin').addEventListener('click', function() {
    var container = document.getElementById('containerMin');
  
    // Создаем поле ввода
    var input = document.createElement('input');
    input.type = 'number';
    input.placeholder = 'Укажите цену биткоина';
  
    // Создаем кнопку для запуска процедуры
    var button = document.createElement('button');
    button.textContent = 'Запустить ожидание';
    
    // Добавляем обработчик нажатия для кнопки
    button.addEventListener('click', function() {
      var targetPrice = input.value;
      for (let i = 0; i < ArrayMin.length; i++) {
        if (this == ArrayMin[i].button){
            if (ArrayMin[i].checkInterval == null){
                // Если проверка не активирована, начинаем проверку
                const checkPriceFunc = function(){
                    fetch('https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT')
                    .then(response => response.json())
                    .then(data => {
                        const btcPrice = parseFloat(data.price);
                        if (btcPrice < ArrayMin[i].userPrice) {
                            let btcPriceVoice = Math.round(data.price)
                            speak(`Текущая цена Bitcoin (${btcPriceVoice}) ниже (${ArrayMin[i].userPrice})!`);
                            ArrayMin[i].userPrice = null;
                            clearInterval(ArrayMin[i].checkInterval);
                            ArrayMin[i].checkInterval = null;
                            // Возвращаем цвет кнопки к исходному состоянию
                            ArrayMin[i].button.style.backgroundColor = ''; // Исходный цвет кнопки
                            ArrayMin[i].button.textContent = 'Ожидать ценовое понижение';
                        }
                    })
                    .catch(error => {
                        console.error('Ошибка при получении данных: ', error);
                        alert('Произошла ошибка при получении данных с Binance.');
                    });
                }; // Проверяем цену сразу при нажатии
                checkPriceFunc();
                ArrayMin[i].checkInterval = setInterval(checkPriceFunc, 10000); // Затем устанавливаем интервал
        
                // Изменяем цвет кнопки, указывая на активацию
                ArrayMin[i].button.style.backgroundColor = 'red';
                ArrayMin[i].userPrice = ArrayMin[i].input.value;
                ArrayMin[i].button.textContent = `Остановить ожидание MIN(${ArrayMin[i].userPrice})`;
                
                speak(`Ожидание на цену (${ArrayMin[i].userPrice}) включено`)
             }else{
                ArrayMin[i].userPrice = null;
                speak('Ожидание остановлено')
                clearInterval(ArrayMin[i].checkInterval);
                ArrayMin[i].checkInterval = null;
                // Возвращаем цвет кнопки к исходному состоянию
                ArrayMin[i].button.style.backgroundColor = ''; // Исходный цвет кнопки
                ArrayMin[i].button.textContent = 'Ожидать ценовое понижение';
             }
            break
        }
      }
      
      
    });

    let objMax = {
        input: input,
        button: button,
        checkInterval: null,
        userPrice: 0,
    }  
    ArrayMin.push(objMax)
    // Добавляем элементы в DOM
    container.appendChild(input);
    container.appendChild(button);
  });