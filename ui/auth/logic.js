// Вхід
document.getElementById("login-form").addEventListener("submit", async function (e) {
  e.preventDefault();

  const email = document.getElementById("login-email").value;
  const password = document.getElementById("login-password").value;

  const res = await fetch("/api/v1/auth/login", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ email, password })
  });

  if (!res.ok) {
    alert("Помилка входу. Перевірте дані.");
    return;
  }

  const data = await res.json();
  localStorage.setItem("token", data.token); // Зберігаємо токен у браузері

  // ✅ Зберігаємо ім’я користувача для привітання на homepage
  localStorage.setItem("userName", data.firstName || "користувач");

  // ✅ Перенаправлення на головну сторінку
  window.location.href = "/homepage";
});
