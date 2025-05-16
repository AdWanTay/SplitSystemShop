function closeAllModals() {
    document.querySelectorAll('.modal').forEach(modal => modal.remove());
    unlockBodyScroll();
}

function openAuthModal() {
    closeAllModals();
    lockBodyScroll();

    fetch("/web/templates/components/modals/auth-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            // const input = document.querySelector("input[type='tel']");
            // input.addEventListener("input", mask, false);
            // input.addEventListener("focus", mask, false);
            // input.addEventListener("blur", mask, false);
            // Специальный обработчик для автозаполнения
            // input.addEventListener('change', function (e) {
            //     if (this.value && !this.value.startsWith('+7')) {
            //         setTimeout(() => {
            //             const event = new Event('input');
            //             this.dispatchEvent(event);
            //         }, 100);
            //     }
            // }, false);

            document.getElementById('loginForm').addEventListener('submit', login(modalContainer));
            document.getElementById('registerForm').addEventListener('submit', register(modalContainer));

            showForm("login", modalContainer);

            //todo ВЕЗДЕ СДЕЛАТЬ ТАК = Закрытие по клику вне модалки (доп)
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    closeAllModals();
                }
            });
        });
}

function openArticleEditor() {
    closeAllModals();
    lockBodyScroll();

    fetch("/web/templates/components/modals/article-edit-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            initializeEditor();

        });
}

function openReviewModal(id) {
    closeAllModals();
    lockBodyScroll();

    fetch("/web/templates/components/modals/review-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);
            initReviewModal();

            document.getElementById('sendReviewBtn1').addEventListener('click', submitReview(id));

            //todo ВЕЗДЕ СДЕЛАТЬ ТАК = Закрытие по клику вне модалки (доп)
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    closeAllModals();
                }
            });
        });
}
function submitReview(splitSystemId) {
    return function () {
        const rating = parseInt(document.getElementById('rating-value').value);
        const comment = document.getElementById('review').value.trim();

        if (rating === 0) {
            showToast('', 'Пожалуйста, выберите рейтинг.');
            return;
        }

        if (comment === '') {
            showToast('', 'Пожалуйста, напишите отзыв.');
            return;
        }

        fetch('/api/review', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                split_system_id: splitSystemId,
                rating: rating,
                comment: comment
            })
        })
            .then(async response => {
                const data = await response.json();

                if (!response.ok) {
                    // Если сервер вернул ошибку
                    if (data.error) {
                        showErr(data.error);
                    } else {
                        showErr('Произошла ошибка при отправке отзыва.');
                    }
                    throw new Error(data.error || 'Ошибка при отправке');
                }

                // Успех
                showNotify('Успех', 'Отзыв отправлен успешно!');
                closeAllModals(); // если у тебя есть модальное окно
                // Очистка формы (если нужно)
                document.getElementById('review').value = '';
                document.getElementById('rating-value').value = 0;
            })
            .catch(error => {
                console.error('Ошибка при отправке отзыва:', error);
            });
    };
}



