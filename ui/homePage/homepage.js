document.addEventListener("DOMContentLoaded", async () => {
  const token = localStorage.getItem("token");

  if (!token) {
    window.location.href = "/auth/login";
    return;
  }

  try {
    const res = await fetch("/api/v1/users", {
      headers: {
        "Authorization": "Bearer " + token
      }
    });

    if (!res.ok) {
      throw new Error("Неможливо завантажити дані користувача.");
    }

    const user = await res.json();
    document.getElementById("user-info").textContent =
      `Ви увійшли як ${user.firstName} ${user.secondName} (${user.email})`;
  } catch (err) {
    alert("Помилка: " + err.message);
    logout();
  }
});

function logout() {
  localStorage.removeItem("token");
  window.location.href = "/auth/login";
}
