document.addEventListener("DOMContentLoaded", () => {
  const greetingEl = document.getElementById("greeting");
  const housesListEl = document.getElementById("houses-list");

  const token = localStorage.getItem("token");

  // ✅ Отримуємо firstName через /api/v1/users/me
  fetch("/api/v1/users", {
    headers: {
      "Authorization": "Bearer " + token
    }
  })
    .then(res => {
      if (!res.ok) {
        throw new Error("Не вдалося отримати дані користувача");
      }
      return res.json();
    })
    .then(user => {
      greetingEl.textContent = `Вітаємо, ${user.firstName}!`;
    })
    .catch(err => {
      greetingEl.textContent = `Вітаємо, користувачу!`;
      console.error(err);
    });

  // 🔘 Створюємо кнопку додавання приміщення
  const createBtn = document.createElement("button");
  createBtn.textContent = "Створити нове приміщення";
  createBtn.style.marginBottom = "20px";
  createBtn.onclick = () => {
    // TODO: додати функціонал пізніше
    alert("Функція в розробці");
  };
  housesListEl.before(createBtn); // Кнопка над списком

  // 📦 Отримуємо список будинків
  fetch("/api/v1/houses", {
    headers: {
      "Authorization": "Bearer " + token
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
