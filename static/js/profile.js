document.addEventListener('DOMContentLoaded', function () {
  console.log("Profile loaded");

  // Инициализация приложения
  initChatApp();
});

let currentChatId = null;
let socket = null;

async function initChatApp() {
  try {
    await loadChatList();
    initWebSocket();
  } catch (error) {
    console.error('Initialization error:', error);
    // Здесь можно добавить уведомление для пользователя
  }
}

async function loadChatList() {
  const container = document.getElementById('chat-container');
  if (!container) {
    console.error('Chat container not found');
    return;
  }

  container.innerHTML = ''; // Очистка перед загрузкой

  try {
    const response = await fetch('/api/chats', {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    console.log('Chat list data:', data);

    // Используем фрагмент для оптимизации рендеринга
    const fragment = document.createDocumentFragment();

    data.chats.forEach(chat => {
      const chatItem = document.createElement('div');
      chatItem.className = 'chat-item';
      chatItem.dataset.id = chat.ID; // Исправлено: chat.ID вместо chatItem.id

      // Создаем структуру чата
      chatItem.innerHTML = `
        <div class="chat-avatar"></div>
        <div class="chat-info">
          <div class="chat-name">${escapeHTML(chat.Title)}</div>
        </div>
      `;

      chatItem.addEventListener('click', () => {
        document.querySelectorAll('.chat-item').forEach(el => {
          el.classList.remove('active');
        });
        chatItem.classList.add('active');
        currentChatId = chat.ID;
        loadChatMessages(chat.ID);
      });

      fragment.appendChild(chatItem);
    });

    container.appendChild(fragment);

    // Автовыбор первого чата
    if (data.chats.length > 0) {
      const firstChat = container.querySelector('.chat-item');
      firstChat?.click();
    }

  } catch (error) {
    console.error('Failed to load chat list:', error);
    // Здесь можно добавить уведомление для пользователя
  }
}

async function loadChatMessages(chatID) {
  const container = document.getElementById('chat-content');
  if (!container) return;

  container.innerHTML = ''; // Очищаем контейнер

  try {
    const response = await fetch(`/api/chats/${chatID}/messages`, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' },
    });

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    const data = await response.json();
    console.log('Chat messages data:', data);

    // Используем фрагмент для оптимизации
    const fragment = document.createDocumentFragment();

    data.messages.forEach(message => {
      const messageItem = document.createElement('div');
      messageItem.className = 'message';
      messageItem.dataset.id = message.ID; // Исправлено: message.ID

      const date = new Date(message.SentAt);
      messageItem.innerHTML = `
        <div class="message-text">${escapeHTML(message.Text)}</div>
        <div class="message-time">
          ${date.getHours()}:${date.getMinutes().toString().padStart(2, '0')}
        </div>
      `;

      fragment.appendChild(messageItem);
    });

    container.appendChild(fragment);
    // Прокрутка к последнему сообщению
    container.scrollTop = container.scrollHeight;

  } catch (error) {
    console.error(`Failed to load messages for chat ${chatID}:`, error);
  }
}

function initWebSocket() {
  socket = new WebSocket('ws://localhost:8080/ws');

  socket.onopen = () => {
    console.log('WebSocket connection established');
  };

  socket.onmessage = (event) => {
    console.log('Received message:', event.data);
    try {
      const message = JSON.parse(event.data);
      // Добавляем сообщение, если оно для текущего чата
      if (message.chatId === currentChatId) {
        addMessageToChat(message);
      }
    } catch (e) {
      console.error('Error parsing WebSocket message:', e);
    }
  };

  socket.onerror = (error) => {
    console.error('WebSocket error:', error);
  };

  socket.onclose = (event) => {
    console.log(event.wasClean ? 'Connection closed cleanly' : 'Connection interrupted');
    // Автопереподключение
    setTimeout(initWebSocket, 3000);
  };

  // Инициализация обработчика формы
  document.getElementById('message-form').addEventListener('submit', async function (event) {
    event.preventDefault();

    const input = this.querySelector('.message-input');
    const message = input.value.trim();

    if (!message) return;

    try {
      // Отправка сообщения через WebSocket или fetch
      if (socket && socket.readyState === WebSocket.OPEN) {
        socket.send(JSON.stringify({
          text: message,
          chatId: currentChatId
        }));
        input.value = ''; // Очищаем поле ввода
      } else {
        console.error('WebSocket не подключен');
      }
    } catch (error) {
      console.error('Ошибка отправки сообщения:', error);
    }
  });
}

function handleMessageSubmit(event) {
  event.preventDefault();

  if (!currentChatId) {
    alert('Please select a chat first');
    return;
  }

  const formData = new FormData(event.target);
  const message = {
    text: formData.get('text'),
    chatId: currentChatId,
    timestamp: new Date().toISOString()
  };

  // Отправка через WebSocket
  if (socket && socket.readyState === WebSocket.OPEN) {
    socket.send(JSON.stringify(message));
    // Добавляем сообщение локально
    addMessageToChat({
      ...message,
      ID: Date.now() // Временный ID
    });
    event.target.reset(); // Очистка формы
  } else {
    console.error('WebSocket is not connected');
  }
}

function addMessageToChat(message) {
  const container = document.getElementById('chat-content');
  if (!container) return;

  const messageItem = document.createElement('div');
  messageItem.className = 'message';
  messageItem.dataset.id = message.ID;

  const date = new Date(message.timestamp || Date.now());
  messageItem.innerHTML = `
    <div class="message-text">${escapeHTML(message.text)}</div>
    <div class="message-time">
      ${date.getHours()}:${date.getMinutes().toString().padStart(2, '0')}
    </div>
  `;

  container.appendChild(messageItem);
  // Прокрутка к новому сообщению
  container.scrollTop = container.scrollHeight;
}

// Вспомогательная функция для безопасности
function escapeHTML(str) {
  return str.replace(/[&<>"']/g,
    match => ({
      '&': '&amp;',
      '<': '&lt;',
      '>': '&gt;',
      '"': '&quot;',
      "'": '&#39;'
    }[match]));
}