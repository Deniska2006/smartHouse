document.addEventListener("DOMContentLoaded", () => {
  const loginForm = document.getElementById("login-form");
  const registerForm = document.getElementById("register-form");
  const showRegisterLink = document.getElementById("show-register");
  const showLoginLink = document.getElementById("show-login");

  // Перемикання на форму реєстрації
  showRegisterLink.addEventListener("click", (e) => {
    e.preventDefault();
    loginForm.classList.remove("active");
    registerForm.classList.add("active");
  });

  // Перемикання на форму входу
  showLoginLink.addEventListener("click", (e) => {
    e.preventDefault();
    registerForm.classList.remove("active");
    loginForm.classList.add("active");
  });

  // Обробка входу
  loginForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const email = document.getElementById("login-email").value;
    const password = document.getElementById("login-password").value;

    try {
      const res = await fetch("/api/v1/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email, password }),
      });

      if (!res.ok) {
        alert("Помилка входу. Перевірте дані.");
        return;
      }

      const data = await res.json();
      localStorage.setItem("token", data.token);
      localStorage.setItem("firstName", data.firstName);
      window.location.href = "/homepage";
    } catch (err) {
      alert("Помилка мережі.");
      console.error(err);
    }
  });

  // Обробка реєстрації
  registerForm.addEventListener("submit", async (e) => {
    e.preventDefault();

    const firstName = document.getElementById("register-first-name").value;
    const secondName = document.getElementById("register-second-name").value;
    const email = document.getElementById("register-email").value;
    const password = document.getElementById("register-password").value;

    try {
      const res = await fetch("/api/v1/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ firstName, secondName, email, password }),
      });

      if (!res.ok) {
        alert("Помилка реєстрації. Перевірте дані.");
        return;
      }

      alert("Реєстрація пройшла успішно! Тепер увійдіть.");
      // Після успішної реєстрації показати форму входу
      registerForm.classList.remove("active");
      loginForm.classList.add("active");
    } catch (err) {
      alert("Помилка мережі.");
      console.error(err);
    }
  });
});
