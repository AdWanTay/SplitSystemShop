// admin.js - скрипты для административной панели

document.addEventListener('DOMContentLoaded', function () {
    const uploadInput = document.getElementById('upload-btn');
    const imagePreview = document.getElementById('image-preview');

    uploadInput.addEventListener('change', function () {
        const file = uploadInput.files[0];
        if (file) {
            const reader = new FileReader();
            reader.onload = function (e) {
                imagePreview.src = e.target.result;
            };
            reader.readAsDataURL(file);
        }
    });

    const defaultTabId = "tab-btn-1";
    const savedTabId = localStorage.getItem("activeTabId");
    const tabIdToActivate = savedTabId || defaultTabId;

    const tabInput = document.getElementById(tabIdToActivate);
    if (tabInput) {
        tabInput.checked = true;
    }

    document.querySelectorAll('input[name="tab-btn"]').forEach(input => {
        input.addEventListener("change", () => {
            if (input.checked) {
                localStorage.setItem("activeTabId", input.id);
            }
        });
    });
});

const tableBody = document.querySelector('#products-table tbody');
const searchBtn = document.getElementById('search-btn');
const searchInput = document.getElementById('search-input');
const addBtn = document.getElementById('add-btn');
const deleteBtn = document.getElementById('delete-btn');
const totalCount = document.getElementById('total-count');
const selectAllCheckbox = document.getElementById('select-all');
const form = document.getElementById("product-form");

let hasUnsavedChanges = false;
let addingNewProduct = true;

let selectedId = 0;
let allProducts = [];
const visibleRows = 5;

document.addEventListener('DOMContentLoaded', function () {
    document.getElementById('openChatBtn').classList.add('hidden');
    fetchOrders();

    form.addEventListener("input", () => {
        hasUnsavedChanges = true;
    });

    document.getElementById("product-form").addEventListener("submit", async function (e) {
        e.preventDefault();
        if (addingNewProduct) {
            await createProduct()
        } else { // обновление товара
            await updateProduct()
        }
    });

    searchInput.addEventListener('input', function () {
        const searchTerm = searchInput.value.trim().toLowerCase();
        if (searchTerm.length !== 0) {
            filterProducts(searchTerm);
        }
    });

    // Обработчик "Выбрать все"
    selectAllCheckbox.addEventListener('change', function () {
        const checkboxes = tableBody.querySelectorAll('input[type="checkbox"]');
        const rows = tableBody.querySelectorAll('tr');

        checkboxes.forEach((checkbox, index) => {
            checkbox.checked = selectAllCheckbox.checked;
            rows[index].classList.toggle('selected', selectAllCheckbox.checked);
        });
    });

    // Обработчик клика по строке
    tableBody.addEventListener('click', async function (e) {
        const row = e.target.closest('tr');
        if (!row) return;
        const id = row.dataset.id;
        if (!id) return;

        if (hasUnsavedChanges) {
            const proceed = confirm("Есть несохранённые изменения. Сохранить?");
            hasUnsavedChanges = false;
            if (proceed) {
                await updateProduct(form)
            } else {
                return
            }
        }
        addingNewProduct = false;

        try {
            const res = await fetch(`/api/split-systems/${id}`);
            if (!res.ok) throw new Error("Ошибка загрузки товара");
            const data = await res.json();

            fillForm(data.item);
            hasUnsavedChanges = false;
            selectedId = id

        } catch (err) {
            console.error(err);
            alert("Ошибка загрузки данных товара");
        }
    });

    // Обработчик клика по чекбоксу
    tableBody.addEventListener('change', function (e) {
        if (e.target.tagName === 'INPUT' && e.target.type === 'checkbox') {
            const row = e.target.closest('tr');
            row.classList.toggle('selected', e.target.checked);
            updateSelectAllCheckbox();
        }
    });

    // Обработчик поиска
    searchBtn.addEventListener('click', function () {
        const searchTerm = searchInput.value.trim().toLowerCase();
        filterProducts(searchTerm);
    });

    // Обработчик добавления товара
    addBtn.addEventListener('click', async function () {
        console.log('Добавить новый товар');

        if (hasUnsavedChanges) {
            const proceed = confirm("Есть несохранённые изменения. Сохранить?");
            hasUnsavedChanges = false;
            if (proceed) {
                await updateProduct(form)
            } else {
                return
            }
        }
        resetFrom(form)

    });

    // Обработчик удаления товаров
    deleteBtn.addEventListener('click', function () {
        const selectedIds = getSelectedIds();
        if (selectedIds.length === 0) {
            alert('Пожалуйста, выберите товары для удаления');
            return;
        }

        console.log('Удалить товары с ID:', selectedIds.join(', '));
        // Здесь будет логика удаления
    });

    loadProducts();

    deleteBtn.addEventListener('click', () => {
        const ids = getSelectedProductIds();
        if (ids.length === 0) {
            alert("Выберите хотя бы один товар для удаления.");
            return;
        }
        deleteProducts(ids);
    });

});


