document.getElementById("show-register").addEventListener("click", function (e) {
  e.preventDefault();
  document.getElementById("login-form").classList.remove("active");
  document.getElementById("register-form").classList.add("active");
});

document.getElementById("show-login").addEventListener("click", function (e) {
  e.preventDefault();
  document.getElementById("register-form").classList.remove("active");
  document.getElementById("login-form").classList.add("active");
});

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

  alert("Вхід успішний!");
  // Можна перенаправити користувача на іншу сторінку
  // window.location.href = "/dashboard.html";
});

// Реєстрація
document.getElementById("register-form").addEventListener("submit", async function (e) {
  e.preventDefault();

  const email = document.getElementById("register-email").value;
  const password = document.getElementById("register-password").value;
  const firstName = document.getElementById("register-first-name").value;
  const secondName = document.getElementById("register-second-name").value;

  const res = await fetch("/api/v1/auth/register", {
    method: "POST",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ email, password, firstName, secondName })
  });

  if (!res.ok) {
    alert("Помилка реєстрації.");
    return;
  }

  alert("Реєстрація успішна! Тепер увійдіть.");
  document.getElementById("register-form").classList.remove("active");
  document.getElementById("login-form").classList.add("active");
});