function showForm(type) {
    const btnLogin = document.getElementById('btn-login');
    const btnRegister = document.getElementById('btn-register');
    const loginForm = document.getElementById('loginForm');
    const registerForm = document.getElementById('registerForm');

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

function openCalculator() {
    lockBodyScroll();

    fetch("/web/templates/components/modals/calc-modal.html")
        .then((res) => res.text())
        .then((html) => {
            const modalContainer = document.createElement("div");
            modalContainer.innerHTML = html;
            document.body.appendChild(modalContainer);

            const form = document.querySelector('.cooling-calculator');
            const qValueEl = form.querySelector('.q-value');
            const qRangeEl = form.querySelector('.q-range');

            function calculate() {
                const area = parseFloat(form.area.value) || 0;
                const height = parseFloat(form.height.value) || 0;
                const people = parseInt(form.people.value) || 0;
                const computers = parseInt(form.computers.value) || 0;
                const tvs = parseInt(form.tvs.value) || 0;
                const otherPower = parseFloat(form.otherPower.value) || 0;

                const insolation = form.insolation.value;
                const ventilation = form.ventilation.checked;
                const airExchange = parseFloat(form.airExchange.value) || 0;
                const constantTemp = form.constantTemp.checked;
                const topFloor = form.topFloor.checked;
                const largeWindow = form.largeWindow.checked;
                const windowArea = parseFloat(form.windowArea.value) || 0;

                // Базовая мощность: объем помещения * коэффициент
                const volume = area * height;
                let q = volume * 0.04; // 40 Вт на м³

                // Корректировка по инсоляции
                if (insolation === 'high') q *= 1.2;
                else if (insolation === 'low') q *= 0.9;

                // Дополнительные источники тепла
                q += people * 0.1; // 100 Вт на человека
                q += computers * 0.3; // 300 Вт на компьютер
                q += tvs * 0.2; // 200 Вт на телевизор
                q += otherPower;

                // Вентиляция
                if (ventilation) {
                    q += volume * airExchange * 0.005; // 5 Вт на м³ при воздухообмене
                }

                // Дополнительные условия
                if (constantTemp) q *= 1.1;
                if (topFloor) q *= 1.05;
                if (largeWindow) q += windowArea * 0.1; // 100 Вт на м² окна

                q = Math.round(q * 100) / 100;

                // Рекомендуемый диапазон ±10%
                const qMin = Math.round(q * 0.9 * 100) / 100;
                const qMax = Math.round(q * 1.1 * 100) / 100;

                qValueEl.textContent = q.toFixed(2);
                qRangeEl.textContent = `${qMin.toFixed(2)} – ${qMax.toFixed(2)}`;
            }

            console.log("хуй")

            // Обновление при изменении значений
            form.addEventListener('input', function () {
                console.log("Вот")
                return calculate
            }());

            // Управление доступностью полей
            form.ventilation.addEventListener('change', () => {

                form.airExchange.disabled = !form.ventilation.checked;
                if (!form.ventilation.checked) form.airExchange.value = 1.0;
                calculate();
            });

            form.largeWindow.addEventListener('change', () => {
                form.windowArea.disabled = !form.largeWindow.checked;
                if (!form.largeWindow.checked) form.windowArea.value = 0;
                calculate();
            });

            // Начальный расчёт
            calculate();
            modalContainer.addEventListener("click", (e) => {
                if (e.target.classList.contains("modal")) {
                    closeAllModals();
                }
            });
        });
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
    console.log("register")
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
                // location.reload();
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
            location.href = '/#contact-us';
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
    },
    // articleDeleteConfirm: {
    //     title: "Подтверждение удаления",
    //     body: `
    //         <form id="deleteForm">
    //             <input type="text" id="deleteConfirmation" autocomplete="off" placeholder="Напишите 'УДАЛИТЬ'" required>
    //         </form>
    //     `,
    //     description: "ⓘ Это действие приведет к безвозвратному удалению статьи",
    //     mainBtnText: "Удалить статью",
    //     mainBtnAction: async function () {
    //         const confirmation = document.getElementById('deleteConfirmation').value;
    //
    //         if (confirmation !== 'УДАЛИТЬ') {
    //             showErr("Пожалуйста, введите пароль и напишите 'УДАЛИТЬ' для подтверждения");
    //             return;
    //         }
    //
    //         try {
    //             const response = await fetch('/api/article/${id}', {
    //                 method: 'DELETE',
    //                 headers: {
    //                     'Content-Type': 'application/json'
    //                 }
    //             });
    //
    //             if (!response.ok) {
    //                 const error = await response.json();
    //                 throw new Error(error.message || "Не удалось удалить статья");
    //             }
    //
    //             showNotify("Успех", "Статья была успешно удалена");
    //             this.closest('.modal').remove();
    //         } catch (error) {
    //             showErr(error.message);
    //         }
    //     }
    // },
};


function initReviewModal() {
    const stars = document.querySelectorAll('#rating .star');
    const ratingValueInput = document.getElementById('rating-value');
    let currentRating = 0;

    stars.forEach((star, index) => {
        star.addEventListener('mouseover', () => {
            updateStars(index + 1);
        });

        star.addEventListener('mouseout', () => {
            updateStars(currentRating);
        });

        star.addEventListener('click', () => {
            currentRating = index + 1;
            ratingValueInput.value = currentRating;
        });
    });

    function updateStars(rating) {
        stars.forEach((star, index) => {
            star.classList.toggle('active', index < rating);
        });
    }


}
