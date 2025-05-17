// calculator.js - скрипты для работы с калькулятором мощности охлаждения

document.addEventListener('DOMContentLoaded', function () {
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

    // Обновление при изменении значений
    form.addEventListener('input', calculate);

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
});