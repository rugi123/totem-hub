document.getElementById('login-form').addEventListener('submit', async function (e) {
    e.preventDefault();

    const formData = new FormData(this);

    const formObject = {};
    formData.forEach((value, key) => {
        formObject[key] = value;
    });

    const jsonData = JSON.stringify(formObject);
    console.log(jsonData)

    try {
        // Отправляем данные на сервер
        const response = await fetch('/auth/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: jsonData
        });

        if (!response.ok) {
            throw new Error('Ошибка сети');
        }

        const result = await response.json();
        console.log('Успешно:', result);
        window.location.href = "/profile";
    } catch (error) {
        console.error('Ошибка:', error);
        // Обработка ошибок
    }

});