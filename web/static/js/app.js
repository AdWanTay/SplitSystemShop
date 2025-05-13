function closeAllModals() {
    document.querySelectorAll('.modal').forEach(modal => modal.remove());
    unlockBodyScroll();
}

function showForm(type) {
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');
    const btnLogin = document.getElementById('btn-login');
    const btnRegister = document.getElementById('btn-register');

    if (type === 'login') {
        loginForm.style.display = 'block';
        registerForm.style.display = 'none';
        btnLogin.checked = true;
    } else {
        loginForm.style.display = 'none';
        registerForm.style.display = 'block';
        btnRegister.checked = true;
    }
}


document.addEventListener("DOMContentLoaded", () => {
    try {
        const loginBtn = document.getElementById("loginBtn");
        loginBtn.addEventListener("click", () => openAuthModal("login"));

        const registerBtn = document.getElementById("registerBtn");
        registerBtn.addEventListener("click", () => openAuthModal("register"));
    } catch {
    }

    try {
        const logoutBtn = document.querySelector(".logoutBtn");
        logoutBtn.addEventListener("click", () => logout());
    } catch {
    }

    try {
        const input = document.querySelector("input[type='tel']");
        input.addEventListener("input", mask, false);
        input.addEventListener("focus", mask, false);
        input.addEventListener("blur", mask, false);
    } catch {
    }
});


function openPurchaseModal(title, id) {
    closeAllModals();

    fetch("/web/templates/modals/purchase.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);
            if (title != undefined)
                document.querySelector('.purchase-modal .course-title').innerHTML = "«" + title + "»";

            setTimeout(async () => {
                const response = await fetch(`/api/course/buy/${id}`, {
                    method: 'POST',
                    headers: {'Content-Type': 'application/json'},
                });

                const result = await response.json();

                if (response.ok) {
                    showNotify("Успех", "Оплата прошла успешно! Страница будет перезагружена");
                    document.querySelector('.qr-code-container').style.display = 'none';
                    document.querySelector('.purchase-description').style.display = 'none';
                    document.querySelector('.payment-link-container').style.display = 'none';
                    document.querySelector('.success-text').style.height = 'auto';
                    document.querySelector('.success-text').style.opacity = 1;

                    const animation = lottie.loadAnimation({
                        container: document.getElementById('success'),
                        renderer: 'svg',
                        path: '/web/static/js/success.json'
                    });
                    animation.play();
                } else {
                    showErr("Ошибка", "Оплата не прошла. Страница будет перезагружена");
                }
            }, 6100);
            setTimeout(() => {
                window.location.reload();
            }, 13500);
        });
}

function openCartModal() {
    lockBodyScroll();

    fetch("/web/templates/components/modals/cart-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            //todo ВЕЗДЕ СДЕЛАТЬ ТАК = Закрытие по клику вне модалки (доп)
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    closeAllModals();
                }
            });
        });
}

function openCalculator() {
    lockBodyScroll();

    fetch("/web/templates/components/modals/calc-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            // реалиция

            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    closeAllModals();
                }
            });
            addAuthListener(type, modalContainer)
        });
}

function openAuthModal(type) {
    closeAllModals();

    lockBodyScroll();

    fetch("/web/templates/components/modals/auth-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            const input = document.querySelector("input[type='tel']");
            input.addEventListener("input", mask, false);
            input.addEventListener("focus", mask, false);
            input.addEventListener("blur", mask, false);

            // Специальный обработчик для автозаполнения
            input.addEventListener('change', function(e) {
                if (this.value && !this.value.startsWith('+7')) {
                    setTimeout(() => {
                        const event = new Event('input');
                        this.dispatchEvent(event);
                    }, 100);
                }
            }, false);

            // Запускаем нужную форму
            if (type === "login") {
                showForm("login");
            } else {
                showForm("register");
            }

            //todo ВЕЗДЕ СДЕЛАТЬ ТАК = Закрытие по клику вне модалки (доп)
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    closeAllModals();
                }
            });
            addAuthListener(type, modalContainer)
        });
}

function addAuthListener(type, modalContainer) {
    if (type === "login") {
        const loginForm = document.getElementById("loginForm");
        const loginHandler = login(modalContainer);
        loginForm.addEventListener('submit', loginHandler);
    } else {
        const registerForm = document.getElementById("registerForm");
        const registerHandler = register(modalContainer);
        registerForm.addEventListener('submit', registerHandler);
    }
}

function login(modalContainer) {
    return async (event) => {
        event.preventDefault();
        const form = event.target;
        const formData = {
            email: form.email.value,
            password: form.password.value,
        };

        try {
            const response = await fetch('/api/auth/login', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(formData)
            });
            const result = await response.json();

            if (response.ok) {
                modalContainer.remove();
                location.reload();
            } else {
                showErr(result.error)
            }
        } catch (error) {
            showErr('Ошибка при отправке запроса')
        }
    }
}

