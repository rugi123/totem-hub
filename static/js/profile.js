document.addEventListener('DOMContentLoaded', function () {
    const chatItems = document.querySelectorAll('.chat-item');
    const chatContent = document.querySelector('.chat-content');

    loadChatList();

});

async function loadChatList() {
    try {
        // Показываем индикатор загрузки
        chatList.innerHTML = '<div class="loading">Загрузка чатов...</div>';

        // Запрашиваем список чатов с сервера
        const response = await fetch('/api/chats');
        const chats = await response.json();

        // Очищаем список чатов
        chatList.innerHTML = '';

        // Добавляем чаты в список
        chats.forEach(chat => {
            const chatItem = document.createElement('div');
            chatItem.className = 'chat-item';
            chatItem.dataset.chatId = chat.id;
            chatItem.innerHTML = `
          <div class="chat-avatar" style="background-color: ${getRandomColor()}">
            ${getInitials(chat.name)}
          </div>
          <div class="chat-info">
            <div class="chat-name">${chat.name}</div>
            <div class="chat-preview">${chat.lastMessage || 'Нет сообщений'}</div>
          </div>
        `;
            chatList.appendChild(chatItem);

            // Добавляем обработчик клика
            chatItem.addEventListener('click', function () {
                selectChat(chat.id);
            });
        });

        // Проверяем URL - если есть ID чата, загружаем его
        const pathParts = window.location.pathname.split('/');
        if (pathParts.length > 2 && pathParts[1] === 'chat') {
            const chatId = pathParts[2];
            selectChat(chatId);
        }

    } catch (error) {
        console.error('Ошибка загрузки списка чатов:', error);
        chatList.innerHTML = '<div class="error">Не удалось загрузить чаты</div>';
    }
}

async function selectChat(chatId) {
    // Удаляем класс active у всех чатов
    document.querySelectorAll('.chat-item').forEach(item => {
        item.classList.remove('active');
    });

    // Добавляем класс active к выбранному чату
    const selectedChat = document.querySelector(`.chat-item[data-chat-id="${chatId}"]`);
    if (selectedChat) {
        selectedChat.classList.add('active');
    }

    // Загружаем содержимое чата
    await loadChatContent(chatId);

    // Обновляем URL
    history.pushState({ chatId }, '', `/chat/${chatId}`);
}

// Функция загрузки содержимого чата
async function loadChatContent(chatId) {
    try {
        // Показываем индикатор загрузки
        chatContent.innerHTML = '<div class="loading">Загрузка чата...</div>';

        // Запрашиваем данные чата с сервера
        const response = await fetch(`/api/chats/${chatId}`);
        const chatData = await response.json();

        // Отображаем содержимое чата
        chatContent.innerHTML = `
        <div class="chat-header">
          <h3>${chatData.name}</h3>
          <div class="chat-members">Участников: ${chatData.membersCount}</div>
        </div>
        <div class="chat-messages">
          ${chatData.messages.map(msg => `
            <div class="message ${msg.isMy ? 'my-message' : ''}">
              <div class="message-sender">${msg.sender}:</div>
              <div class="message-text">${msg.text}</div>
              <div class="message-time">${formatTime(msg.time)}</div>
            </div>
          `).join('')}
        </div>
        <div class="chat-input">
          <input type="text" placeholder="Введите сообщение...">
          <button>Отправить</button>
        </div>
      `;

    } catch (error) {
        console.error('Ошибка загрузки чата:', error);
        chatContent.innerHTML = '<div class="error">Не удалось загрузить чат</div>';
    }
}

// Обработка кнопки "Назад" в браузере
window.addEventListener('popstate', function (event) {
    if (event.state && event.state.chatId) {
        selectChat(event.state.chatId);
    } else {
        chatContent.innerHTML = '<div class="chat-empty">Выберите чат для начала общения</div>';
        document.querySelectorAll('.chat-item').forEach(item => {
            item.classList.remove('active');
        });
    }
});