async function updateProduct() {
    const formData = new FormData(form);
    // Приведение чекбокса к строке "true"/"false"
    formData.set("has_inverter", form.has_inverter.checked ? "true" : "false");
    formData.set("price", String(form.price.value * 100));

    try {
        const res = await fetch(`/api/split-systems/${selectedId}`, {
            method: "PATCH",
            body: formData
        });
        const data = await res.json();
        if (res.ok) {
            showNotify('Успех', "Товар успешно обновлен!");
            resetFrom(form)
            hasUnsavedChanges = false;
            addingNewProduct = true;
            loadProducts()
        } else {
            showErr("Ошибка: " + (data.error || "неизвестная ошибка"));
        }
    } catch (err) {
        console.error(err);
        showErr("Ошибка отправки запроса.")
    }
}

function resetFrom(form) {
    form.reset();
    document.querySelector(".input-file-text").textContent = "Максимум 5мб"
    document.getElementById("image-preview").src = "/web/static/uploads/placeholder.jpg"
    document.querySelectorAll('.checkbox-button').forEach((label) => {
        const input = label.querySelector('input');
        input.checked = false
        label.classList.remove('checked');
    })
}

async function createProduct() {
    const formData = new FormData(form);
    // Приведение чекбокса к строке "true"/"false"
    formData.set("has_inverter", form.has_inverter.checked ? "true" : "false");
    formData.set("price", String(form.price.value * 100));
    try {
        const res = await fetch("/api/split-systems", {
            method: "POST",
            body: formData
        });

        const data = await res.json();

        if (res.ok) {
            showNotify('Успех', "Товар успешно создан!");
            resetFrom(form)
            hasUnsavedChanges = false;
            loadProducts()
        } else {
            showErr("Ошибка: " + (data.error || "неизвестная ошибка"));
        }
    } catch (err) {
        console.error(err);
        showErr("Ошибка отправки запроса.")
    }
}

// Получение ID выбранных товаров
function getSelectedIds() {
    const selectedIds = [];
    const checkboxes = tableBody.querySelectorAll('input[type="checkbox"]:checked');

    checkboxes.forEach(checkbox => {
        const row = checkbox.closest('tr');
        const id = row.querySelector('td:nth-child(2)').textContent;
        selectedIds.push(id);
    });

    return selectedIds;
}

// Фильтрация товаров
function filterProducts(searchTerm) {
    if (!searchTerm) {
        renderProducts(allProducts);
        return;
    }

    const filtered = allProducts.filter(product =>
        product.title.toLowerCase().includes(searchTerm) ||
        String(product.id).toLowerCase().includes(searchTerm) ||
        product.brand.name.toLowerCase().includes(searchTerm) ||
        product.short_description.toLowerCase().includes(searchTerm)
    );

    renderProducts(filtered);
}


