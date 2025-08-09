document.getElementById('register-form').addEventListener('submit', async function (e) {
    e.preventDefault();

    const formData = new FormData(this);

    const formObject = {};
    formData.forEach((value, key) => {
        formObject[key] = value;
    });

    const jsonData = JSON.stringify(formObject);

    try {
        // Отправляем данные на сервер
        const response = await fetch('/auth/register', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: jsonData
        });

        if (!response.ok) {
            const error = await response.json();
            throw error
        }

        window.location.href = "/auth/login";
    } catch (error) {
        const errorMessage = error.error.replace(/;/g, '');

        const parts = errorMessage
            .split('error in field ')
            .slice(1);
        const result = parts.reduce((obj, item) => {
            const [key, value] = item.split(":");
            obj[key] = value;
            return obj;
        }, {});
        console.log(result);

        const errPanel = document.getElementById("error-panel")

        errPanel.innerHTML = '';

        if (result['Password'] === 'min') {
            errPanel.innerHTML += '<snap class="err-item">пароль слишком короткий</snap>';
        }
        if (result['Name'] === 'min') {
            errPanel.innerHTML += '<snap class="err-item">никнейм слишком короткий</snap>';
        }
    }

});