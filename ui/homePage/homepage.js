document.addEventListener("DOMContentLoaded", () => {
  const greetingEl = document.getElementById("greeting");
  const housesListEl = document.getElementById("houses-list");

  // Відображення імені користувача (якщо є)
  const userName = localStorage.getItem("userName") || "користувачу";
  greetingEl.textContent = `Вітаємо, ${userName}!`;

  // Запит приміщень з авторизацією
  fetch("/api/v1/houses", {
    headers: {
      "Authorization": "Bearer " + localStorage.getItem("token")
    }
  })
  .then(res => {
    if (!res.ok) {
      throw new Error("Не вдалося отримати список приміщень");
    }
    return res.json();
  })
  .then(houses => {
    if (houses.length === 0) {
      housesListEl.innerHTML = "<p>Список приміщень порожній.</p>";
      return;
    }

    housesListEl.innerHTML = ""; // Очистити перед додаванням

    houses.forEach(house => {
      const card = document.createElement("div");
      card.className = "house-card";
      card.innerHTML = `
        <h3>${house.name}</h3>
        <p><strong>Місто:</strong> ${house.city}</p>
        <p><strong>Адреса:</strong> ${house.address}</p>
        <p><strong>Опис:</strong> ${house.description}</p>
      `;
      housesListEl.appendChild(card);
    });
  })
  .catch(err => {
    housesListEl.innerHTML = `<p style="color: red;">${err.message}</p>`;
    console.error(err);
  });
});
