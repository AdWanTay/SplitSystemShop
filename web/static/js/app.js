function closeAllModals() {
    document.querySelectorAll('.modal').forEach(modal => modal.remove());
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

function openAboutModal(title, author, fullDescription, id, isAuth) {
    closeAllModals();

    fetch("/web/templates/modals/about-course.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);
            document.querySelector(".course-title").innerHTML = title;
            document.querySelector(".course-description").innerHTML = "<p>" + fullDescription + "</p>"
            document.querySelector(".course-category").innerHTML = "Автор курса: " + author;
            const buyBtn = modalContainer.querySelector(".buy-btn");
            if (isAuth){
                buyBtn.setAttribute("onclick", `openPurchaseModal("${title}", ${id})`);
            }else{
                buyBtn.setAttribute("onclick", `openAuthModal('login')`);

            }
            //todo ВЕЗДЕ СДЕЛАТЬ ТАК = Закрытие по клику вне модалки (доп)
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    modalContainer.remove();
                }
            });
        });
}

function openAuthModal(type) {
    closeAllModals();

    fetch("/web/templates/components/modals/auth-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            var input = document.querySelector("input[type='tel']");
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
                    modalContainer.remove();
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


// Функция для открытия модального окна с нужным содержимым
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
                    <button id="mainBtn" class="modal-button primary">${config.mainBtnText}</button>
                    <button class="modal-button cancel" onclick="this.closest('.modal').remove()">Отмена</button>
                </div>
            </div>
        </div>
    `;

    // Вставляем модальное окно в body
    document.body.insertAdjacentHTML('beforeend', modalHtml);

    // Назначаем обработчик для основной кнопки
    document.getElementById('mainBtn').addEventListener('click', config.mainBtnAction);

    try {
        if (config == modalConfigs.orderKit) {
            document.getElementById('fio').value = document.getElementById('full_name').value
            document.getElementById('email1').value = document.getElementById('email_address').value
            document.getElementById('tel1').value = document.getElementById('phone_number').value
        }
    } catch { }

    try {
        var input = document.querySelector("input[type='tel']");
        input.addEventListener("input", mask, false);
        input.addEventListener("focus", mask, false);
        input.addEventListener("blur", mask, false);
    } catch {
    }
}

// Конфигурации для разных модальных окон
const modalConfigs = {
    orderKit: {
        title: "Оставить заявку",
        body: `
            <form id="submitForm" autocomplete="off">
                <input id="fio" type="text" id="name" autocomplete="off" placeholder="ФИО" value="" required>
                <input id="email1" type="email" id="newEmail" autocomplete="off" placeholder="Почта" required>
                <input id="tel1" type="tel" id="email" autocomplete="off" placeholder="Телефон" required>
            </form>
        `,
        description: "ⓘ Необходимо ввести свои данные без ошибок",
        mainBtnText: "Готово",
        mainBtnAction: async function () {
            const email = document.getElementById('email1').value;
            const fullName = document.getElementById('fio').value;
            const phoneNumber = document.getElementById('tel1').value;

            if (!email || !fullName || !phoneNumber) {
                showErr("Все поля должны быть заполнены");
                return;
            }

            if (!validateEmail(email)) {
                showErr("Введен некорректный адрес электронной почты");
                return;
            }

            try {
                const response = await fetch('/api/starter-kit/request', {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    },
                    body: JSON.stringify({
                        email: email,
                        full_name: fullName,
                        phone_number: phoneNumber
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось");
                }

                showNotify("Успех", "Страница будет перезагружена");
                this.closest('.modal').remove();
                setTimeout(() => {
                    window.location.reload();
                }, 3100);
            } catch (error) {
                showErr(error.message);
            }
        }
    },
    emailEdit: {
        title: "Изменение почты",
        body: `
            <form id="emailForm" autocomplete="off">
                <input type="email" id="newEmail" autocomplete="off" placeholder="Новая почта" required>
            </form>
        `,
        description: "ⓘ Необходимо ввести новый адрес без ошибок",
        mainBtnText: "Сохранить",
        mainBtnAction: async function () {
            const newEmail = document.getElementById('newEmail').value;

            if (!newEmail) {
                showErr("Все поля должны быть заполнены");
                return;
            }

            if (!validateEmail(newEmail)) {
                showErr("Введен некорректный адрес электронной почты");
                return;
            }

            try {
                const response = await fetch('/api/auth/change-email', {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    },
                    body: JSON.stringify({
                        new_email: newEmail
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось изменить email");
                }

                showNotify("Успех", "Email успешно изменён. Страница будет перезагружена");
                this.closest('.modal').remove();
                setTimeout(() => {
                    window.location.reload();
                }, 3100);
            } catch (error) {
                showErr(error.message);
            }
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
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
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
    nameEdit: {
        title: "Изменение ФИО",
        body: `
            <form id="nameForm">
                <input type="text" id="newLastName" autocomplete="off" placeholder="Фамилия" required>
                <input type="text" id="newFirstName" autocomplete="off" placeholder="Имя" required>
                <input type="text" id="newPatronymic" autocomplete="off" placeholder="Отчество (необязательно)">
            </form>
        `,
        description: "ⓘ Укажите ваши реальные фамилию, имя и отчество",
        mainBtnText: "Сохранить",
        mainBtnAction: async function () {
            const lastName = document.getElementById('newLastName').value;
            const firstName = document.getElementById('newFirstName').value;
            const patronymic = document.getElementById('newPatronymic').value;

            if (!lastName || !firstName) {
                showErr("Фамилия и имя обязательны для заполнения");
                return;
            }

            try {
                const response = await fetch('/api/auth/change-name', {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    },
                    body: JSON.stringify({
                        new_last_name: lastName,
                        new_first_name: firstName,
                        new_patronymic: patronymic
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось изменить ФИО");
                }

                showNotify("Успех", "Персональные данные успешно изменены. Страница будет перезагружена");
                this.closest('.modal').remove();
                setTimeout(() => {
                    window.location.reload();
                }, 3100);
            } catch (error) {
                showErr(error.message);
            }
        }
    },
    phoneEdit: {
        title: "Изменение номера телефона",
        body: `
            <form id="phoneForm" autocomplete="off">
                <input type="tel" id="newPhoneNumber" autocomplete="off" placeholder="Новый номер телефона" required>
            </form>
        `,
        description: "ⓘ Необходимо ввести корректный номер телефона",
        mainBtnText: "Сохранить",
        mainBtnAction: async function () {
            const newPhoneNumber = document.getElementById('newPhoneNumber').value;

            if (!newPhoneNumber) {
                showErr("Все поля должны быть заполнены");
                return;
            }

            try {
                const response = await fetch('/api/auth/change-phone', {
                    method: 'PATCH',
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
                    },
                    body: JSON.stringify({
                        new_phone_number: newPhoneNumber
                    })
                });

                if (!response.ok) {
                    const error = await response.json();
                    throw new Error(error.message || "Не удалось изменить номер телефона");
                }
                showNotify("Успех", "Номер телефона успешно изменён. Страница будет перезагружена");
                this.closest('.modal').remove();
                setTimeout(() => {
                    window.location.reload();
                }, 3100);
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
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('token')}`
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


function openCalculator() {
    fetch("/web/templates/components/modals/calc-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            // реалиция

            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    modalContainer.remove();
                }
            });
            addAuthListener(type, modalContainer)
        });
}