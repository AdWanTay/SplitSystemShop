const navContainer = document.querySelector('.question-navigation');
const questionButtons = navContainer.querySelectorAll('.question-number');

document.addEventListener('DOMContentLoaded', function() {
    let elapsedSeconds = 0;

    // 1. Получаем ID теста из URL
    const pathParts = window.location.pathname.split('/').filter(Boolean);
    const testId = pathParts[pathParts.length - 1];

    if (!testId || isNaN(testId)) {
        console.error('Invalid test ID');
        return;
    }

    let testData = null;
    let userAnswers = {};
    let currentQuestionIndex = 0;
    let timerInterval = null;

    // 2. Модальное окно для начала теста
    async function showStartModal() {
        try {
            const response = await fetch('/web/templates/test/modals/window.html');
            const html = await response.text();

            const modal = document.createElement('div');
            modal.innerHTML = html;
            document.body.appendChild(modal);

            // Настраиваем модальное окно
            const modalTitle = modal.querySelector('.modal-title');
            modalTitle.textContent = 'Готовы начать тест?';

            const modalBody = modal.querySelector('.modal-body');
            modalBody.innerHTML = '<p>Тест содержит вопросы только с одним правильным ответом.</p>';

            const modalDesc = modal.querySelector('.modal-description');
            modalDesc.innerHTML = '<p>На выполнение теста отводится <span style="color: #0055ff">15 минут</span>. Таймер запустится автоматически после начала.</p>';

            const startBtn = modal.querySelector('#mainBtn');
            startBtn.textContent = 'Начать тест';
            startBtn.id = 'start-test'; // Меняем ID для обработчика

            startBtn.addEventListener('click', function() {
                modal.remove();
                startTimer();
                showQuestion(0);
            });

            const escBtn = modal.querySelector('.cancel');
            escBtn.addEventListener('click', function() {
                window.close();
            });

            const escBtn2 = modal.querySelector('.modal-close');
            escBtn2.addEventListener('click', function() {
                window.close();
            });

        } catch (error) {
            console.error('Ошибка загрузки модального окна:', error);
            startTimer();
            showQuestion(0);
        }
    }


    // 3. Таймер
    function startTimer() {
        const timeElement = document.getElementById('time');
        let time = 15*60; // 15 минут в секундах

        timerInterval = setInterval(function() {
            time--;
            elapsedSeconds++;
            if (time <= 0) {
                clearInterval(timerInterval);
                finishTest();
                return;
            }

            const minutes = Math.floor(time / 60);
            const seconds = time % 60;
            timeElement.textContent = `${minutes}:${seconds < 10 ? '0' : ''}${seconds}`;
        }, 1000);
    }


    // 4. Загрузка данных теста
    async function loadTestData() {
        try {
            const response = await fetch(`/api/tests/${testId}`);
            if (!response.ok) throw new Error(`HTTP error! status: ${response.status}`);

            testData = await response.json();

            if (!testData.questions || !testData.questions.length) {
                throw new Error('No questions in test');
            }

            document.querySelector('.test-title').textContent = testData.title;
            initializeTest();
            showStartModal();
        } catch (error) {
            showErr('Не удалось загрузить тест. Пожалуйста, попробуйте позже.');
        }
    }


    // 5. Инициализация теста
    function initializeTest() {
        const questionNav = document.querySelector('.question-navigation');
        questionNav.innerHTML = '';

        testData.questions.forEach((question, index) => {
            const label = document.createElement('label');
            label.className = 'question-label';
            if (index === 0) label.classList.add('active');

            label.innerHTML = `
                <input class="question-number" type="radio" name="question" value="${index}" ${index === 0 ? 'checked' : ''}>
                <span>${index + 1}</span>
            `;
            questionNav.appendChild(label);
        });

        // Обработчики для навигации
        document.querySelectorAll('.question-number').forEach(radio => {
            radio.addEventListener('change', function() {
                if (this.checked) {
                    saveCurrentAnswer();
                    currentQuestionIndex = parseInt(this.value);
                    showQuestion(currentQuestionIndex);
                }
            });
        });

        // Кнопка "Далее"
        const nextButton = document.getElementById('next-button');
        if (nextButton) {
            nextButton.addEventListener('click', function() {
                saveCurrentAnswer();
                if (currentQuestionIndex < testData.questions.length - 1) {
                    currentQuestionIndex++;
                    const nextRadio = document.querySelector(`.question-number[value="${currentQuestionIndex}"]`);
                    if (nextRadio) nextRadio.checked = true;
                    showQuestion(currentQuestionIndex);
                }
            });
        }
    }


    // 6. Показать вопрос
    function showQuestion(index) {
        if (!testData || !testData.questions || index >= testData.questions.length) {
            console.error('Invalid question index');
            return;
        }

        const question = testData.questions[index];

        // Обновляем текст вопроса
        const questionTextEl = document.querySelector('.question-text');
        if (questionTextEl) {
            questionTextEl.textContent = `${index + 1}. ${question.question_text}`;
        }

        // Обновляем варианты ответов
        const answersContainer = document.querySelector('.answers-box');
        if (answersContainer) {
            answersContainer.innerHTML = `
                <div class="answers-title">Варианты ответов (один ответ):</div>
                <ul class="answers-list">
                    ${question.answers.map(answer => `
                        <li>
                            <label>
                                <input type="radio" name="answer" value="${answer.id}"
                                    ${userAnswers[question.id] && userAnswers[question.id].includes(answer.id) ? 'checked' : ''}>
                                ${answer.answer_text}
                            </label>
                        </li>
                    `).join('')}
                </ul>
            `;
        }

        // Обновляем активные классы в навигации
        document.querySelectorAll('.question-label').forEach((label, i) => {
            label.classList.toggle('active', i === index);
        });

        // Обновляем состояние кнопки "Далее"
        const nextButton = document.getElementById('next-button');
        if (nextButton) {
            nextButton.style.display = index === testData.questions.length - 1 ? 'none' : 'block';
        }
    }


    // 7. Сохранить текущий ответ
    function saveCurrentAnswer() {
        if (!testData || currentQuestionIndex >= testData.questions.length) return;

        const questionId = testData.questions[currentQuestionIndex].id;
        const selected = document.querySelector('input[name="answer"]:checked');

        const label = document.querySelector(`.question-navigation label:nth-child(${currentQuestionIndex + 1})`);

        if (selected) {
            userAnswers[questionId] = [parseInt(selected.value)];
            if (label) label.classList.add('answered'); // <--- Добавляем класс
        } else {
            delete userAnswers[questionId];
            if (label) label.classList.remove('answered'); // <--- Убираем класс, если ответ убрали
        }
    }



    // 8. Завершение теста
    async function finishTest() {
        clearInterval(timerInterval);
        saveCurrentAnswer();

        const resultData = {
            test_id: parseInt(testId),
            result: Object.fromEntries(
                testData.questions.map((question) => [
                    question.id,
                    userAnswers[question.id] ? userAnswers[question.id][0] : null
                ])
            )
        };

        try {
            const response = await fetch('/api/tests/send-result', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(resultData)
            });

            if (response.ok) {
                const result = await response.json(); // Получаем {correct: X, total: Y}

                // Показываем модальное окно с результатами
                await showResultsModal(result.correct, result.total, elapsedSeconds);

            } else {
                throw new Error(`Server returned ${response.status}`);
            }
        } catch (error) {
            showErr('Ошибка при отправке результатов. Пожалуйста, попробуйте ещё раз.');
        }
    }

    async function showResultsModal(correctAnswers, totalQuestions, elapsedSeconds) {
        try {
            // Загружаем шаблон модального окна
            const response = await fetch('/web/templates/test/modals/window.html');
            const html = await response.text();

            const modal = document.createElement('div');
            modal.innerHTML = html;

            // Настраиваем содержимое
            const percentage = Math.round((correctAnswers / totalQuestions) * 100);
            const minutesSpent = Math.floor(elapsedSeconds / 60);
            const secondsSpent = elapsedSeconds % 60;
            const formattedTimeSpent = `${minutesSpent}:${secondsSpent < 10 ? '0' : ''}${secondsSpent}`;

            modal.querySelector('.modal-title').textContent = 'Результаты тестирования';

            modal.querySelector('.modal-close').style.display = "none";

            const modalBody = modal.querySelector('.modal-body');
            modalBody.innerHTML = `
                <div class="results-container">
                    <div class="result-circle">
                        <div class="circle-progress" style="--percentage: ${percentage}">
                            <span>${percentage}%</span>
                        </div>
                    </div>
                    <div class="result-details">
                        <div class="result-item">
                            <span class="label">Правильных ответов:</span>
                            <span class="value correct">${correctAnswers}</span>
                        </div>
                        <div class="result-item">
                            <span class="label">Всего вопросов:</span>
                            <span class="value">${totalQuestions}</span>
                        </div>
                        <div class="result-item">
                            <span class="label">Затраченное время:</span>
                            <span class="value">${formattedTimeSpent}</span>
                        </div>
                    </div>
                </div>
                <div class="result-message">
                    ${getResultMessage(percentage)}
                </div>
            `;

            // Настраиваем кнопки
            const modalActions = modal.querySelector('.modal-actions');
            modalActions.innerHTML = `
            <button id="restartTest" class="modal-button primary">Пройти тест снова</button>
            <button id="returnToCourse" class="modal-button cancel">Вернуться к курсу</button>
        `;

            // Удаляем описание (не нужно)
            modal.querySelector('.modal-description').remove();

            // Добавляем обработчики
            document.body.appendChild(modal);

            document.getElementById('restartTest').addEventListener('click', function() {
                modal.remove();
                location.reload(); // Или ваша функция для перезапуска теста
            });

            document.getElementById('returnToCourse').addEventListener('click', function() {
                modal.remove();
                window.location.href = '/profile#bought-courses'; // Или ваш путь к курсу
            });

            // Добавляем стили
            const style = document.createElement('style');
            style.textContent = `
            .results-container {
                display: flex;
                align-items: center;
                gap: 30px;
                margin: 20px 0;
            }
            
            .result-circle {
                position: relative;
                width: 120px;
                height: 120px;
            }
            
            .circle-progress {
                width: 100%;
                height: 100%;
                border-radius: 50%;
                background: conic-gradient(
                    #4CAF50 calc(var(--percentage) * 3.6deg),
                    #f0f0f0 0
                );
                display: flex;
                align-items: center;
                justify-content: center;
            }
            
            .circle-progress span {
                background: white;
                width: 90px;
                height: 90px;
                border-radius: 50%;
                display: flex;
                align-items: center;
                justify-content: center;
                font-weight: bold;
                font-size: 20px;
            }
            
            .result-details {
                display: flex;
                flex-direction: column;
                gap: 10px;
            }
            
            .result-item {
                display: flex;
                justify-content: space-between;
                min-width: 200px;
            }
            
            .label {
                color: #666;
            }
            
            .value {
                font-weight: bold;
            }
            
            .correct {
                color: #4CAF50;
            }
            
            .result-message {
                margin-top: 20px;
                padding: 15px;
                background: #f8f9fa;
                border-radius: 8px;
                text-align: center;
            }
        `;
            document.head.appendChild(style);

        } catch (error) {
            console.error('Ошибка загрузки модального окна:', error);
            showNotify("Тест завершен!", `Правильных ответов: ${correctAnswers}/${totalQuestions}`);
        }
    }

    function getResultMessage(percentage) {
        if (percentage == 100) return 'Отличный результат! Вы прекрасно усвоили материал.';
        if (percentage >= 70) return 'Хороший результат! Есть немного тем для повторения.';
        if (percentage >= 50) return 'Средний результат. Рекомендуем повторить материал.';
        return 'Низкий результат. Внимательно изучите материал и попробуйте снова.';
    }


    // 9. Модальное окно подтверждения завершения (использует шаблон force-finish.html)
    window.forceFinishModal = async function() {
        saveCurrentAnswer();

        const answeredCount = Object.keys(userAnswers).length;
        const totalQuestions = testData.questions.length;

        if (answeredCount === totalQuestions) {
            // Все вопросы отвечены — сразу завершаем тест без подтверждения
            finishTest();
        } else {
            // Есть неотвеченные — показываем подтверждение
            try {
                const response = await fetch('/web/templates/test/modals/force-finish.html');
                if (!response.ok) throw new Error('Не удалось загрузить модальное окно');

                const html = await response.text();
                const parser = new DOMParser();
                const doc = parser.parseFromString(html, 'text/html');
                const modalTemplate = doc.querySelector('.modal').outerHTML;

                const modalContainer = document.createElement('div');
                modalContainer.innerHTML = modalTemplate;
                const modal = modalContainer.firstChild;

                modal.querySelector('.modal-button.primary').addEventListener('click', function() {
                    modal.remove();
                    finishTest();
                });
                document.body.appendChild(modal);

            } catch (error) {
                console.error('Ошибка загрузки модального окна:', error);
            }
        }
    };


    // Запускаем загрузку теста
    loadTestData();
});