function register(modalContainer) {
    return async (event) => {
        event.preventDefault();
        const form = event.target;
        const formData = {
            lastName: form.lastName.value,
            firstName: form.firstName.value,
            patronymic: form.patronymic.value,
            phoneNumber: form.phoneNumber.value,
            email: form.email.value,
            password: form.password.value,
        };

        try {
            const response = await fetch('/api/auth/register', {
                method: 'POST',
                headers: {'Content-Type': 'application/json'},
                body: JSON.stringify(formData)
            });

            const result = await response.json();

            if (response.ok) {
                modalContainer.remove()
                location.reload();
            } else {
                showErr('Ошибка входа: ' + (result.error || 'Неизвестная ошибка'))
            }
        } catch (error) {
            console.error('Ошибка:', error);
            showErr('Ошибка при отправке запроса')
        }
    }
}

async function logout() {
    try {
        const response = await fetch("/api/auth/logout", {
            method: "GET",
            credentials: "include", // Обязательно, чтобы cookie ушла
        });

        if (response.ok) {
            window.location.href = "/";
        } else {
            showErr("Ошибка в запросе выхода. Попробуйте позже")
        }
    } catch (err) {
        showErr("Ошибка при запросе выхода:", err);
    }
}

function openModal(config) {
    closeAllModals();

    const modalHtml = `
        <div class="modal">
            <div class="modal-content main-modal">
                <span class="modal-close" onclick="this.closest('.modal').remove()">✖</span>
                <h2 class="modal-title">${config.title}</h2>
                <div class="modal-body">${config.body}</div>
                <div class="modal-description">${config.description}</div>
                <div class="modal-actions">
                    <button id="mainBtn" class="accent-btn btn-standard">${config.mainBtnText}</button>
                    <button class="primary-btn btn-standard" onclick="this.closest('.modal').remove()">Отмена</button>
                </div>
            </div>
        </div>
    `;

    document.body.insertAdjacentHTML('beforeend', modalHtml);
    document.getElementById('mainBtn').addEventListener('click', config.mainBtnAction);
}

const modalConfigs = {
    contactUs: {
        title: "Свяжитесь с нами, или оставьте заявку",
        body: `
            <div class="contact__content">
                <div class="contact-card">
                    <h3 class="contact-card-title">Контакты</h3>
                    <a href="tel:89998002001">+7 (999) 800-20-01</a>
                    <a href="mailto:info@climatehome.online">info@climatehome.online</a>
                    <a href="mailto:admin@climatehome.online">admin@climatehome.online</a>
                </div>
                <div class="contact-card">
                    <h3 class="contact-card-title">Адрес</h3>
                    <p>Респ. Башкортостан, г. Уфа</p>
                </div>
            </div>
        `,
        description: "ⓘ Рабочие дни: Пн-Пт с 10:00 до 19:00, Сб-Вс Выходной",
        mainBtnText: "Оставить заявку",
        mainBtnAction: async function () {
            this.closest('.modal').remove();
            location.href='/#contact-us';
        }
    },
    passwordEdit: {
        title: "Изменение пароля",
        body: `
            <form id="passwordForm">
                <input type="password" id="newPassword" autocomplete="off" placeholder="Новый пароль" required>
                <input type="password" id="repeatNewPassword" autocomplete="off" placeholder="Повторите новый пароль" required>
            </form>
        `,
        description: "ⓘ Пароль должен содержать не менее 8 символов",
        mainBtnText: "Сохранить",
        mainBtnAction: async function () {
            const newPassword = document.getElementById('newPassword').value;
            const repeatNewPassword = document.getElementById('repeatNewPassword').value;

            if (!newPassword || !repeatNewPassword) {
                showErr("Все поля должны быть заполнены");
                return;
            }

            if (newPassword !== repeatNewPassword) {
                showErr("Пароли не совпадают");
                return;
            }

            if (newPassword.length < 8) {
                showErr("Пароль должен содержать не менее 8 символов");
                return;
            }

            try {
                const response = await fetch('/api/auth/change-password', {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        new_password: newPassword
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось изменить пароль");
                }
                showNotify("Успех", "Пароль успешно изменён");
                this.closest('.modal').remove();
            } catch (error) {
                showErr(error.message);
            }
        }
    },
    deleteAccount: {
        title: "Удаление аккаунта",
        body: `
            <form id="deleteForm">
                <input type="text" id="deleteConfirmation" autocomplete="off" placeholder="Напишите 'УДАЛИТЬ'" required>
            </form>
        `,
        description: "ⓘ Это действие необратимо. Все ваши данные будут удалены.",
        mainBtnText: "Удалить аккаунт",
        mainBtnAction: async function () {
            const confirmation = document.getElementById('deleteConfirmation').value;

            if (confirmation !== 'УДАЛИТЬ') {
                showErr("Пожалуйста, введите пароль и напишите 'УДАЛИТЬ' для подтверждения");
                return;
            }

            try {
                const response = await fetch('/api/auth/delete-account', {
                    method: 'DELETE',
                    headers: {
                        'Content-Type': 'application/json'
                    }
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось удалить аккаунт");
                }

                showNotify("Успех", "Аккаунт успешно удалён");
                showNotify("Внимание!", "Аккаунт будет удален в течение");
                this.closest('.modal').remove();

                setTimeout(() => {
                    window.location.href = '/';
                }, 3100);
            } catch (error) {
                showErr(error.message);
            }
        }
    }
};