function fillForm(product) {
    document.querySelectorAll('.checkbox-button').forEach((label) => {
        const input = label.querySelector('input');
        input.checked = false
        label.classList.remove('checked');
    })

    form.title.value = product.title;
    form.short_description.value = product.short_description;
    form.long_description.value = product.long_description;
    form.brand_id.value = product.brand_id;
    form.type_id.value = product.type_id;
    form.price.value = product.price / 100;
    form.cooling_power.value = product.cooling_power;
    form.recommended_area.value = product.recommended_area;
    form.has_inverter.checked = product.has_inverter;
    form.energy_class_cooling_id.value = product.energy_class_cooling_id || "";
    form.energy_class_heating_id.value = product.energy_class_heating_id || "";
    form.min_noise_level.value = product.min_noise_level;
    form.max_noise_level.value = product.max_noise_level;

    form.internal_weight.value = product.internal_weight;
    form.internal_width.value = product.internal_width;
    form.internal_height.value = product.internal_height;
    form.internal_depth.value = product.internal_depth;

    document.querySelectorAll('.checkbox-button').forEach((label) => {
        const input = label.querySelector('input');
        product.modes.forEach(mode => {
            const mode_id = mode.id.toString()
            if (input.value === mode_id) {
                input.checked = true
                label.classList.add('checked');
            }
        })
    })

    form.external_weight.value = product.external_weight;
    form.external_width.value = product.external_width;
    form.external_height.value = product.external_height;
    form.external_depth.value = product.external_depth;

    document.querySelector('.input-file-text').textContent = product.image_url;

    // Устанавливаем картинку
    const imagePreview = document.getElementById('image-preview');
    if (product.image_url) {
        imagePreview.src = '/web/static/uploads/' + product.image_url;
    } else {
        imagePreview.src = '/web/static/uploads/placeholder.jpg';
    }

    // сбрасываем выбор файла
    form.image.value = '';
}

// Обновление состояния чекбокса "Выбрать все"
function updateSelectAllCheckbox() {
    const checkboxes = tableBody.querySelectorAll('input[type="checkbox"]');
    const checkedCount = Array.from(checkboxes).filter(cb => cb.checked).length;

    selectAllCheckbox.checked = checkedCount === checkboxes.length;
    selectAllCheckbox.indeterminate = checkedCount > 0 && checkedCount < checkboxes.length;
}

function getSelectedProductIds() {
    const checkboxes = tableBody.querySelectorAll('input[type="checkbox"]:checked');
    const ids = Array.from(checkboxes).map(checkbox => {
        return parseInt(checkbox.closest('tr').dataset.id, 10);
    });
    return ids;
}

async function deleteProducts(ids) {
    const confirmed = confirm(`Удалить ${ids.length} товар(ов)?`);
    if (!confirmed) return;

    try {
        await Promise.all(ids.map(id =>
            fetch(`/api/split-systems/${id}`, {
                method: 'DELETE'
            })
        ));

        // Удалить строки из таблицы сразу (оптимизация UX)
        ids.forEach(id => {
            const row = tableBody.querySelector(`tr[data-id="${id}"]`);
            if (row) row.remove();
        });

        alert("Удаление завершено.");
    } catch (error) {
        console.error("Ошибка удаления:", error);
        alert("Ошибка при удалении товаров.");
    }
}

// Настройка высоты таблицы
function adjustTableHeight() {
    const rows = tableBody.querySelectorAll('tr');
    if (rows.length > 0) {
        const rowHeight = rows[0].offsetHeight;
        const headerHeight = document.querySelector('thead tr').offsetHeight;
        const tableScroll = document.querySelector('.table-scroll');

        tableScroll.style.maxHeight = (rowHeight * visibleRows + headerHeight) + 'px';
        rows.forEach(row => row.style.display = '');
    }
}

// Загрузка данных (замените на реальный API запрос)
function formatPrice(price) {
    const intPrice = parseInt(price); // или Number(price)
    return new Intl.NumberFormat('ru-RU').format(intPrice / 100) + ' ₽';
}

