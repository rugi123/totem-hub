document.addEventListener('DOMContentLoaded', async function (e) {
  e.preventDefault();
  console.log("profile");

  loadChatList();
});

async function loadChatList() {
  try {
    const response = await fetch('/api/chats', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
      }
    });

    const result = await response.json();

    if (!response.ok) {
      throw result
    }



  } catch (error) {
    console.log(error);
  }
}