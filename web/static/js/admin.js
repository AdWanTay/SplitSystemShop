document.addEventListener('DOMContentLoaded', function () {
    const tableBody = document.querySelector('#products-table tbody');
    const searchBtn = document.getElementById('search-btn');
    const searchInput = document.getElementById('search-input');
    const addBtn = document.getElementById('add-btn');
    const deleteBtn = document.getElementById('delete-btn');
    const totalCount = document.getElementById('total-count');
    const selectAllCheckbox = document.getElementById('select-all');

    let allProducts = [];
    const visibleRows = 5;

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
    tableBody.addEventListener('click', function (e) {
        const row = e.target.closest('tr');
        if (!row) return;

        // Если кликнули на чекбокс, не переключаем выделение (обрабатывается отдельно)
        if (e.target.tagName === 'INPUT' && e.target.type === 'checkbox') {
            return;
        }

        const checkbox = row.querySelector('input[type="checkbox"]');
        checkbox.checked = !checkbox.checked;
        row.classList.toggle('selected', checkbox.checked);

        updateSelectAllCheckbox();
    });

    // Обработчик клика по чекбоксу
    tableBody.addEventListener('change', function (e) {
        if (e.target.tagName === 'INPUT' && e.target.type === 'checkbox') {
            const row = e.target.closest('tr');
            row.classList.toggle('selected', e.target.checked);
            updateSelectAllCheckbox();
        }
    });

    // Обновление состояния чекбокса "Выбрать все"
    function updateSelectAllCheckbox() {
        const checkboxes = tableBody.querySelectorAll('input[type="checkbox"]');
        const checkedCount = Array.from(checkboxes).filter(cb => cb.checked).length;

        selectAllCheckbox.checked = checkedCount === checkboxes.length;
        selectAllCheckbox.indeterminate = checkedCount > 0 && checkedCount < checkboxes.length;
    }

    // Обработчик поиска
    searchBtn.addEventListener('click', function () {
        const searchTerm = searchInput.value.trim().toLowerCase();
        filterProducts(searchTerm);
    });

    // Обработчик добавления товара
    addBtn.addEventListener('click', function () {
        console.log('Добавить новый товар');
        // Здесь будет логика добавления
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
              <td>${item.recommended_area} м²</td>
              <td>${item.cooling_power} кВт</td>
              <td>Охлаждение: ${item.energy_class_cooling?.name || '—'}, Обогрев: ${item.energy_class_heating?.name || '—'}</td>
              <td>${item.min_noise_level} – ${item.max_noise_level} дБ</td>
              <td>${item.internal_width}×${item.internal_height}×${item.internal_depth} мм / ${item.internal_weight} кг</td>
              <td>${item.external_width}×${item.external_height}×${item.external_depth} мм / ${item.external_weight} кг</td>
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
                allProducts = data.items
                renderProducts(data.items || []);
            })
            .catch(error => {
                alert("Ошибка загрузки: " + error.message);
            });
    }

    loadProducts()

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

    deleteBtn.addEventListener('click', () => {
        const ids = getSelectedProductIds();
        if (ids.length === 0) {
            alert("Выберите хотя бы один товар для удаления.");
            return;
        }
        deleteProducts(ids);
    });
});
