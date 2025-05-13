document.addEventListener('DOMContentLoaded', function() {
    const tableBody = document.querySelector('#products-table tbody');
    const searchBtn = document.getElementById('search-btn');
    const searchInput = document.getElementById('search-input');
    const addBtn = document.getElementById('add-btn');
    const deleteBtn = document.getElementById('delete-btn');
    const totalCount = document.getElementById('total-count');
    const selectAllCheckbox = document.getElementById('select-all');

    let allProducts = [];
    const visibleRows = 5;

    // Обработчик "Выбрать все"
    selectAllCheckbox.addEventListener('change', function() {
        const checkboxes = tableBody.querySelectorAll('input[type="checkbox"]');
        const rows = tableBody.querySelectorAll('tr');

        checkboxes.forEach((checkbox, index) => {
            checkbox.checked = selectAllCheckbox.checked;
            rows[index].classList.toggle('selected', selectAllCheckbox.checked);
        });
    });

    // Обработчик клика по строке
    tableBody.addEventListener('click', function(e) {
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
    tableBody.addEventListener('change', function(e) {
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
    searchBtn.addEventListener('click', function() {
        const searchTerm = searchInput.value.trim().toLowerCase();
        filterProducts(searchTerm);
    });

    // Обработчик добавления товара
    addBtn.addEventListener('click', function() {
        console.log('Добавить новый товар');
        // Здесь будет логика добавления
    });

    // Обработчик удаления товаров
    deleteBtn.addEventListener('click', function() {
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
            product.name.toLowerCase().includes(searchTerm) ||
            product.description.toLowerCase().includes(searchTerm)
        );

        renderProducts(filtered);
    }

    // Рендеринг товаров
    function renderProducts(products) {
        totalCount.textContent = products.length;
        selectAllCheckbox.checked = false;
        selectAllCheckbox.indeterminate = false;

        tableBody.innerHTML = '';

        products.forEach((item, index) => {
            const row = document.createElement('tr');
            row.innerHTML = `
                        <td class="checkbox-cell"><input type="checkbox"></td>
                        <td>${item.id}</td>
                        <td>${item.name}</td>
                        <td>${item.description}</td>
                        <td>${item.price}</td>
                        <td>...</td>
                        <td>...</td>
                        <td>...</td>
                    `;
            tableBody.appendChild(row);

            if (index >= visibleRows) {
                row.style.display = 'none';
            }
        });

        adjustTableHeight();
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
    function loadSampleData() {
        allProducts = [];
        for (let i = 1; i <= 20; i++) {
            allProducts.push({
                id: i,
                name: `Товар ${i}`,
                description: `Описание товара ${i}. Подробное описание характеристик и свойств товара.`,
                price: `${i * 1000} ₽`
            });
        }

        renderProducts(allProducts);
    }

    loadSampleData();
});