function renderProducts(products) {
    totalCount.textContent = products.length;
    selectAllCheckbox.checked = false;
    selectAllCheckbox.indeterminate = false;

    tableBody.innerHTML = '';

    products.forEach(item => {
        const row = document.createElement('tr');
        row.dataset.id = item.id; // ← обязательно

        row.innerHTML = `
              <td class="checkbox-cell"><input type="checkbox"></td>
              <td>${item.id}</td>
              <td>${item.title}</td>
              <td>${item.short_description}</td>
              <td>${item.brand?.name || '—'}</td>
              <td>${item.type?.name || '—'}</td>
              <td>${formatPrice(item.price)}</td>
              <td>${item.has_inverter ? 'Да' : 'Нет'}</td>
              <td>${item.recommended_area}</td>
              <td>${item.cooling_power}</td>
              <td>Охлаждение: ${item.energy_class_cooling?.name || '—'}, Обогрев: ${item.energy_class_heating?.name || '—'}</td>
              <td>${item.min_noise_level} – ${item.max_noise_level} дБ</td>
              <td>${item.internal_width}×${item.internal_height}×${item.internal_depth} / ${item.internal_weight} </td>
              <td>${item.external_width}×${item.external_height}×${item.external_depth} / ${item.external_weight} </td>
              <td>${item.modes?.map(mode => mode.name).join(', ') || '—'}</td>
              <td>${item.average_rating}</td>           
            `;
        tableBody.appendChild(row);
        adjustTableHeight(products)
    });

}

function loadProducts() {
    fetch('/api/split-systems/')
        .then(response => {
            if (!response.ok) throw new Error("Ошибка при загрузке данных");
            return response.json();
        })
        .then(data => {
            allProducts = data.items;
            renderProducts(data.items || []);
        })
        .catch(error => {
            alert("Ошибка загрузки: " + error.message);
        });
}

function fetchOrders() {
    fetch("/api/order", {
        method: "GET",
        headers: {
            "Content-Type": "application/json"
        }
    })
        .then(response => response.json())
        .then(data => {
            const tbody = document.querySelector("#orders-table tbody");
            tbody.innerHTML = ""; // очистить перед вставкой

            data.data.forEach(order => {
                const row = document.createElement("tr");

                row.innerHTML = `
                    <td>${order.id}</td>
                    <td>${formatDate(order.created_at)}</td>
                    <td>${order.user.last_name + " " + order.user.first_name + " " + order.user.patronymic || "Неизвестно"}</td>
                    <td>${order.user?.email + " | " + order.user?.phone_number || "—"}</td>
                    <td>${renderSplitSystems(order.split_systems)}</td>
                    <td>${formatPrice(order.total_price) || 0}</td>
                `;

                const statusCell = document.createElement("td");
                const select = document.createElement("select");
                ["в обработке", "принят", "завершен"].forEach(status => {
                    const option = document.createElement("option");
                    option.value = status;
                    option.textContent = status;
                    if (order.status === status) {
                        option.selected = true;
                    }
                    select.appendChild(option);
                });

                select.addEventListener("change", () => {
                    updateOrderStatus(order.id, select.value);
                });

                statusCell.appendChild(select);
                row.appendChild(statusCell);
                tbody.appendChild(row);
            });
        })
        .catch(error => {
            console.error("Ошибка загрузки заказов:", error);
        });
}

function renderSplitSystems(splitSystems) {
    if (!Array.isArray(splitSystems) || splitSystems.length === 0) {
        return "—";
    }
    return splitSystems.map(system => {
        return `<li><a href="/products/${system.id}" class="split-link" target="_blank">${system.title}</a></li>`;
    }).join("");
}

function updateOrderStatus(orderId, newStatus) {
    fetch(`/api/order/${orderId}?status=${encodeURIComponent(newStatus)}`, {
        method: "PATCH"
    })
        .then(response => {
            if (!response.ok) {
                throw new Error("Ошибка обновления статуса");
            }
            return response.json();
        })
        .then(data => {
            showNotify("Статус обновлён", data.message);
        })
        .catch(error => {
            showErr("Ошибка при обновлении статуса:", error);
        });
}

function formatDate(dateString) {
    const date = new Date(dateString);
    return date.toLocaleString("ru-RU", {
        year: "numeric",
        month: "2-digit",
        day: "2-digit",
        hour: "2-digit",
        minute: "2-digit"
    });
}


document.querySelectorAll('.input-file input[type="file"]').forEach(function (input) {
    input.addEventListener('change', function () {
        const file = this.files[0];
        if (file) {
            const wrapper = this.closest('.input-file');
            const textElement = wrapper.querySelector('.input-file-text');
            if (textElement) {
                textElement.textContent = file.name;
            }
        }
    });
});